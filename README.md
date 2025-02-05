# basin-cli

[![License](https://img.shields.io/github/license/tablelandnetwork/basin-cli.svg)](./LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg)](https://github.com/RichardLitt/standard-readme)

> Continuously publish data from your database to the Tableland network.

# Table of Contents

- [Install](#install)
- [Postgres Setup](#postgres-setup)
  - [Self-hosted](#self-hosted)
  - [Amazon RDS](#amazon-rds)
  - [Supabase](#supabase)
- [Create a publication](#create-a-publication)
- [Start replicating a publication](#start-replicating-a-publication)
- [Upload a Parquet file](#upload-a-parquet-file)
- [Listing Publications](#listing-publications)
- [Listing Deals](#listing-deals)
- [Running](#running)
- [Run tests](#run-tests)
- [Generate Cap'N Proto code](#generate-capn-proto-code)
- [Retrieving](#retrieving)

# Background

Tableland Basin is a secure and verifiable open data platform. The Basin CLI is a tool that allows you to continuously replicate a table or view from your database to the network. Currently, only PostgreSQL is supported.

> 🚧 Basin is currently not in a production-ready state. Any data that is pushed to the network may be subject to deletion. 🚧

# Usage

## Install

```bash
git clone https://github.com/tablelandnetwork/basin-cli.git
cd basin-cli
go install ./cmd/basin
```

## Postgres Setup

### Self-hosted

- Make sure you have access to a superuser role. For example, you can create a new role such as `CREATE ROLE basin WITH PASSWORD NULL LOGIN SUPERUSER;`.
- Check that your Postgres installation has the [wal2json](https://github.com/eulerto/wal2json) plugin installed.
- Check if logical replication is enabled:

  ```sql
  SHOW wal_level;
  ```

  The `wal_level` setting must be set to logical: `ALTER SYSTEM SET wal_level = logical;`.

- Restart the database in order for the new `wal_level` to take effect (be careful!).

### Amazon RDS

- Make sure you have a user with the `rds_superuser` role, and use `psql` to connect to your database.

  ```console
  psql -h [HOST] -U [USER] -d [DATABASE]
  ```

- Check if logical replication is enabled:

  ```sql
  SELECT name, setting
  FROM pg_settings
  WHERE name = 'rds.logical_replication';
  ```

- If it's on, go to [Create a publication](#create-a-publication)
- If it's off, follow the next steps:
  - [Create a custom RDS parameter group](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithDBInstanceParamGroups.html#USER_WorkingWithParamGroups.Creating)
  - After creation, edit it and set the `rds.logical_replication` parameter to `1`
  - [Associate the recently created parameter group with you DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithDBInstanceParamGroups.html#USER_WorkingWithParamGroups.Associating)
    - You can choose **Apply immediately** to apply the changes immediately
    - You'll probably need to reboot the instance for changes to take effect (be careful!)
- After reboot, check if logical replication is enabled

### Supabase

- Log into the [Supabase](https://supabase.io/) dashboard and go to your project, or create a new one.
- Check if logical replication is enabled. This should be the default setting, so you shouldn't have to change anything. You can do this in the `SQL Editor` section on the left hand side of the Supabase dashboard by running `SHOW wal_level;` query, which should log `logical`.
- You can find the database connection information on the left hand side under `Project Settings` > `Database`. You will need the `Host`, `Port`, `Database`, `Username`, and `Password` to connect to your database.
  - When you create a publication, the `--dburi` should follow this format:
    ```sh
    postgresql://postgres:[PASSWORD]@db.[PROJECT_ID].supabase.co:5432/postgres
    ```

## Create a publication

_Publications_ define the data you are pushing to Basin.

Basin uses public key authentication, so you will need an Ethereum style (ECDSA, secp256k1) wallet to create a new publication. You can use an existing wallet or set up a new one with `basin wallet create`. Your private key is only used locally for signing.

```bash
basin wallet create [FILENAME]
```

A new private key will be written to `FILENAME`.

The name of a publication contains a `namespace` (e.g. `my_company`) and the name of an existing database relation (e.g. `my_table`), separated by a period (`.`). Use `basin publication create` to create a new publication. See `basin publication create --help` for more info.

```bash
basin publication create  --dburi [DBURI] --address [WALLET_ADDRESS] namespace.relation_name
```

🚧 Basin currently only replicates `INSERT` statements, which means that it only replicates append-only data (e.g., log-style data). Row updates and deletes will be ignored. 🚧

## Start replicating a publication

Use `basin publication start` to start a daemon that will continuously push changes to the underlying table/view to the network. See `basin publication start --help` for more info.

```bash
basin publication start --private-key [PRIVATE_KEY] namespace.relation_name
```

## Upload a Parquet file

Before uploading a Parquet file, you need to [Create a publication](#create-a-publication), if not already created. You can omit the `--dburi` flag, in this case.

Then, use `basin publication upload` to upload a Parquet file.

```bash
basin publication upload --name [namespace.relation_name] --private-key [PRIVATE_KEY] filepath
```

You can attach a timestamp to that file upload, e.g. 

```bash
basin publication upload --name [namespace.relation_name] --private-key [PRIVATE_KEY] --timestamp 1699984703 filepath

# or use data format
basin publication upload --name [namespace.relation_name] --private-key [PRIVATE_KEY] --timestamp 2006-01-02 filepath

# or use RFC3339 format
basin publication upload --name [namespace.relation_name] --private-key [PRIVATE_KEY] --timestamp 2006-01-02T15:04:05Z07:00 filepath
```

If a timestamp is not provided, the CLI will assume the timestamp is the current client epoch in UTC.

## Listing Publications

You can list the publications from an address by running:

```bash
basin publication list --address [ETH_ADDRESS]
```

## Listing Deals

You can list deals of a given publication by running:

```bash
basin publication deals --publication [PUBLICATION] --latest 5
```

Deals command accept `--before`,`--after` , and `--at` flags to filter deals by timestamp

```bash
# examples
basin publication deals --publication demotest.data --at 1699569502
basin publication deals --publication demotest.data --before 2023-11-09T19:38:23-03:00
basin publication deals --publication demotest.data --after 2023-11-09  
```

## Retrieving

```bash
basin publication retrieve bafybeifr5njnrw67yyb2h2t7k6ukm3pml4fgphsxeurqcmgmeb7omc2vlq
```

# Development

## Running

You can make use of the scripts inside `scripts` to facilitate running the CLI locally without building.

```bash
# Starting the Provider Server
PORT=8888 ./scripts/server.sh

# Create a wallet
./scripts/run.sh wallet create pk.out

# Start replicating
./scripts/run.sh publication start --private-key [PRIVATE_KEY] namespace.relation_name
```

## Run tests

```bash
make test
```

Note: One of the tests requires Docker Engine to be running.

## Generate Cap'N Proto code

```bash
make generate
```

# Contributing

PRs accepted.

Small note: If editing the README, please conform to the
[standard-readme](https://github.com/RichardLitt/standard-readme) specification.

# License

MIT AND Apache-2.0, © 2021-2023 Tableland Network Contributors

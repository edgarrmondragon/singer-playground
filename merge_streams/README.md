# Merge streams

- [Motivation](#Motivation)
- [Instructions](#Instructions)

## Motivation

There are situations in multi-tenant architectures where the data schema is shared between a number of sources (databases, filesystem directories, etc.)
but the tap discovery behavior results in separate streams. Rather that modifying the existing tap or writing a new tap from scratch, streams can be merged by applying the
correct transformations to each Singer message type.

Examples:

- A MongoDB architecture where there is a database for each customer but collections have the same structure.
- An S3 bucket with CSV files, where the files for each customer are under a different prefix, but the files have the same schemas.

## Instructions

Create a virtual environment and install the python dependencies.

```shell
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

To merge the streams in the tap output, you just have to do

```shell
python tap.py | python map.py
```

## Test with a Singer target

The output of the above command can itself be piped into a singer target.

For example, to use with `target-sqlite` install it

```shell
pipx install target-sqlite
```

and validate that the streams are merged and land in a single table in the target:

```shell
$ python tap.py | python map.py | target-sqlite -c target_sqlite.json
$ sqlite3 example.db -markdown "select * from orders"
| __partition | id | nested__a | nested__b |        __loaded_at         |
|-------------|----|-----------|-----------|----------------------------|
| 123         | 1  | 42        | 0         | 2021-06-26 01:23:27.920824 |
| 123         | 2  | 314       | 1         | 2021-06-26 01:23:27.921140 |
| 456         | 1  | 1513      | 0         | 2021-06-26 01:23:27.921658 |
```

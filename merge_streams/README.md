# Merge streams

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

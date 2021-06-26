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
| __tenant |    id     | nested__a | nested__b |        __loaded_at         |
|----------|-----------|-----------|-----------|----------------------------|
| 123      | a93u01982 | 42        | 0         | 2021-06-25 23:18:27.880821 |
| 123      | 0hi984h92 | 314       | 1         | 2021-06-25 23:18:27.880931 |
| 456      | 1ijf98203 | 1513      | 0         | 2021-06-25 23:18:27.881026 |
```

# tap-data-jobs

Singer tap for positions posted in [Data Stack Jobs](https://datastackjobs.com/).

## Instructions

```shell
$ go build
$ ./tap-data-jobs | target-sqlite --config target_sqlite.json
$ sqlite3 datajobs.db -markdown "
    select
        company_name,
        position,
        location_type,
        published_at
    from jobs
    where location_type = 'remote'
    order by published_at desc
    limit 15
"
|    company_name     |             position             | location_type |     published_at     |
|---------------------|----------------------------------|---------------|----------------------|
| Sticker Mule        | Machine Learning Engineer        | remote        | 2021-06-29T07:41:48Z |
| Stripe              | Machine Learning Engineer        | remote        | 2021-06-29T07:38:23Z |
| Binance             | Big Data DevOps Engineer         | remote        | 2021-06-25T17:40:47Z |
| GitLab              | Senior Data Engineer, Analytics  | remote        | 2021-06-25T17:37:54Z |
| GitLab              | Senior Data Engineer, Analytics  | remote        | 2021-06-22T12:56:20Z |
| Inflection          | Senior Data Engineer             | remote        | 2021-06-22T12:55:08Z |
| Uptake Technologies | Data Engineer                    | remote        | 2021-06-22T12:52:51Z |
| ReCharge Payments   | Staff Data Engineer              | remote        | 2021-06-11T17:15:54Z |
| Zapier              | Data Engineer                    | remote        | 2021-06-11T17:13:24Z |
| Welocalize          | Director, Machine Learning       | remote        | 2021-06-11T17:08:13Z |
| Socure              | Sr. Data Engineer                | remote        | 2021-06-11T17:03:25Z |
| Iterative           | Developer Advocate               | remote        | 2021-06-07T20:23:09Z |
| Coinbase            | Senior Machine Learning Engineer | remote        | 2021-06-07T20:09:11Z |
| Urbint              | Manager, Data Engineering        | remote        | 2021-06-07T20:06:23Z |
| Urbint              | Senior Data Engineering          | remote        | 2021-06-07T20:03:53Z |
```

# FilterOp

we also support '==', but we're not exporting that to swagger, cause it can't deal with it
add >, >=, <, <= to that list, it breaks codegen


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `FilterOpAnd`           | and                     |
| `FilterOpOr`            | or                      |
| `FilterOpNot`           | not                     |
| `FilterOpEq`            | eq                      |
| `FilterOpNotEqual`      | !=                      |
| `FilterOpNe`            | ne                      |
| `FilterOpExists`        | exists                  |
| `FilterOpNotExists`     | not-exists              |
| `FilterOpGt`            | gt                      |
| `FilterOpGte`           | gte                     |
| `FilterOpLt`            | lt                      |
| `FilterOpLte`           | lte                     |
| `FilterOpStartsWith`    | starts-with             |
| `FilterOpNotStartsWith` | not-starts-with         |
| `FilterOpEndsWith`      | ends-with               |
| `FilterOpNotEndsWith`   | not-ends-with           |
| `FilterOpContains`      | contains                |
| `FilterOpNotContains`   | not-contains            |
| `FilterOpRegexp`        | regexp                  |
| `FilterOpNotRegexp`     | not-regexp              |
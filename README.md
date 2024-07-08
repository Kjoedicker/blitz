
<h1 align="center">Blitz</h3>

# About The Project
A tool for determining the limits of your app with targeted blitzes

```
With requests as artillery
Load tests are a blitz
```

# Usage

## Test plans

### Example file

```yml test-plan.yml
---
host: http://localhost:3000
Targets:
  - description: 'GET /health/mysql'
    method: "GET"
    path: "/health/mysql"
    headers: ~
    hits: 100
    interval: second
    duration: 2
  - description: 'GET /health/redis'
    method: "GET"
    path: "/health/mysql"
    headers: ~
    hits: 1000
    interval: second
    duration: 2
```

### Root configurations
| Field       | Details                                            |
| ----------- | -------------------------------------------------- |
| host        | Being targeted                                     |
| Targets     | Different paths, with different intervals and hits |

### Target Configurations
| Field       | Details                                                |
| ----------- | ------------------------------------------------------ |
| description | Of whats happening                                     |
| method      | HTTP methods                                           |
| path        | Of the endpoint                                        |
| headers     | To add to the request                                  |
| hits        | Total requests to make every configured `interval`     |
| interval    | Interval to run the configured number of `hits`        |
| duration    | The duration that `hits` will execute every `interval` |

## Execution

When executed, blitz will look for a `test-plan.yml` file in the current directory. You can also specify the test plan by passing a `--target-file` flag with the path.

```shell
kjoedicker@arch ~ % ./blitz --target-file test-plan.yml

▄▄▄▄   ██▓    ██▄▄▄█████▒███████▒
▓█████▄▓██▒   ▓██▓  ██▒ ▓▒ ▒ ▒ ▄▀░
▒██▒ ▄█▒██░   ▒██▒ ▓██░ ▒░ ▒ ▄▀▒░ 
▒██░█▀ ▒██░   ░██░ ▓██▓ ░  ▄▀▒   ░
░▓█  ▀█░██████░██░ ▒██▒ ░▒███████▒
░▒▓███▀░ ▒░▓  ░▓   ▒ ░░  ░▒▒ ▓░▒░▒
▒░▒   ░░ ░ ▒  ░▒ ░   ░   ░░▒ ▒ ░ ▒
 ░    ░  ░ ░   ▒ ░ ░     ░ ░ ░ ░ ░
 ░         ░  ░░           ░ ░    
          ░                  ░        
Test Plan Number: 0 Request Group: 1 Request Number 1 Response Time: 0.004990 Errors: <nil> 
Test Plan Number: 0 Request Group: 1 Request Number 2 Response Time: 0.002833 Errors: <nil> 
Test Plan Number: 0 Request Group: 1 Request Number 3 Response Time: 0.002370 Errors: <nil> 
Test Plan Number: 0 Request Group: 1 Request Number 4 Response Time: 0.003792 Errors: <nil> 
Test Plan Number: 0 Request Group: 1 Request Number 5 Response Time: 0.002648 Errors: <nil> 
Test Plan Number: 0 Request Group: 2 Request Number 1 Response Time: 0.002620 Errors: <nil> 
Test Plan Number: 0 Request Group: 2 Request Number 2 Response Time: 0.002490 Errors: <nil> 
Test Plan Number: 0 Request Group: 2 Request Number 3 Response Time: 0.002523 Errors: <nil> 
Test Plan Number: 0 Request Group: 2 Request Number 4 Response Time: 0.002855 Errors: <nil> 
Test Plan Number: 0 Request Group: 2 Request Number 5 Response Time: 0.002772 Errors: <nil> 
Test Plan Number: 0 Request Group: 3 Request Number 1 Response Time: 0.002332 Errors: <nil> 
Test Plan Number: 0 Request Group: 3 Request Number 2 Response Time: 0.002259 Errors: <nil> 
Test Plan Number: 0 Request Group: 3 Request Number 3 Response Time: 0.002207 Errors: <nil> 
Test Plan Number: 0 Request Group: 3 Request Number 4 Response Time: 0.003105 Errors: <nil> 
Test Plan Number: 0 Request Group: 3 Request Number 5 Response Time: 0.002483 Errors: <nil> 
Test Plan Number: 0 Request Group: 4 Request Number 1 Response Time: 0.002468 Errors: <nil> 
Test Plan Number: 0 Request Group: 4 Request Number 2 Response Time: 0.002469 Errors: <nil> 
Test Plan Number: 0 Request Group: 4 Request Number 3 Response Time: 0.003077 Errors: <nil> 
Test Plan Number: 0 Request Group: 4 Request Number 4 Response Time: 0.003054 Errors: <nil> 
Test Plan Number: 0 Request Group: 4 Request Number 5 Response Time: 0.003257 Errors: <nil> 
Test Plan Number: 0 Request Group: 5 Request Number 1 Response Time: 0.003294 Errors: <nil> 
Test Plan Number: 0 Request Group: 5 Request Number 2 Response Time: 0.003077 Errors: <nil> 
Test Plan Number: 0 Request Group: 5 Request Number 3 Response Time: 0.003031 Errors: <nil> 
Test Plan Number: 0 Request Group: 5 Request Number 4 Response Time: 0.002556 Errors: <nil> 
Test Plan Number: 0 Request Group: 5 Request Number 5 Response Time: 0.003569 Errors: <nil>
...
```

There is also a `--format` flag for setting the format of the printed request results. 

Currently supports: `text` and `csv`.

```shell
kjoedicker@arch ~ % ./blitz --target-file test-plan.yml --format=csv

▄▄▄▄   ██▓    ██▄▄▄█████▒███████▒
▓█████▄▓██▒   ▓██▓  ██▒ ▓▒ ▒ ▒ ▄▀░
▒██▒ ▄█▒██░   ▒██▒ ▓██░ ▒░ ▒ ▄▀▒░ 
▒██░█▀ ▒██░   ░██░ ▓██▓ ░  ▄▀▒   ░
░▓█  ▀█░██████░██░ ▒██▒ ░▒███████▒
░▒▓███▀░ ▒░▓  ░▓   ▒ ░░  ░▒▒ ▓░▒░▒
▒░▒   ░░ ░ ▒  ░▒ ░   ░   ░░▒ ▒ ░ ▒
 ░    ░  ░ ░   ▒ ░ ░     ░ ░ ░ ░ ░
 ░         ░  ░░           ░ ░    
          ░                  ░        
test_plan_number,request_group,request_number,response_time,error_response
0,1,1,0.022877,<nil>
0,1,2,0.003424,<nil>
0,1,3,0.003336,<nil>
0,1,4,0.003406,<nil>
0,1,5,0.003744,<nil>
0,2,1,0.003749,<nil>
0,2,2,0.003331,<nil>
0,2,3,0.004240,<nil>
0,2,4,0.002546,<nil>
0,2,5,0.003934,<nil>
0,3,1,0.003401,<nil>
0,3,2,0.003043,<nil>
0,3,3,0.003264,<nil>
0,3,4,0.004571,<nil>
0,3,5,0.003470,<nil>
0,4,1,0.003380,<nil>
0,4,2,0.003768,<nil>
0,4,3,0.003534,<nil>
0,4,4,0.002685,<nil>
0,4,5,0.004360,<nil>
0,5,1,0.004094,<nil>
0,5,2,0.003457,<nil>
0,5,3,0.002425,<nil>
0,5,4,0.003210,<nil>
0,5,5,0.002810,<nil>
0,6,1,0.002681,<nil>
0,6,2,0.002722,<nil>
0,6,3,0.003560,<nil>
0,6,4,0.003339,<nil>
0,6,5,0.003848,<nil>
...
```

The logo that is printed at the start is goes through `stderr`, freeing up `stdout` for printed request results.

This way the following `~ % ./blitz --format=csv > test-results.csv` results in a file that only contains headers and rows.

The logo can also be disabled by providing the `--print-logo=false` flag

```shell
kjoedicker@arch ~ % ./blitz --target-file test-plan.yml --format=csv --print-logo=false

test_plan_number,request_group,request_number,response_time,error_response
0,1,1,0.004684,<nil>
0,1,2,0.002428,<nil>
0,1,3,0.002652,<nil>
0,1,4,0.003209,<nil>
0,1,5,0.003162,<nil>
0,2,1,0.002575,<nil>
0,2,2,0.003284,<nil>
0,2,3,0.003205,<nil>
0,2,4,0.003218,<nil>
0,2,5,0.002806,<nil>
0,3,1,0.002636,<nil>
0,3,2,0.003304,<nil>
0,3,3,0.003232,<nil>
0,3,4,0.002537,<nil>
0,3,5,0.002681,<nil>
0,4,1,0.002438,<nil>
0,4,2,0.002277,<nil>
0,4,3,0.002366,<nil>
0,4,4,0.002592,<nil>
0,4,5,0.002572,<nil>
...
```
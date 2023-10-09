
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
  - description: 'GET /'
    method: "GET"
    path: "/"
    headers: ~
    hits: 1000
    interval: 2
  - description: 'GET /moo'
    method: "GET"
    path: "/"
    headers: ~
    hits: 675
    interval: 1
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


Scenario: GET /
1000 requests every 2 minutes for 1 minutes

Request group: 0
Request 1: 0.017142 seconds 
Request 2: 0.017813 seconds 
Request 3: 0.017534 seconds 
Request 4: 0.017292 seconds 
Request 5: 0.016807 seconds 
Request 6: 0.018161 seconds 
Request 7: 0.017936 seconds 
Request 8: 0.015053 seconds 

Request group: 1
Request 1: 0.010527 seconds 
Request 2: 0.010758 seconds 
Request 3: 0.008926 seconds 
Request 4: 0.009085 seconds 
Request 5: 0.009358 seconds 
Request 6: 0.009874 seconds 
Request 7: 0.010084 seconds 
Request 8: 0.009577 seconds 

Request group: 2
Request 1: 0.002936 seconds 
Request 2: 0.003148 seconds 
Request 3: 0.003853 seconds 
Request 4: 0.002284 seconds 
Request 5: 0.003585 seconds 
Request 6: 0.003318 seconds 
Request 7: 0.002503 seconds 
Request 8: 0.001966 seconds 

...
```
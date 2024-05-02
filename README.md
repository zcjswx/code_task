## Prerequisites

- Go
- Docker

## How to run

In the project's root dir

- Setup Postgres, and init tables
    ```shell
    make setup
    ```

- Run unit tests
  ```shell
  make test
  ```

- Download XML, parse and save in database
    ```shell
    make run_parse
    ```

  and paste the url of a XML file, exp
    ```
    https://raw.githubusercontent.com/zcjswx/code_task/main/misc/graph.xml
    ```


- Run queries and get answers

    ```shell
    make run_query
    ```
  and paste your queries, exp
  ```json
  {
    "queries": [
      {
        "paths": {
          "start": "3",
          "end": "1"
        }
      },
      {
        "cheapest": {
          "start": "3",
          "end": "1"
        }
      },
      {
        "cheapest": {
          "start": "3",
          "end": "2"
        }
      },
      {
        "cheapest": {
          "start": "1",
          "end": "0"
        }
      }
    ]
  }
  ```
  
## Todo

- Add config file
- Run as a web service
- Add Customized error type to do error handling
- CI pipline
- dockerfile
- More comments, tests and documents
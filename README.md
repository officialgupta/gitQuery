# gitQuery
A minimal CLI app to query GitHub users.

## Installation
    go get github.com/officialgupta/gitQuery
Can make global by:

    go install gitQuery

## Usage
- Help:

    ```
    $ gitQuery
    Usage: gitQuery [options]
    Options:
        -u, --user string
            Search Users
    ```  

- Single User Query:

    ```
    $ gitQuery -u officialgupta
    Searching user(s): [officialgupta]
    Username:        officialgupta
    Name:            Mayank Gupta
    Email:
    Bio:
    ```

- Multi-User Query:

    ```
    $ gitQuery -u pmbenjamin,defunkt                                                                                                                                    
    Searching user(s): [officialgupta abakhai]
    Username:        officialgupta
    Name:            Mayank Gupta
    Email:
    Bio:

    Username:        abakhai
    Name:            Ami
    Email:
    Bio:
    ```


Based on the go CLI introduction by petermbenjamin

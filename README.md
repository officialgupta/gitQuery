# gitQuery
A minimal CLI app to query GitHub users.

## Installation
    go get github.com/officialgupta/gitgo
Can make global by:

    go install gitgo

## Usage
- Help:

    ```
    $ gitgo
    Usage: gitgo [options]
    Options:
        -u, --user string
            Search Users
    ```  

- Single User Query:

    ```
    $ gitgo -u officialgupta
    Searching user(s): [officialgupta]
    Username:        officialgupta
    Name:            Mayank Gupta
    Email:
    Bio:
    ```

- Multi-User Query:

    ```
    $ gitgo -u pmbenjamin,defunkt                                                                                                                                    
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

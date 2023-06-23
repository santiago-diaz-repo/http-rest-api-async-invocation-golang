# http-rest-api-versioning-golang

This project shows two ways to manage versioning in an API, they are Accept header and URL. Both ways are implementing based on [json-api](https://jsonapi.org/) standard:

1. By sending a header called `Accept` whose value can be `application/vnd.json+api;version=`_number_, you can ask for a response in a particular version.
2. By invoking the URL /v*number*, you can ask for a response in a particular version.

## How to test?
I have left a **postman collection** and a **visual code project** in order for you to test this API.

## What will you get as response?
When calling any of both APIs you will get something like the following response:

```json
{
    "data": [
        {
            "type": "customers",
            "id": "221234",
            "attributes": {
                "firstName": "Test1",
                "lastName": "Test1"
            },
            "relationships": {
                "loans": {
                    "data": [
                        {
                            "type": "loans",
                            "id": "77565543"
                        }
                    ],
                    "links": {
                        "related": "http://localhost:8080/customers/221234/loans"
                    }
                }
            },
            "links": {
                "self": "http://localhost:8080/customers/221234"
            },
            "meta": {
                "detail": "This is version 2"
            }
        },
        {
            "type": "customers",
            "id": "964532",
            "attributes": {
                "firstName": "Test2",
                "lastName": "Test2"
            },
            "relationships": {
                "loans": {
                    "data": [
                        {
                            "type": "loans",
                            "id": "5466878"
                        },
                        {
                            "type": "loans",
                            "id": "67654590"
                        }
                    ],
                    "links": {
                        "related": "http://localhost:8080/customers/964532/loans"
                    }
                }
            },
            "links": {
                "self": "http://localhost:8080/customers/964532"
            },
            "meta": {
                "detail": "This is version 2"
            }
        }
    ],
    "included": [
        {
            "type": "loans",
            "id": "77565543",
            "attributes": {
                "balance": 343354.66
            }
        },
        {
            "type": "loans",
            "id": "5466878",
            "attributes": {
                "balance": 343354.66
            }
        },
        {
            "type": "loans",
            "id": "67654590",
            "attributes": {
                "balance": 100006.5
            }
        }
    ]
}

```

As you can see, based on `github.com/google/jsonapi`, it is possible to generate a response that follows the json API standard.

If you have any questions, do not hesitate to get in touch

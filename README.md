# fiz
A file wizard command line tool

### Initial wizard json format

Variables are assumed to be string types unless there is a default value

Conditions can either be 1 variable, or multiple variables with operators. Operators include:

    - `||` Or
    - `&&` And


```
{
    "templates": [
        { "name": "hello.cpp" },
        { "name": "hello.h" },
        { "name": "mock_hello.h", "condition": ["CreateMock"] }
    ],
    "variables": [
        {
            "name": "ClassName",
            "type": "string"
        },
        {
            "name": "CreateMock",
            "type": "bool"
        },
        {
            "name": "CreateNamespace",
            "type": "bool"
        },
        {
            "name": "Namespace",
            "condition": ["CreateNamespace"]
        }
    ]
}
```

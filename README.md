# fiz
A file wizard command line tool

### Initial wizard json format

```
{
    "templates": [
        "hello.cpp",
        "hello.h"
    ],
    "variables": [
        {
            "name": "ClassName",
            "type": "string"
        },
        {
            "name": "CreateNamespace",
            "type": "bool",
            "default": false
        },
        {
            "name": "Namespace",
            "type": "string",
            "condition": ["CreateNamespace"]
        }
    ]
}
```

# tenryumon: sink
![Test Status](https://github.com/nyelonong/nsqsink/actions/workflows/test.yml/badge.svg)

# how to run
## config example
```
[
  {
    "id": "channel-name",
    "topic": "topic-name",
    "source": {
      "nsqd": ["127.0.0.1:4150"],
      "nsqlookupd": ["127.0.0.1:4161"]
    },
    "max_attempt": 5,
    "max_in_flight": 4,
    "concurrent": 2,
    "sinker": {
      "type": "http",
      "parser": {
        "type": "json",
        "template": "{\"name\":\"$user.name\",\"age\":$user.age}"
      },
      "config": {
        "http": {
          "url": "http://something.com",
          "method": "POST",
          "headers": {
            "Accept": "application/json"
          }
        },
        "file": {
          "file_name": ""
        }
      }
    },
    "active": true
  }
]
```

## with incoming message example
```
http://localhost:4151/pub?topic=topic-name
{
    "user": {
        "name": "steph",
        "age": 123
    }
}
```

## the result will be
```
POST http://something.com
json payload: {"name":"steph","age":123}
```

# command to run
go build && ./sink -config-path="./example-config.json"

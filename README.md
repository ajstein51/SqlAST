# SQL Parser
The objective for this program is simply to be able to parse a given sql definition and return a Json of its results with a reasonable amount of accuracy to be used further in other programs. I have quite a lot of hopes and dreams for this repo but have yet to think about it in its totality so TBD on future plans.

An example being

```
SELECT *
FROM Database.dbo.TableName
```

The following should return something as follows (To change as this gets further along)

{
    "Token": {
        "Type": Select,
        "Value": "SELECT"
    },
    "Token": {
        "Type": WildCard,
        "Value": "*"
    },
    "Token": {
        "Type": From,
        "Value": "FROM"
    },
    "Token": {
        "Type": Identifier,
        "Value": "Database"
    },
    "Token": {
        "Type": TBD,
        "Value": "."
    },
    "Token": {
        "Type": Identifier,
        "Value": "Database"
    },
    "Token": {
        "Type": TBD,
        "Value": "."
    },
    "Token": {
        "Type": Identifier,
        "Value": "Database"
    }
}

<sup>PS: this is also serving as my intro to Golang and a learning project for it</sup>

## To Do
### Testing
- Individual run tests IE: go run -test FileName.go
- Test flag testing IE: `go run -test`

### Lexing
- Figure out how to distinguish a * from a wildcard vs operator
- other token types
    - Research what the . operator is called

# Version

Version is a simple package to allow you to manage version for your application

# Installation

``` 
go get go.ajitem.com/version
```

# Usage

```
var v *version.Version

// the package ensures that the version used follows semver spec
v = version.NewVersion("1.2.3-alpha2.0+201907015-e3fef0c")

log.Println(v)
// prints 1.2.3-alpha2.0+201907015-e3fef0c

// or 

v.PrettyPrint = true
fmt.Println(v)
// prints 1.2.3 (alpha2.0) (201907015 e3fef0c)
```

# Contributing

Any PRs for adding new features / fixing bugs are more than welcome :)
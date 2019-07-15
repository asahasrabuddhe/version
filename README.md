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

fmt.Println(v)


// or 

log.Println(v)
```

# Contributing

Any PRs for adding new features / fixing bugs are more than welcome :)
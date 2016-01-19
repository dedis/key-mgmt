#CONIKS Protos

Copyright (C) 2015 Princeton University.

http://www.coniks.org

## Original documentation

See the [original documentation](https://raw.githubusercontent.com/coniks-sys/coniks-ref-implementation/master/protos/README.md) for more information.

## Compiling into Golang
Assuming you have protoc installed, run the following commands from the projects `protos` directory.

```
 protoc --go_out=../common/ *.proto
```

You will have to fix the files manually.

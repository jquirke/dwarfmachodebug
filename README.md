# dwarfmachodebug
tool used for inspecting DWARF LPT on Mach-O files. Used in root causing golang bug: https://github.com/golang/go/issues/54320

## Build 
```
% go build 
```

## Run 
```
% ./dwarfmachodebug /path/to/go/linked/macho/file <optional CU names>
```

If no CU (Compile Unit) is specified all compile units are dumped: e.g.: 

```
% ../dwarfmachodebug/dwarfmachodebug ./helloworld main math
dumping compile unit main
[1000A36B0 - 1000A3730] (/Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:5)
unnamed file
0x1000A36B0: IsStmt= true prologend=false epilogbegin=false /Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:5
0x1000A36BC: IsStmt= true prologend= true epilogbegin=false /Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:5
0x1000A36C8: IsStmt= true prologend=false epilogbegin=false /Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:6
0x1000A36CC: IsStmt=false prologend=false epilogbegin=false /Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:6
0x1000A3710: IsStmt= true prologend=false epilogbegin=false /Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:7
0x1000A371C: IsStmt= true prologend=false epilogbegin=false /Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:5
0x1000A3730: IsStmt= true prologend=false epilogbegin=false /Users/qjeremy/go/src/qjeremy/helloworld/hello-world.go:5
dumping compile unit math
[100062430 - 100062490] (/usr/local/go/src/math/unsafe.go:12)
unnamed file
0x100062430: IsStmt= true prologend= true epilogbegin=false /usr/local/go/src/math/unsafe.go:12
0x100062440: IsStmt=false prologend=false epilogbegin=false /usr/local/go/src/math/unsafe.go:12
0x100062460: IsStmt=false prologend=false epilogbegin=false /usr/local/go/src/math/unsafe.go:12
0x100062460: IsStmt= true prologend= true epilogbegin=false /usr/local/go/src/math/unsafe.go:23
0x100062470: IsStmt=false prologend=false epilogbegin=false /usr/local/go/src/math/unsafe.go:23
0x100062490: IsStmt=false prologend=false epilogbegin=false /usr/local/go/src/math/unsafe.go:23
```

The intepretation of the LPT values are as specified in [DWARF section 6.2](https://dwarfstd.org/doc/Dwarf3.pdf)


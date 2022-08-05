package main

import (
	"debug/dwarf"
	"debug/macho"
	"fmt"
	"io"
	"os"
)

// hacked up utility to debug DWARF ARM64 issue using Go's libraries
// qjeremy@uber.com

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("usage: %v <machofile> [compileunit [compileunit ... ] ]\n", os.Args[0])
		return
	}
	macho, err := macho.Open(os.Args[1])
	if err != nil {
		fmt.Printf("macho open: %v", err)
		return
	}

	cuNames := make(map[string]bool)
	for _, cuName := range os.Args[2:] {
		cuNames[cuName] = true
	}

	dwarfdata, err := macho.DWARF()
	if err != nil {
		fmt.Printf("dwarf: %v", err)
		return
	}
	reader := dwarfdata.Reader()
nextCu:
	for {
		entry, err := reader.Next()
		if err != nil {
			fmt.Printf("reader err: %v", err)
			return
		}
		if entry == nil {
			fmt.Printf("finished all entries\n")
			break
		}
		if entry.Tag == dwarf.TagCompileUnit {
			cuName, _ := entry.Val(dwarf.AttrName).(string)
			if cuNames[cuName] || len(cuNames) == 0 {
				fmt.Printf("dumping compile unit %v\n", cuName)
				lineReader, err := dwarfdata.LineReader(entry)
				if err != nil {
					fmt.Printf("lineReader err: %v", err)
					continue
				}
				files := lineReader.Files()
				for _, file := range files {
					if file == nil {
						fmt.Printf("unnamed file\n")
						continue
					}
					//fmt.Printf("file: %v\n", file.Name)
				}
				var lineEntry dwarf.LineEntry
				for {
					if err := lineReader.Next(&lineEntry); err != nil {
						if err != io.EOF {
							fmt.Printf("error on linereader.Next, skipping to next CU %v", err)
						}
						continue nextCu
					}
					fmt.Printf("0x%08X: IsStmt=%5t prologend=%5t epilogbegin=%5t %s:%d \n",
						lineEntry.Address, lineEntry.IsStmt, lineEntry.PrologueEnd, lineEntry.EpilogueBegin,
						lineEntry.File.Name, lineEntry.Line,
					)
				}
			}

		}
		reader.SkipChildren()
	}
}

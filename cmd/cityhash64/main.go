package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hungrybirder/cityhash"
)

const zeroSeed uint64 = 0

func showTitle() {
	fmt.Println("==========Show CityHash64 Demos==========")
	fmt.Println()
}

func showUsageAndExit() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("cityhash64 [-seed=N] str [str...]")
	fmt.Println()
	os.Exit(0)
}

func process(keys []string, seed uint64) {
	showTitle()
	fmt.Printf("%-30s\t%-30s\t%s\n", "hash", "hash(hex)", "key")
	if seed == zeroSeed {
		for _, key := range keys {
			keyBytes := []byte(key)
			keyLen := uint32(len(keyBytes))
			hashVal := cityhash.CityHash64(keyBytes, keyLen)
			fmt.Printf("%-30d\t%-30x\t%s\n", hashVal, hashVal, key)
		}
	} else {
		for _, key := range keys {
			keyBytes := []byte(key)
			keyLen := uint32(len(keyBytes))
			hashVal := cityhash.CityHash64WithSeed(keyBytes, keyLen, seed)
			fmt.Printf("%-30d\t%-30x\t%s\n", hashVal, hashVal, key)
		}
		fmt.Println()
		fmt.Printf("Seed=%d\n", seed)
	}
	fmt.Println()
}

func main() {
	seedPtr := flag.Uint64("seed", uint64(0), "Hash seed, optional, must greater than 0")
	flag.Parse()
	keys := flag.Args()
	if len(keys) == 0 {
		showUsageAndExit()
	}

	process(flag.Args(), *seedPtr)
}

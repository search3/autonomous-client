package main

import (
	ipfs "github.com/ipfs/go-ipfs-api"
)

func main() {
	ipfsShell := ipfs.NewShell("localhost:5001")
	err := ipfsShell.Get("QmRxJSPm37oihhVkFJ5BpmmRpN4rFxkev22A76eykdrYbw", "facenet_ipfs")
	if err != nil {
		panic(err)
	}
}

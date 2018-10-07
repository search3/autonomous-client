package ipfsconnector 

import (
	ipfs "github.com/ipfs/go-ipfs-api"
)

type ipfsClient struct {
	shell Shell
	host string 
}

func NewIpfsClient() (*ipfsClient) {

}


func GetModel(string: hash) error {

} 


func UploadPhoto(string: photo_file) (string, error) {

}

func main() {
	ipfsShell := ipfs.NewShell("localhost:5001")
	err := ipfsShell.Get("QmRxJSPm37oihhVkFJ5BpmmRpN4rFxkev22A76eykdrYbw", "facenet_ipfs")
	if err != nil {
		panic(err)
	}
}

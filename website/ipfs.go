package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/rootfounders/backend/config"
)

var ipfsShell *ipfs.Shell

func InitIpfs(cnf *config.Config) (err error) {
	url := cnf.IpfsNode
	if url == "" {
		return errors.New("ipfs node url required")
	}
	ipfsShell = ipfs.NewShellWithClient(url, http.DefaultClient)
	if !ipfsShell.IsUp() {
		return fmt.Errorf("ipfs node down at %s", url)
	}
	return
}

func Pin(r io.Reader) (cid string, err error) {
	cid, err = ipfsShell.Add(r)
	return
}

func ReadFromIpfs(cid string) (content []byte, err error) {
	r, err := ipfsShell.Cat(cid)
	if err != nil {
		return
	}
	content, err = io.ReadAll(r)
	if err != nil {
		return
	}
	r.Close()

	return
}

package node

import finalcachepb "github.com/nienie/finalcache/pb"

//Node ...
type Node interface {

	finalcachepb.FinalCacheHandler

	//WhoAmI ...
	WhoAmI() string

	//Lookup ...
	Lookup(key string) string

	//Run ...
	Run() error
}


package chatserv

import "github.com/TasToys/eventstructs/common"

type JoinChannelCommand struct {
	node     common.Node
	platform string
	channel  string
}

type JoinedChannelResponce struct {
	node     common.Node
	platform string
	channel  string
}

type PartChannelCommand struct {
	node     common.Node
	platform string
	channel  string
}

type PartChannelResponce struct {
	node     common.Node
	platform string
	channel  string
}

type ListChannelsCommand struct {
	node common.Node
}

type ListChannelsRespons struct {
	node         common.Node
	listChannels []string
}

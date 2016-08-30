// Auto-generated by avdl-compiler v1.3.4 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/remote.avdl

package chat1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type MessageBoxed struct {
	ServerHeader    *MessageServerHeader `codec:"serverHeader,omitempty" json:"serverHeader,omitempty"`
	ClientHeader    MessageClientHeader  `codec:"clientHeader" json:"clientHeader"`
	HeaderSignature SignatureInfo        `codec:"headerSignature" json:"headerSignature"`
	BodyCiphertext  EncryptedData        `codec:"bodyCiphertext" json:"bodyCiphertext"`
	BodySignature   SignatureInfo        `codec:"bodySignature" json:"bodySignature"`
	KeyGeneration   int                  `codec:"keyGeneration" json:"keyGeneration"`
}

type ThreadViewBoxed struct {
	Messages   []MessageBoxed `codec:"messages" json:"messages"`
	Pagination *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetInboxRemoteArg struct {
	Pagination *Pagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetThreadRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MarkAsRead     bool           `codec:"markAsRead" json:"markAsRead"`
	Pagination     *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type PostRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MessageBoxed   MessageBoxed   `codec:"messageBoxed" json:"messageBoxed"`
}

type NewConversationRemoteArg struct {
	IdTriple ConversationIDTriple `codec:"idTriple" json:"idTriple"`
}

type GetMessagesRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MessageIDs     []MessageID    `codec:"messageIDs" json:"messageIDs"`
}

type MarkAsReadArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MsgID          MessageID      `codec:"msgID" json:"msgID"`
}

type RemoteInterface interface {
	GetInboxRemote(context.Context, *Pagination) (InboxView, error)
	GetThreadRemote(context.Context, GetThreadRemoteArg) (ThreadViewBoxed, error)
	PostRemote(context.Context, PostRemoteArg) (MessageID, error)
	NewConversationRemote(context.Context, ConversationIDTriple) (ConversationID, error)
	GetMessagesRemote(context.Context, GetMessagesRemoteArg) ([]MessageBoxed, error)
	MarkAsRead(context.Context, MarkAsReadArg) error
}

func RemoteProtocol(i RemoteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.remote",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getInboxRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetInboxRemote(ctx, (*typedArgs)[0].Pagination)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getThreadRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetThreadRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetThreadRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetThreadRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetThreadRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postRemote": {
				MakeArg: func() interface{} {
					ret := make([]PostRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostRemoteArg)(nil), args)
						return
					}
					ret, err = i.PostRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationRemote": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationRemoteArg)(nil), args)
						return
					}
					ret, err = i.NewConversationRemote(ctx, (*typedArgs)[0].IdTriple)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getMessagesRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetMessagesRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetMessagesRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetMessagesRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetMessagesRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"markAsRead": {
				MakeArg: func() interface{} {
					ret := make([]MarkAsReadArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]MarkAsReadArg)
					if !ok {
						err = rpc.NewTypeError((*[]MarkAsReadArg)(nil), args)
						return
					}
					err = i.MarkAsRead(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RemoteClient struct {
	Cli rpc.GenericClient
}

func (c RemoteClient) GetInboxRemote(ctx context.Context, pagination *Pagination) (res InboxView, err error) {
	__arg := GetInboxRemoteArg{Pagination: pagination}
	err = c.Cli.Call(ctx, "chat.1.remote.getInboxRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetThreadRemote(ctx context.Context, __arg GetThreadRemoteArg) (res ThreadViewBoxed, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.getThreadRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) PostRemote(ctx context.Context, __arg PostRemoteArg) (res MessageID, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.postRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) NewConversationRemote(ctx context.Context, idTriple ConversationIDTriple) (res ConversationID, err error) {
	__arg := NewConversationRemoteArg{IdTriple: idTriple}
	err = c.Cli.Call(ctx, "chat.1.remote.newConversationRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetMessagesRemote(ctx context.Context, __arg GetMessagesRemoteArg) (res []MessageBoxed, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.getMessagesRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) MarkAsRead(ctx context.Context, __arg MarkAsReadArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.markAsRead", []interface{}{__arg}, nil)
	return
}

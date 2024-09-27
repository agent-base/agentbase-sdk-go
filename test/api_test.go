package test

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/agent-base/agentbase-sdk-go"
)

var (
	host         = "这里填写你的host"
	apiSecretKey = "这里填写你的api secret key"
)

func TestApi3(t *testing.T) {
	var c = &agentbase.ClientConfig{
		Host:         host,
		ApiSecretKey: apiSecretKey,
	}
	var client = agentbase.NewClientWithConfig(c)

	ctx := context.Background()

	var (
		ch  = make(chan agentbase.ChatMessageStreamChannelResponse)
		err error
	)

	ch, err = client.Api().ChatMessagesStream(ctx, &agentbase.ChatMessageRequest{
		Query: "你是谁?",
		User:  "这里换成你创建的",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	var (
		strBuilder strings.Builder
		cId        string
	)
	for {
		select {
		case <-ctx.Done():
			t.Log("ctx.Done", strBuilder.String())
			return
		case r, isOpen := <-ch:
			if !isOpen {
				goto done
			}
			strBuilder.WriteString(r.Answer)
			cId = r.ConversationID
			log.Println("Answer2", r.Answer, r.ConversationID, cId, r.ID, r.TaskID)
		}
	}

done:
	t.Log(strBuilder.String())
	t.Log(cId)
}

func TestMessages(t *testing.T) {
	var cId = "ec373942-2d17-4f11-89bb-f9bbf863ebcc"
	var err error
	ctx := context.Background()

	// messages
	var messageReq = &agentbase.MessagesRequest{
		ConversationID: cId,
		User:           "jiuquan AI",
	}

	var client = agentbase.NewClient(host, apiSecretKey)

	var msg *agentbase.MessagesResponse
	if msg, err = client.Api().Messages(ctx, messageReq); err != nil {
		t.Fatal(err.Error())
		return
	}
	j, _ := json.Marshal(msg)
	t.Log(string(j))
}

func TestMessagesFeedbacks(t *testing.T) {
	var client = agentbase.NewClient(host, apiSecretKey)
	var err error
	ctx := context.Background()

	var id = "72d3dc0f-a6d5-4b5e-8510-bec0611a6048"

	var res *agentbase.MessagesFeedbacksResponse
	if res, err = client.Api().MessagesFeedbacks(ctx, &agentbase.MessagesFeedbacksRequest{
		MessageID: id,
		Rating:    agentbase.FeedbackLike,
		User:      "jiuquan AI",
	}); err != nil {
		t.Fatal(err.Error())
	}

	j, _ := json.Marshal(res)

	log.Println(string(j))
}

func TestConversations(t *testing.T) {
	var client = agentbase.NewClient(host, apiSecretKey)
	var err error
	ctx := context.Background()

	var res *agentbase.ConversationsResponse
	if res, err = client.Api().Conversations(ctx, &agentbase.ConversationsRequest{
		User: "jiuquan AI",
	}); err != nil {
		t.Fatal(err.Error())
	}

	j, _ := json.Marshal(res)

	log.Println(string(j))
}

func TestConversationsRename(t *testing.T) {
	var client = agentbase.NewClient(host, apiSecretKey)
	var err error
	ctx := context.Background()

	var res *agentbase.ConversationsRenamingResponse
	if res, err = client.Api().ConversationsRenaming(ctx, &agentbase.ConversationsRenamingRequest{
		ConversationID: "ec373942-2d17-4f11-89bb-f9bbf863ebcc",
		Name:           "rename!!!",
		User:           "jiuquan AI",
	}); err != nil {
		t.Fatal(err.Error())
	}

	j, _ := json.Marshal(res)

	log.Println(string(j))
}

func TestParameters(t *testing.T) {
	var client = agentbase.NewClient(host, apiSecretKey)
	var err error
	ctx := context.Background()

	var res *agentbase.ParametersResponse
	if res, err = client.Api().Parameters(ctx, &agentbase.ParametersRequest{
		User: "jiuquan AI",
	}); err != nil {
		t.Fatal(err.Error())
	}

	j, _ := json.Marshal(res)

	log.Println(string(j))
}

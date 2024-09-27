# AgentBase Go SDK
This is the Go SDK for the AgentBase API, which allows you to easily integrate AgentBase into your Go applications.

## Install
```bash
go get github.com/agent-base/agentbase-sdk-go
```

## Usage
After installing the SDK, you can use it in your project like this:

```go
package main

import (
	"context"
	"log"
	"strings"

	"github.com/agent-base/agentbase-sdk-go"
)

func main() {
	var (
		ctx = context.Background()
		c = agentbase.NewClient("your-agentbase-server-host", "your-api-key-here")

		req = &agentbase.ChatMessageRequest{
			Query: "your-question",
			User: "your-user",
		}

		ch chan agentbase.ChatMessageStreamChannelResponse
		err error
	)

	if ch, err = c.Api().ChatMessagesStream(ctx, req); err != nil {
		return
	}

	var strBuilder strings.Builder

	for {
		select {
		case <-ctx.Done():
			return
		case streamData, isOpen := <-ch:
			if err = streamData.Err; err != nil {
				log.Println(err.Error())
				return
			}
			if !isOpen {
				log.Println(strBuilder.String())
				return
			}

			strBuilder.WriteString(streamData.Answer)
		}
	}
}

```

## License
This SDK is released under the MIT License.
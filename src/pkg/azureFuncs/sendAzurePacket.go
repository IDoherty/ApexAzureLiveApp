package azureFuncs

import (
	"context"
	"fmt"
	"time"

	/*/ Azure Event Hub Libraries
	"github.com/Azure/azure-amqp-common-go/persist"
	"github.com/Azure/azure-event-hubs-go"
	//*/
	"github.com/Azure/azure-service-bus-go"
)

func sendAzurePacket(strMsgChan <-chan string) {

	var tMsg *servicebus.Message
	var strMsg string

	connStr := "Endpoint=sb://fenwaybus.servicebus.windows.net/;SharedAccessKeyName=FenwayTopicPolicy;SharedAccessKey=XNPixakOFUu+7I4LZZ3DrNSfCJQtnoQSPasLWwgLPUc=;EntityPath=fenwaytopic"
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable SERVICEBUS_CONNECTION_STRING not set")
		return
	}

	// Create a client to communicate with a Service Bus Namespace.
	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a client to communicate with the queue. (The queue must have already been created, see `QueueManager`)
	topic, err := ns.NewTopic("fenwayTopic")
	if err != nil {
		fmt.Println("FATAL: ", err)
		return
	}
	fmt.Println()
	fmt.Println("Topic Name - ", topic.Name)

	for {
		messageTimeout := time.Now().Add(10000 * time.Millisecond)

		strMsg = <-strMsgChan

		/*/ Incoming String Printout
		fmt.Println("string in - ", strMsg)
		fmt.Println()
		//*/

		if strMsg != "[]" {
			tMsg = buildFenwayMessage("Fenway", "18", strMsg)

			/*/ Output Packet Payload
			fmt.Println("output string", string(strMsg))
			fmt.Println()
			//*/

			go func(tMsg *servicebus.Message) {
				ctx, cancel := context.WithDeadline(context.Background(), messageTimeout)
				defer cancel()
				err := topic.Send(ctx, tMsg)
				if err != nil {
					fmt.Println("Error Sending to Topic: ", err)
				}
			}(tMsg)
		}

	}
}

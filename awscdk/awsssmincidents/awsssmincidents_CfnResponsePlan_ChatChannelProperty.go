package awsssmincidents


// The AWS Chatbot chat channel used for collaboration during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chatChannelProperty := &chatChannelProperty{
//   	chatbotSns: []*string{
//   		jsii.String("chatbotSns"),
//   	},
//   }
//
type CfnResponsePlan_ChatChannelProperty struct {
	// The SNS targets that AWS Chatbot uses to notify the chat channel of updates to an incident.
	//
	// You can also make updates to the incident through the chat channel by using the SNS topics.
	ChatbotSns *[]*string `field:"optional" json:"chatbotSns" yaml:"chatbotSns"`
}


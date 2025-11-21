package awsssmincidents


// The  chat channel used for collaboration during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chatChannelProperty := &ChatChannelProperty{
//   	ChatbotSns: []*string{
//   		jsii.String("chatbotSns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-chatchannel.html
//
type CfnResponsePlan_ChatChannelProperty struct {
	// The Amazon  targets that  uses to notify the chat channel of updates to an incident.
	//
	// You can also make updates to the incident through the chat channel by using the Amazon  topics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-chatchannel.html#cfn-ssmincidents-responseplan-chatchannel-chatbotsns
	//
	ChatbotSns *[]*string `field:"optional" json:"chatbotSns" yaml:"chatbotSns"`
}


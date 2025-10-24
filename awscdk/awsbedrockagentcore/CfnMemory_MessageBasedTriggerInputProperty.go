package awsbedrockagentcore


// The message based trigger input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageBasedTriggerInputProperty := &MessageBasedTriggerInputProperty{
//   	MessageCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-messagebasedtriggerinput.html
//
type CfnMemory_MessageBasedTriggerInputProperty struct {
	// The memory trigger condition input for the message based trigger message count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-messagebasedtriggerinput.html#cfn-bedrockagentcore-memory-messagebasedtriggerinput-messagecount
	//
	MessageCount *float64 `field:"optional" json:"messageCount" yaml:"messageCount"`
}


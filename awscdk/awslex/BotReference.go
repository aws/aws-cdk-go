package awslex


// A reference to a Bot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botReference := &BotReference{
//   	BotArn: jsii.String("botArn"),
//   	BotId: jsii.String("botId"),
//   }
//
type BotReference struct {
	// The ARN of the Bot resource.
	BotArn *string `field:"required" json:"botArn" yaml:"botArn"`
	// The Id of the Bot resource.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
}


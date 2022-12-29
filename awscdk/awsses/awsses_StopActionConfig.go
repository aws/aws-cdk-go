package awsses


// StopAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stopActionConfig := &stopActionConfig{
//   	scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	topicArn: jsii.String("topicArn"),
//   }
//
type StopActionConfig struct {
	// The scope of the StopAction.
	//
	// The only acceptable value is RuleSet.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the stop action is taken.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}


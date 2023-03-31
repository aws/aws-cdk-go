package awsiot


// Parameters to define a mitigation action that publishes findings to Amazon SNS.
//
// You can implement your own custom actions in response to the Amazon SNS messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publishFindingToSnsParamsProperty := &publishFindingToSnsParamsProperty{
//   	topicArn: jsii.String("topicArn"),
//   }
//
type CfnMitigationAction_PublishFindingToSnsParamsProperty struct {
	// The ARN of the topic to which you want to publish the findings.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}


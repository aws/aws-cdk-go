package awslogs


// Properties for defining a `CfnDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDestinationProps := &cfnDestinationProps{
//   	destinationName: jsii.String("destinationName"),
//   	roleArn: jsii.String("roleArn"),
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	destinationPolicy: jsii.String("destinationPolicy"),
//   }
//
type CfnDestinationProps struct {
	// The name of the destination.
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the physical target where the log events are delivered (for example, a Kinesis stream).
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	DestinationPolicy *string `field:"optional" json:"destinationPolicy" yaml:"destinationPolicy"`
}


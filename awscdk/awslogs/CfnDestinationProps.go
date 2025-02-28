package awslogs


// Properties for defining a `CfnDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDestinationProps := &CfnDestinationProps{
//   	DestinationName: jsii.String("destinationName"),
//   	RoleArn: jsii.String("roleArn"),
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	DestinationPolicy: jsii.String("destinationPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html
//
type CfnDestinationProps struct {
	// The name of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationname
	//
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the physical target where the log events are delivered (for example, a Kinesis stream).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationpolicy
	//
	DestinationPolicy *string `field:"optional" json:"destinationPolicy" yaml:"destinationPolicy"`
}


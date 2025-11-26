package previewawslogsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDestinationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDestinationMixinProps := &CfnDestinationMixinProps{
//   	DestinationName: jsii.String("destinationName"),
//   	DestinationPolicy: jsii.String("destinationPolicy"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetArn: jsii.String("targetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html
//
type CfnDestinationMixinProps struct {
	// The name of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationname
	//
	DestinationName *string `field:"optional" json:"destinationName" yaml:"destinationName"`
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationpolicy
	//
	DestinationPolicy *string `field:"optional" json:"destinationPolicy" yaml:"destinationPolicy"`
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags that have been assigned to this delivery destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the physical target where the log events are delivered (for example, a Kinesis stream).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-targetarn
	//
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}


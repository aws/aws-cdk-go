package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeliverySource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliverySourceProps := &CfnDeliverySourceProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	LogType: jsii.String("logType"),
//   	ResourceArn: jsii.String("resourceArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverysource.html
//
type CfnDeliverySourceProps struct {
	// The unique name of the delivery source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverysource.html#cfn-logs-deliverysource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of log that the source is sending.
	//
	// For valid values for this parameter, see the documentation for the source service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverysource.html#cfn-logs-deliverysource-logtype
	//
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// The Amazon Resource Name (ARN) that uniquely identifies this delivery source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverysource.html#cfn-logs-deliverysource-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The tags that have been assigned to this delivery source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverysource.html#cfn-logs-deliverysource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


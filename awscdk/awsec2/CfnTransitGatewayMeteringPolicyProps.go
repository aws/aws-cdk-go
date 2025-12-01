package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitGatewayMeteringPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayMeteringPolicyProps := &CfnTransitGatewayMeteringPolicyProps{
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//
//   	// the properties below are optional
//   	MiddleboxAttachmentIds: []*string{
//   		jsii.String("middleboxAttachmentIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicy.html
//
type CfnTransitGatewayMeteringPolicyProps struct {
	// The ID of the transit gateway associated with the metering policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicy.html#cfn-ec2-transitgatewaymeteringpolicy-transitgatewayid
	//
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The IDs of the middlebox attachments associated with the metering policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicy.html#cfn-ec2-transitgatewaymeteringpolicy-middleboxattachmentids
	//
	MiddleboxAttachmentIds *[]*string `field:"optional" json:"middleboxAttachmentIds" yaml:"middleboxAttachmentIds"`
	// The tags assigned to the transit gateway metering policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicy.html#cfn-ec2-transitgatewaymeteringpolicy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


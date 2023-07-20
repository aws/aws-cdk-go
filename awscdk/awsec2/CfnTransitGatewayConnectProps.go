package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitGatewayConnect`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayConnectProps := &CfnTransitGatewayConnectProps{
//   	Options: &TransitGatewayConnectOptionsProperty{
//   		Protocol: jsii.String("protocol"),
//   	},
//   	TransportTransitGatewayAttachmentId: jsii.String("transportTransitGatewayAttachmentId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html
//
type CfnTransitGatewayConnectProps struct {
	// The Connect attachment options.
	//
	// - protocol (gre).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-options
	//
	Options interface{} `field:"required" json:"options" yaml:"options"`
	// The ID of the attachment from which the Connect attachment was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-transporttransitgatewayattachmentid
	//
	TransportTransitGatewayAttachmentId *string `field:"required" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// The tags for the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


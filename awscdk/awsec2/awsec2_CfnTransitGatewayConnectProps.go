package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTransitGatewayConnect`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayConnectProps := &cfnTransitGatewayConnectProps{
//   	options: &transitGatewayConnectOptionsProperty{
//   		protocol: jsii.String("protocol"),
//   	},
//   	transportTransitGatewayAttachmentId: jsii.String("transportTransitGatewayAttachmentId"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayConnectProps struct {
	// The Connect attachment options.
	//
	// - protocol (gre).
	Options interface{} `field:"required" json:"options" yaml:"options"`
	// The ID of the attachment from which the Connect attachment was created.
	TransportTransitGatewayAttachmentId *string `field:"required" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// The tags for the attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


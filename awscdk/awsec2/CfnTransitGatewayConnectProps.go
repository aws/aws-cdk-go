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


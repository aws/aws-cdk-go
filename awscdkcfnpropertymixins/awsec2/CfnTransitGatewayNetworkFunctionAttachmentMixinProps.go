package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTransitGatewayNetworkFunctionAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTransitGatewayNetworkFunctionAttachmentMixinProps := &CfnTransitGatewayNetworkFunctionAttachmentMixinProps{
//   	AvailabilityZoneIds: []*string{
//   		jsii.String("availabilityZoneIds"),
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaynetworkfunctionattachment.html
//
type CfnTransitGatewayNetworkFunctionAttachmentMixinProps struct {
	// The Availability Zone IDs for the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaynetworkfunctionattachment.html#cfn-ec2-transitgatewaynetworkfunctionattachment-availabilityzoneids
	//
	AvailabilityZoneIds *[]*string `field:"optional" json:"availabilityZoneIds" yaml:"availabilityZoneIds"`
	// The Availability Zones for the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaynetworkfunctionattachment.html#cfn-ec2-transitgatewaynetworkfunctionattachment-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The name of the network function service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaynetworkfunctionattachment.html#cfn-ec2-transitgatewaynetworkfunctionattachment-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The tags for the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaynetworkfunctionattachment.html#cfn-ec2-transitgatewaynetworkfunctionattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaynetworkfunctionattachment.html#cfn-ec2-transitgatewaynetworkfunctionattachment-transitgatewayid
	//
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
}


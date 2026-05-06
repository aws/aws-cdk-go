package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transitGatewayConfigurationProperty := &TransitGatewayConfigurationProperty{
//   	AvailabilityZoneIds: []*string{
//   		jsii.String("availabilityZoneIds"),
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html
//
type CfnClientVpnEndpointPropsMixin_TransitGatewayConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html#cfn-ec2-clientvpnendpoint-transitgatewayconfiguration-availabilityzoneids
	//
	AvailabilityZoneIds *[]*string `field:"optional" json:"availabilityZoneIds" yaml:"availabilityZoneIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html#cfn-ec2-clientvpnendpoint-transitgatewayconfiguration-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html#cfn-ec2-clientvpnendpoint-transitgatewayconfiguration-transitgatewayid
	//
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
}


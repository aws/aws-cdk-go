package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayConfigurationProperty := &TransitGatewayConfigurationProperty{
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//
//   	// the properties below are optional
//   	AvailabilityZoneIds: []*string{
//   		jsii.String("availabilityZoneIds"),
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html
//
type CfnClientVpnEndpoint_TransitGatewayConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html#cfn-ec2-clientvpnendpoint-transitgatewayconfiguration-transitgatewayid
	//
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html#cfn-ec2-clientvpnendpoint-transitgatewayconfiguration-availabilityzoneids
	//
	AvailabilityZoneIds *[]*string `field:"optional" json:"availabilityZoneIds" yaml:"availabilityZoneIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-transitgatewayconfiguration.html#cfn-ec2-clientvpnendpoint-transitgatewayconfiguration-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
}


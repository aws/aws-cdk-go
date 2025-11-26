package previewawsredshiftserverlessmixins


// Contains information about a network interface in an Amazon Redshift Serverless managed VPC endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkInterfaceProperty := &NetworkInterfaceProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-networkinterface.html
//
type CfnWorkgroupPropsMixin_NetworkInterfaceProperty struct {
	// The availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-networkinterface.html#cfn-redshiftserverless-workgroup-networkinterface-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The unique identifier of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-networkinterface.html#cfn-redshiftserverless-workgroup-networkinterface-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The IPv4 address of the network interface within the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-networkinterface.html#cfn-redshiftserverless-workgroup-networkinterface-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// The unique identifier of the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-networkinterface.html#cfn-redshiftserverless-workgroup-networkinterface-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}


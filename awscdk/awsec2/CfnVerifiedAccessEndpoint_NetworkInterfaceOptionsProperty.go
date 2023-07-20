package awsec2


// Describes the network interface options when creating an AWS Verified Access endpoint using the `network-interface` type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceOptionsProperty := &NetworkInterfaceOptionsProperty{
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-networkinterfaceoptions.html
//
type CfnVerifiedAccessEndpoint_NetworkInterfaceOptionsProperty struct {
	// The ID of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-networkinterfaceoptions.html#cfn-ec2-verifiedaccessendpoint-networkinterfaceoptions-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The IP port number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-networkinterfaceoptions.html#cfn-ec2-verifiedaccessendpoint-networkinterfaceoptions-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The IP protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-networkinterfaceoptions.html#cfn-ec2-verifiedaccessendpoint-networkinterfaceoptions-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


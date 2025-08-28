package awsquicksight


// The structure that contains information about a network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &NetworkInterfaceProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	ErrorMessage: jsii.String("errorMessage"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	Status: jsii.String("status"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-vpcconnection-networkinterface.html
//
type CfnVPCConnection_NetworkInterfaceProperty struct {
	// The availability zone that the network interface resides in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-vpcconnection-networkinterface.html#cfn-quicksight-vpcconnection-networkinterface-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// An error message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-vpcconnection-networkinterface.html#cfn-quicksight-vpcconnection-networkinterface-errormessage
	//
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// The network interface ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-vpcconnection-networkinterface.html#cfn-quicksight-vpcconnection-networkinterface-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The status of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-vpcconnection-networkinterface.html#cfn-quicksight-vpcconnection-networkinterface-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The subnet ID associated with the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-vpcconnection-networkinterface.html#cfn-quicksight-vpcconnection-networkinterface-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}


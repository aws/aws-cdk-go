package awsmedialive


// The properties for a VPC type input destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDestinationVpcProperty := &InputDestinationVpcProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationvpc.html
//
type CfnInput_InputDestinationVpcProperty struct {
	// The availability zone of the Input destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationvpc.html#cfn-medialive-input-inputdestinationvpc-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The network interface ID of the Input destination in the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationvpc.html#cfn-medialive-input-inputdestinationvpc-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}


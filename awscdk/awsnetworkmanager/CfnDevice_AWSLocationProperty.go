package awsnetworkmanager


// The Amazon Web Services location of the device, if applicable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSLocationProperty := &AWSLocationProperty{
//   	SubnetArn: jsii.String("subnetArn"),
//   	Zone: jsii.String("zone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-awslocation.html
//
type CfnDevice_AWSLocationProperty struct {
	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-awslocation.html#cfn-networkmanager-device-awslocation-subnetarn
	//
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// The Zone that the device is located in.
	//
	// Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-awslocation.html#cfn-networkmanager-device-awslocation-zone
	//
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}


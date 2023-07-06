package awsnetworkmanager


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
type CfnDevice_AWSLocationProperty struct {
	// `CfnDevice.AWSLocationProperty.SubnetArn`.
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// `CfnDevice.AWSLocationProperty.Zone`.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}


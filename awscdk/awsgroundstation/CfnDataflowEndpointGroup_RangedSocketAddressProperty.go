package awsgroundstation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rangedSocketAddressProperty := &RangedSocketAddressProperty{
//   	Name: jsii.String("name"),
//   	PortRange: &IntegerRangeProperty{
//   		Maximum: jsii.Number(123),
//   		Minimum: jsii.Number(123),
//   	},
//   }
//
type CfnDataflowEndpointGroup_RangedSocketAddressProperty struct {
	// `CfnDataflowEndpointGroup.RangedSocketAddressProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnDataflowEndpointGroup.RangedSocketAddressProperty.PortRange`.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
}


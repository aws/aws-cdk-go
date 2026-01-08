package previewawsgroundstationmixins


// A socket address with a port range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rangedSocketAddressProperty := &RangedSocketAddressProperty{
//   	Name: jsii.String("name"),
//   	PortRange: &IntegerRangeProperty{
//   		Maximum: jsii.Number(123),
//   		Minimum: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-rangedsocketaddress.html
//
type CfnDataflowEndpointGroupV2PropsMixin_RangedSocketAddressProperty struct {
	// IPv4 socket address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-rangedsocketaddress.html#cfn-groundstation-dataflowendpointgroupv2-rangedsocketaddress-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Port range of a socket address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-rangedsocketaddress.html#cfn-groundstation-dataflowendpointgroupv2-rangedsocketaddress-portrange
	//
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
}


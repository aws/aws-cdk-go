package previewawsgroundstationmixins


// An integer range that has a minimum and maximum value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   integerRangeProperty := &IntegerRangeProperty{
//   	Maximum: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-integerrange.html
//
type CfnDataflowEndpointGroupV2PropsMixin_IntegerRangeProperty struct {
	// A maximum value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-integerrange.html#cfn-groundstation-dataflowendpointgroupv2-integerrange-maximum
	//
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// A minimum value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroupv2-integerrange.html#cfn-groundstation-dataflowendpointgroupv2-integerrange-minimum
	//
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}


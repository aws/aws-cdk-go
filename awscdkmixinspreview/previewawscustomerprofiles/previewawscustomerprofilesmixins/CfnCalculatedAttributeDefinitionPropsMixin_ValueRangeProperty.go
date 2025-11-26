package previewawscustomerprofilesmixins


// A structure letting customers specify a relative time window over which over which data is included in the Calculated Attribute.
//
// Use positive numbers to indicate that the endpoint is in the past, and negative numbers to indicate it is in the future. ValueRange overrides Value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   valueRangeProperty := &ValueRangeProperty{
//   	End: jsii.Number(123),
//   	Start: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-valuerange.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_ValueRangeProperty struct {
	// The ending point for this overridden range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-valuerange.html#cfn-customerprofiles-calculatedattributedefinition-valuerange-end
	//
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// The starting point for this overridden range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-valuerange.html#cfn-customerprofiles-calculatedattributedefinition-valuerange-start
	//
	Start *float64 `field:"optional" json:"start" yaml:"start"`
}


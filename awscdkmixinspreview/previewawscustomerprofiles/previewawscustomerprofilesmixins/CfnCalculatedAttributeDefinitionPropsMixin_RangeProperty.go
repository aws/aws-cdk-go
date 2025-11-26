package previewawscustomerprofilesmixins


// The relative time period over which data is included in the aggregation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rangeProperty := &RangeProperty{
//   	TimestampFormat: jsii.String("timestampFormat"),
//   	TimestampSource: jsii.String("timestampSource"),
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   	ValueRange: &ValueRangeProperty{
//   		End: jsii.Number(123),
//   		Start: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_RangeProperty struct {
	// The format the timestamp field in your JSON object is specified.
	//
	// This value should be one of EPOCHMILLI (for Unix epoch timestamps with second/millisecond level precision) or ISO_8601 (following ISO_8601 format with second/millisecond level precision, with an optional offset of Z or in the format HH:MM or HHMM.). E.g. if your object type is MyType and source JSON is {"generatedAt": {"timestamp": "2001-07-04T12:08:56.235-0700"}}, then TimestampFormat should be "ISO_8601"
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-timestampformat
	//
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// An expression specifying the field in your JSON object from which the date should be parsed.
	//
	// The expression should follow the structure of \"{ObjectTypeName.<Location of timestamp field in JSON pointer format>}\". E.g. if your object type is MyType and source JSON is {"generatedAt": {"timestamp": "1737587945945"}}, then TimestampSource should be "{MyType.generatedAt.timestamp}"
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-timestampsource
	//
	TimestampSource *string `field:"optional" json:"timestampSource" yaml:"timestampSource"`
	// The unit of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The amount of time of the specified unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
	// A structure letting customers specify a relative time window over which over which data is included in the Calculated Attribute.
	//
	// Use positive numbers to indicate that the endpoint is in the past, and negative numbers to indicate it is in the future. ValueRange overrides Value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-valuerange
	//
	ValueRange interface{} `field:"optional" json:"valueRange" yaml:"valueRange"`
}


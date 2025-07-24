package awscustomerprofiles


// The relative time period over which data is included in the aggregation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rangeProperty := &RangeProperty{
//   	Unit: jsii.String("unit"),
//
//   	// the properties below are optional
//   	TimestampFormat: jsii.String("timestampFormat"),
//   	TimestampSource: jsii.String("timestampSource"),
//   	Value: jsii.Number(123),
//   	ValueRange: &ValueRangeProperty{
//   		End: jsii.Number(123),
//   		Start: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html
//
type CfnCalculatedAttributeDefinition_RangeProperty struct {
	// The unit of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// The format the timestamp field in your JSON object is specified.
	//
	// This value should be one of EPOCHMILLI or ISO_8601. E.g. if your object type is MyType and source JSON is {"generatedAt": {"timestamp": "2001-07-04T12:08:56.235Z"}}, then TimestampFormat should be "ISO_8601".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-timestampformat
	//
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// An expression specifying the field in your JSON object from which the date should be parsed.
	//
	// The expression should follow the structure of \"{ObjectTypeName.<Location of timestamp field in JSON pointer format>}\". E.g. if your object type is MyType and source JSON is {"generatedAt": {"timestamp": "1737587945945"}}, then TimestampSource should be "{MyType.generatedAt.timestamp}".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-timestampsource
	//
	TimestampSource *string `field:"optional" json:"timestampSource" yaml:"timestampSource"`
	// The amount of time of the specified unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
	// A structure specifying the endpoints of the relative time period over which data is included in the aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-valuerange
	//
	ValueRange interface{} `field:"optional" json:"valueRange" yaml:"valueRange"`
}


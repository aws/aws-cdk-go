package mixinsawscustomerprofiles


// The conditions including range, object count, and threshold for the calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionsProperty := &ConditionsProperty{
//   	ObjectCount: jsii.Number(123),
//   	Range: &RangeProperty{
//   		TimestampFormat: jsii.String("timestampFormat"),
//   		TimestampSource: jsii.String("timestampSource"),
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   		ValueRange: &ValueRangeProperty{
//   			End: jsii.Number(123),
//   			Start: jsii.Number(123),
//   		},
//   	},
//   	Threshold: &ThresholdProperty{
//   		Operator: jsii.String("operator"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-conditions.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_ConditionsProperty struct {
	// The number of profile objects used for the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-conditions.html#cfn-customerprofiles-calculatedattributedefinition-conditions-objectcount
	//
	ObjectCount *float64 `field:"optional" json:"objectCount" yaml:"objectCount"`
	// The relative time period over which data is included in the aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-conditions.html#cfn-customerprofiles-calculatedattributedefinition-conditions-range
	//
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The threshold for the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-conditions.html#cfn-customerprofiles-calculatedattributedefinition-conditions-threshold
	//
	Threshold interface{} `field:"optional" json:"threshold" yaml:"threshold"`
}


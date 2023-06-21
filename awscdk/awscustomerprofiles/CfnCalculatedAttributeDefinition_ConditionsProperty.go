package awscustomerprofiles


// The conditions including range, object count, and threshold for the calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionsProperty := &ConditionsProperty{
//   	ObjectCount: jsii.Number(123),
//   	Range: &RangeProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	Threshold: &ThresholdProperty{
//   		Operator: jsii.String("operator"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
type CfnCalculatedAttributeDefinition_ConditionsProperty struct {
	// The number of profile objects used for the calculated attribute.
	ObjectCount *float64 `field:"optional" json:"objectCount" yaml:"objectCount"`
	// The relative time period over which data is included in the aggregation.
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The threshold for the calculated attribute.
	Threshold interface{} `field:"optional" json:"threshold" yaml:"threshold"`
}


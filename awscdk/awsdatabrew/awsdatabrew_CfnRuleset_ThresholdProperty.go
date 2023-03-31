package awsdatabrew


// The threshold used with a non-aggregate check expression.
//
// The non-aggregate check expression will be applied to each row in a specific column. Then the threshold will be used to determine whether the validation succeeds.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thresholdProperty := &thresholdProperty{
//   	value: jsii.Number(123),
//
//   	// the properties below are optional
//   	type: jsii.String("type"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnRuleset_ThresholdProperty struct {
	// The value of a threshold.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// The type of a threshold.
	//
	// Used for comparison of an actual count of rows that satisfy the rule to the threshold value.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Unit of threshold value.
	//
	// Can be either a COUNT or PERCENTAGE of the full sample size used for validation.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}


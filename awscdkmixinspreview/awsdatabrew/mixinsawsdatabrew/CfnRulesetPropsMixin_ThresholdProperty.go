package mixinsawsdatabrew


// The threshold used with a non-aggregate check expression.
//
// The non-aggregate check expression will be applied to each row in a specific column. Then the threshold will be used to determine whether the validation succeeds.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   thresholdProperty := &ThresholdProperty{
//   	Type: jsii.String("type"),
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-threshold.html
//
type CfnRulesetPropsMixin_ThresholdProperty struct {
	// The type of a threshold.
	//
	// Used for comparison of an actual count of rows that satisfy the rule to the threshold value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-threshold.html#cfn-databrew-ruleset-threshold-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Unit of threshold value.
	//
	// Can be either a COUNT or PERCENTAGE of the full sample size used for validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-threshold.html#cfn-databrew-ruleset-threshold-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The value of a threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-threshold.html#cfn-databrew-ruleset-threshold-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


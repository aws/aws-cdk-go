package awsbedrockagentcore


// The filter that applies conditions to agent traces during online evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Key: jsii.String("key"),
//   	Operator: jsii.String("operator"),
//   	Value: &FilterValueProperty{
//   		BooleanValue: jsii.Boolean(false),
//   		DoubleValue: jsii.Number(123),
//   		StringValue: jsii.String("stringValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-filter.html
//
type CfnOnlineEvaluationConfig_FilterProperty struct {
	// The key or field name to filter on within the agent trace data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-filter.html#cfn-bedrockagentcore-onlineevaluationconfig-filter-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The comparison operator to use for filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-filter.html#cfn-bedrockagentcore-onlineevaluationconfig-filter-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value used in filter comparisons.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-filter.html#cfn-bedrockagentcore-onlineevaluationconfig-filter-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}


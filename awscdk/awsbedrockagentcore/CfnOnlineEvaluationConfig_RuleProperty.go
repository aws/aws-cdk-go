package awsbedrockagentcore


// The evaluation rule that defines sampling configuration, filtering criteria, and session detection settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	SamplingConfig: &SamplingConfigProperty{
//   		SamplingPercentage: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			Key: jsii.String("key"),
//   			Operator: jsii.String("operator"),
//   			Value: &FilterValueProperty{
//   				BooleanValue: jsii.Boolean(false),
//   				DoubleValue: jsii.Number(123),
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	SessionConfig: &SessionConfigProperty{
//   		SessionTimeoutMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-rule.html
//
type CfnOnlineEvaluationConfig_RuleProperty struct {
	// The configuration that controls what percentage of agent traces are sampled for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-rule.html#cfn-bedrockagentcore-onlineevaluationconfig-rule-samplingconfig
	//
	SamplingConfig interface{} `field:"required" json:"samplingConfig" yaml:"samplingConfig"`
	// The list of filters that determine which agent traces should be included in the evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-rule.html#cfn-bedrockagentcore-onlineevaluationconfig-rule-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The configuration that defines how agent sessions are detected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-rule.html#cfn-bedrockagentcore-onlineevaluationconfig-rule-sessionconfig
	//
	SessionConfig interface{} `field:"optional" json:"sessionConfig" yaml:"sessionConfig"`
}


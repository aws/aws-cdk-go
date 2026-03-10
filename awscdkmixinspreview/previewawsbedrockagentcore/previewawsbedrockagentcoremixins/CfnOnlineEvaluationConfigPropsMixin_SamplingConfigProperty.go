package previewawsbedrockagentcoremixins


// The configuration that controls what percentage of agent traces are sampled for evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   samplingConfigProperty := &SamplingConfigProperty{
//   	SamplingPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-samplingconfig.html
//
type CfnOnlineEvaluationConfigPropsMixin_SamplingConfigProperty struct {
	// The percentage of agent traces to sample for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-samplingconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-samplingconfig-samplingpercentage
	//
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
}


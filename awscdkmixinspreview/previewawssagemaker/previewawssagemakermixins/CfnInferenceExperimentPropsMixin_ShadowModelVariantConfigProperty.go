package previewawssagemakermixins


// The name and sampling percentage of a shadow variant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   shadowModelVariantConfigProperty := &ShadowModelVariantConfigProperty{
//   	SamplingPercentage: jsii.Number(123),
//   	ShadowModelVariantName: jsii.String("shadowModelVariantName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodelvariantconfig.html
//
type CfnInferenceExperimentPropsMixin_ShadowModelVariantConfigProperty struct {
	// The percentage of inference requests that Amazon SageMaker replicates from the production variant to the shadow variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodelvariantconfig.html#cfn-sagemaker-inferenceexperiment-shadowmodelvariantconfig-samplingpercentage
	//
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
	// The name of the shadow variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodelvariantconfig.html#cfn-sagemaker-inferenceexperiment-shadowmodelvariantconfig-shadowmodelvariantname
	//
	ShadowModelVariantName *string `field:"optional" json:"shadowModelVariantName" yaml:"shadowModelVariantName"`
}


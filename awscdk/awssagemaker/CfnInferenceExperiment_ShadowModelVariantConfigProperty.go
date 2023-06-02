package awssagemaker


// The name and sampling percentage of a shadow variant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shadowModelVariantConfigProperty := &ShadowModelVariantConfigProperty{
//   	SamplingPercentage: jsii.Number(123),
//   	ShadowModelVariantName: jsii.String("shadowModelVariantName"),
//   }
//
type CfnInferenceExperiment_ShadowModelVariantConfigProperty struct {
	// The percentage of inference requests that Amazon SageMaker replicates from the production variant to the shadow variant.
	SamplingPercentage *float64 `field:"required" json:"samplingPercentage" yaml:"samplingPercentage"`
	// The name of the shadow variant.
	ShadowModelVariantName *string `field:"required" json:"shadowModelVariantName" yaml:"shadowModelVariantName"`
}


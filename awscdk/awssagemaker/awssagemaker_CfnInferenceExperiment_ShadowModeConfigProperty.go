package awssagemaker


// The configuration of `ShadowMode` inference experiment type, which specifies a production variant to take all the inference requests, and a shadow variant to which Amazon SageMaker replicates a percentage of the inference requests.
//
// For the shadow variant it also specifies the percentage of requests that Amazon SageMaker replicates.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shadowModeConfigProperty := &shadowModeConfigProperty{
//   	shadowModelVariants: []interface{}{
//   		&shadowModelVariantConfigProperty{
//   			samplingPercentage: jsii.Number(123),
//   			shadowModelVariantName: jsii.String("shadowModelVariantName"),
//   		},
//   	},
//   	sourceModelVariantName: jsii.String("sourceModelVariantName"),
//   }
//
type CfnInferenceExperiment_ShadowModeConfigProperty struct {
	// List of shadow variant configurations.
	ShadowModelVariants interface{} `field:"required" json:"shadowModelVariants" yaml:"shadowModelVariants"`
	// The name of the production variant, which takes all the inference requests.
	SourceModelVariantName *string `field:"required" json:"sourceModelVariantName" yaml:"sourceModelVariantName"`
}


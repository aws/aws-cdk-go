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
//   shadowModeConfigProperty := &ShadowModeConfigProperty{
//   	ShadowModelVariants: []interface{}{
//   		&ShadowModelVariantConfigProperty{
//   			SamplingPercentage: jsii.Number(123),
//   			ShadowModelVariantName: jsii.String("shadowModelVariantName"),
//   		},
//   	},
//   	SourceModelVariantName: jsii.String("sourceModelVariantName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodeconfig.html
//
type CfnInferenceExperiment_ShadowModeConfigProperty struct {
	// List of shadow variant configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodeconfig.html#cfn-sagemaker-inferenceexperiment-shadowmodeconfig-shadowmodelvariants
	//
	ShadowModelVariants interface{} `field:"required" json:"shadowModelVariants" yaml:"shadowModelVariants"`
	// The name of the production variant, which takes all the inference requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodeconfig.html#cfn-sagemaker-inferenceexperiment-shadowmodeconfig-sourcemodelvariantname
	//
	SourceModelVariantName *string `field:"required" json:"sourceModelVariantName" yaml:"sourceModelVariantName"`
}


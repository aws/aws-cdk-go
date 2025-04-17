package awssagemaker


// The configuration for SHAP analysis using SageMaker Clarify Explainer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyShapConfigProperty := &ClarifyShapConfigProperty{
//   	ShapBaselineConfig: &ClarifyShapBaselineConfigProperty{
//   		MimeType: jsii.String("mimeType"),
//   		ShapBaseline: jsii.String("shapBaseline"),
//   		ShapBaselineUri: jsii.String("shapBaselineUri"),
//   	},
//
//   	// the properties below are optional
//   	NumberOfSamples: jsii.Number(123),
//   	Seed: jsii.Number(123),
//   	TextConfig: &ClarifyTextConfigProperty{
//   		Granularity: jsii.String("granularity"),
//   		Language: jsii.String("language"),
//   	},
//   	UseLogit: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html
//
type CfnEndpointConfig_ClarifyShapConfigProperty struct {
	// The configuration for the SHAP baseline of the Kernal SHAP algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-shapbaselineconfig
	//
	ShapBaselineConfig interface{} `field:"required" json:"shapBaselineConfig" yaml:"shapBaselineConfig"`
	// The number of samples to be used for analysis by the Kernal SHAP algorithm.
	//
	// > The number of samples determines the size of the synthetic dataset, which has an impact on latency of explainability requests. For more information, see the *Synthetic data* of [Configure and create an endpoint](https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-online-explainability-create-endpoint.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-numberofsamples
	//
	NumberOfSamples *float64 `field:"optional" json:"numberOfSamples" yaml:"numberOfSamples"`
	// The starting value used to initialize the random number generator in the explainer.
	//
	// Provide a value for this parameter to obtain a deterministic SHAP result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-seed
	//
	Seed *float64 `field:"optional" json:"seed" yaml:"seed"`
	// A parameter that indicates if text features are treated as text and explanations are provided for individual units of text.
	//
	// Required for natural language processing (NLP) explainability only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-textconfig
	//
	TextConfig interface{} `field:"optional" json:"textConfig" yaml:"textConfig"`
	// A Boolean toggle to indicate if you want to use the logit function (true) or log-odds units (false) for model predictions.
	//
	// Defaults to false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyshapconfig.html#cfn-sagemaker-endpointconfig-clarifyshapconfig-uselogit
	//
	UseLogit interface{} `field:"optional" json:"useLogit" yaml:"useLogit"`
}


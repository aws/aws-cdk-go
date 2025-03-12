package awssagemaker


// The configuration parameters for the SageMaker Clarify explainer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyExplainerConfigProperty := &ClarifyExplainerConfigProperty{
//   	ShapConfig: &ClarifyShapConfigProperty{
//   		ShapBaselineConfig: &ClarifyShapBaselineConfigProperty{
//   			MimeType: jsii.String("mimeType"),
//   			ShapBaseline: jsii.String("shapBaseline"),
//   			ShapBaselineUri: jsii.String("shapBaselineUri"),
//   		},
//
//   		// the properties below are optional
//   		NumberOfSamples: jsii.Number(123),
//   		Seed: jsii.Number(123),
//   		TextConfig: &ClarifyTextConfigProperty{
//   			Granularity: jsii.String("granularity"),
//   			Language: jsii.String("language"),
//   		},
//   		UseLogit: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	EnableExplanations: jsii.String("enableExplanations"),
//   	InferenceConfig: &ClarifyInferenceConfigProperty{
//   		ContentTemplate: jsii.String("contentTemplate"),
//   		FeatureHeaders: []*string{
//   			jsii.String("featureHeaders"),
//   		},
//   		FeaturesAttribute: jsii.String("featuresAttribute"),
//   		FeatureTypes: []*string{
//   			jsii.String("featureTypes"),
//   		},
//   		LabelAttribute: jsii.String("labelAttribute"),
//   		LabelHeaders: []*string{
//   			jsii.String("labelHeaders"),
//   		},
//   		LabelIndex: jsii.Number(123),
//   		MaxPayloadInMb: jsii.Number(123),
//   		MaxRecordCount: jsii.Number(123),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		ProbabilityIndex: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyexplainerconfig.html
//
type CfnEndpointConfig_ClarifyExplainerConfigProperty struct {
	// The configuration for SHAP analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyexplainerconfig.html#cfn-sagemaker-endpointconfig-clarifyexplainerconfig-shapconfig
	//
	ShapConfig interface{} `field:"required" json:"shapConfig" yaml:"shapConfig"`
	// A JMESPath boolean expression used to filter which records to explain.
	//
	// Explanations are activated by default. See [`EnableExplanations`](https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-online-explainability-create-endpoint.html#clarify-online-explainability-create-endpoint-enable) for additional information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyexplainerconfig.html#cfn-sagemaker-endpointconfig-clarifyexplainerconfig-enableexplanations
	//
	EnableExplanations *string `field:"optional" json:"enableExplanations" yaml:"enableExplanations"`
	// The inference configuration parameter for the model container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifyexplainerconfig.html#cfn-sagemaker-endpointconfig-clarifyexplainerconfig-inferenceconfig
	//
	InferenceConfig interface{} `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
}


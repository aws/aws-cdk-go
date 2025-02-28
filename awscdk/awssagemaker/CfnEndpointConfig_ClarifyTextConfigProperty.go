package awssagemaker


// A parameter used to configure the SageMaker Clarify explainer to treat text features as text so that explanations are provided for individual units of text.
//
// Required only for natural language processing (NLP) explainability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyTextConfigProperty := &ClarifyTextConfigProperty{
//   	Granularity: jsii.String("granularity"),
//   	Language: jsii.String("language"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifytextconfig.html
//
type CfnEndpointConfig_ClarifyTextConfigProperty struct {
	// The unit of granularity for the analysis of text features.
	//
	// For example, if the unit is `'token'` , then each token (like a word in English) of the text is treated as a feature. SHAP values are computed for each unit/feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifytextconfig.html#cfn-sagemaker-endpointconfig-clarifytextconfig-granularity
	//
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// Specifies the language of the text features in [ISO 639-1](https://docs.aws.amazon.com/ https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) or [ISO 639-3](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_639-3) code of a supported language.
	//
	// > For a mix of multiple languages, use code `'xx'` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifytextconfig.html#cfn-sagemaker-endpointconfig-clarifytextconfig-language
	//
	Language *string `field:"required" json:"language" yaml:"language"`
}


package previewawsdatazonemixins


// The recommendation configuration for the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recommendationConfigurationProperty := &RecommendationConfigurationProperty{
//   	EnableBusinessNameGeneration: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-recommendationconfiguration.html
//
type CfnDataSourcePropsMixin_RecommendationConfigurationProperty struct {
	// Specifies whether automatic business name generation is to be enabled or not as part of the recommendation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-recommendationconfiguration.html#cfn-datazone-datasource-recommendationconfiguration-enablebusinessnamegeneration
	//
	EnableBusinessNameGeneration interface{} `field:"optional" json:"enableBusinessNameGeneration" yaml:"enableBusinessNameGeneration"`
}


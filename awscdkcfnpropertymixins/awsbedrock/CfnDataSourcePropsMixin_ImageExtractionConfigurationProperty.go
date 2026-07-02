package awsbedrock


// Configuration for image extraction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   imageExtractionConfigurationProperty := &ImageExtractionConfigurationProperty{
//   	ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-imageextractionconfiguration.html
//
type CfnDataSourcePropsMixin_ImageExtractionConfigurationProperty struct {
	// Indicates whether a feature is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-imageextractionconfiguration.html#cfn-bedrock-datasource-imageextractionconfiguration-imageextractionstatus
	//
	ImageExtractionStatus *string `field:"optional" json:"imageExtractionStatus" yaml:"imageExtractionStatus"`
}


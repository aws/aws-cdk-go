package awsbedrock


// Configuration for video extraction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   videoExtractionConfigurationProperty := &VideoExtractionConfigurationProperty{
//   	VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-videoextractionconfiguration.html
//
type CfnDataSourcePropsMixin_VideoExtractionConfigurationProperty struct {
	// Indicates whether a feature is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-videoextractionconfiguration.html#cfn-bedrock-datasource-videoextractionconfiguration-videoextractionstatus
	//
	VideoExtractionStatus *string `field:"optional" json:"videoExtractionStatus" yaml:"videoExtractionStatus"`
}


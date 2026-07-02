package awsbedrock


// Configuration for media extraction settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mediaExtractionConfigurationProperty := &MediaExtractionConfigurationProperty{
//   	AudioExtractionConfiguration: &AudioExtractionConfigurationProperty{
//   		AudioExtractionStatus: jsii.String("audioExtractionStatus"),
//   	},
//   	ImageExtractionConfiguration: &ImageExtractionConfigurationProperty{
//   		ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   	},
//   	VideoExtractionConfiguration: &VideoExtractionConfigurationProperty{
//   		VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-mediaextractionconfiguration.html
//
type CfnDataSourcePropsMixin_MediaExtractionConfigurationProperty struct {
	// Configuration for audio extraction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-mediaextractionconfiguration.html#cfn-bedrock-datasource-mediaextractionconfiguration-audioextractionconfiguration
	//
	AudioExtractionConfiguration interface{} `field:"optional" json:"audioExtractionConfiguration" yaml:"audioExtractionConfiguration"`
	// Configuration for image extraction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-mediaextractionconfiguration.html#cfn-bedrock-datasource-mediaextractionconfiguration-imageextractionconfiguration
	//
	ImageExtractionConfiguration interface{} `field:"optional" json:"imageExtractionConfiguration" yaml:"imageExtractionConfiguration"`
	// Configuration for video extraction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-mediaextractionconfiguration.html#cfn-bedrock-datasource-mediaextractionconfiguration-videoextractionconfiguration
	//
	VideoExtractionConfiguration interface{} `field:"optional" json:"videoExtractionConfiguration" yaml:"videoExtractionConfiguration"`
}


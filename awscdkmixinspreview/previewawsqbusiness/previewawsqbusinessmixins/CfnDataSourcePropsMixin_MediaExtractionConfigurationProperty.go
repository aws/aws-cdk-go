package previewawsqbusinessmixins


// The configuration for extracting information from media in documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html
//
type CfnDataSourcePropsMixin_MediaExtractionConfigurationProperty struct {
	// Configuration settings for extracting and processing audio content from media files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html#cfn-qbusiness-datasource-mediaextractionconfiguration-audioextractionconfiguration
	//
	AudioExtractionConfiguration interface{} `field:"optional" json:"audioExtractionConfiguration" yaml:"audioExtractionConfiguration"`
	// The configuration for extracting semantic meaning from images in documents.
	//
	// For more information, see [Extracting semantic meaning from images and visuals](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/extracting-meaning-from-images.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html#cfn-qbusiness-datasource-mediaextractionconfiguration-imageextractionconfiguration
	//
	ImageExtractionConfiguration interface{} `field:"optional" json:"imageExtractionConfiguration" yaml:"imageExtractionConfiguration"`
	// Configuration settings for extracting and processing video content from media files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html#cfn-qbusiness-datasource-mediaextractionconfiguration-videoextractionconfiguration
	//
	VideoExtractionConfiguration interface{} `field:"optional" json:"videoExtractionConfiguration" yaml:"videoExtractionConfiguration"`
}


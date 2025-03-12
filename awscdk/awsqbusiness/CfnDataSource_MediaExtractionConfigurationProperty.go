package awsqbusiness


// The configuration for extracting information from media in documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaExtractionConfigurationProperty := &MediaExtractionConfigurationProperty{
//   	ImageExtractionConfiguration: &ImageExtractionConfigurationProperty{
//   		ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html
//
type CfnDataSource_MediaExtractionConfigurationProperty struct {
	// The configuration for extracting semantic meaning from images in documents.
	//
	// For more information, see [Extracting semantic meaning from images and visuals](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/extracting-meaning-from-images.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html#cfn-qbusiness-datasource-mediaextractionconfiguration-imageextractionconfiguration
	//
	ImageExtractionConfiguration interface{} `field:"optional" json:"imageExtractionConfiguration" yaml:"imageExtractionConfiguration"`
}


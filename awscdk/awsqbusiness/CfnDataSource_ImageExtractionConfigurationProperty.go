package awsqbusiness


// The configuration for extracting semantic meaning from images in documents.
//
// For more information, see [Extracting semantic meaning from images and visuals](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/extracting-meaning-from-images.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageExtractionConfigurationProperty := &ImageExtractionConfigurationProperty{
//   	ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-imageextractionconfiguration.html
//
type CfnDataSource_ImageExtractionConfigurationProperty struct {
	// Specify whether to extract semantic meaning from images and visuals from documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-imageextractionconfiguration.html#cfn-qbusiness-datasource-imageextractionconfiguration-imageextractionstatus
	//
	ImageExtractionStatus *string `field:"required" json:"imageExtractionStatus" yaml:"imageExtractionStatus"`
}


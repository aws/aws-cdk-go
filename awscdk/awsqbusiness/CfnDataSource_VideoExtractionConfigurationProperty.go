package awsqbusiness


// Configuration settings for video content extraction and processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoExtractionConfigurationProperty := &VideoExtractionConfigurationProperty{
//   	VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-videoextractionconfiguration.html
//
type CfnDataSource_VideoExtractionConfigurationProperty struct {
	// The status of video extraction (ENABLED or DISABLED) for processing video content from files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-videoextractionconfiguration.html#cfn-qbusiness-datasource-videoextractionconfiguration-videoextractionstatus
	//
	VideoExtractionStatus *string `field:"required" json:"videoExtractionStatus" yaml:"videoExtractionStatus"`
}


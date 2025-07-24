package awsqbusiness


// Configuration settings for audio content extraction and processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioExtractionConfigurationProperty := &AudioExtractionConfigurationProperty{
//   	AudioExtractionStatus: jsii.String("audioExtractionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-audioextractionconfiguration.html
//
type CfnDataSource_AudioExtractionConfigurationProperty struct {
	// The status of audio extraction (ENABLED or DISABLED) for processing audio content from files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-audioextractionconfiguration.html#cfn-qbusiness-datasource-audioextractionconfiguration-audioextractionstatus
	//
	AudioExtractionStatus *string `field:"required" json:"audioExtractionStatus" yaml:"audioExtractionStatus"`
}


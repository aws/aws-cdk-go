package awsbedrock


// Configuration for audio extraction.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-audioextractionconfiguration.html
//
type CfnDataSource_AudioExtractionConfigurationProperty struct {
	// Indicates whether a feature is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-audioextractionconfiguration.html#cfn-bedrock-datasource-audioextractionconfiguration-audioextractionstatus
	//
	AudioExtractionStatus *string `field:"required" json:"audioExtractionStatus" yaml:"audioExtractionStatus"`
}


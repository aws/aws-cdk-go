package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioExtractionConfigurationProperty := &AudioExtractionConfigurationProperty{
//   	AudioExtractionStatus: jsii.String("audioExtractionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-audioextractionconfiguration.html
//
type CfnKnowledgeBase_AudioExtractionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-audioextractionconfiguration.html#cfn-quicksight-knowledgebase-audioextractionconfiguration-audioextractionstatus
	//
	AudioExtractionStatus *string `field:"required" json:"audioExtractionStatus" yaml:"audioExtractionStatus"`
}


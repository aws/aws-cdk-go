package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoExtractionConfigurationProperty := &VideoExtractionConfigurationProperty{
//   	VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//
//   	// the properties below are optional
//   	VideoExtractionType: jsii.String("videoExtractionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-videoextractionconfiguration.html
//
type CfnKnowledgeBase_VideoExtractionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-videoextractionconfiguration.html#cfn-quicksight-knowledgebase-videoextractionconfiguration-videoextractionstatus
	//
	VideoExtractionStatus *string `field:"required" json:"videoExtractionStatus" yaml:"videoExtractionStatus"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-videoextractionconfiguration.html#cfn-quicksight-knowledgebase-videoextractionconfiguration-videoextractiontype
	//
	VideoExtractionType *string `field:"optional" json:"videoExtractionType" yaml:"videoExtractionType"`
}


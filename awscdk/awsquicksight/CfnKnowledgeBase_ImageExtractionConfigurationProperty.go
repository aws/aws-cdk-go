package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageExtractionConfigurationProperty := &ImageExtractionConfigurationProperty{
//   	ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-imageextractionconfiguration.html
//
type CfnKnowledgeBase_ImageExtractionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-imageextractionconfiguration.html#cfn-quicksight-knowledgebase-imageextractionconfiguration-imageextractionstatus
	//
	ImageExtractionStatus *string `field:"required" json:"imageExtractionStatus" yaml:"imageExtractionStatus"`
}


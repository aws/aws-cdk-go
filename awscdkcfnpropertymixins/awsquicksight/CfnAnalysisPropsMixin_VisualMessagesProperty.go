package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   visualMessagesProperty := &VisualMessagesProperty{
//   	NoDataMessage: &VisualMessageConfigurationProperty{
//   		Description: jsii.String("description"),
//   		DescriptionVisibility: jsii.String("descriptionVisibility"),
//   		Enabled: jsii.Boolean(false),
//   		LinkText: jsii.String("linkText"),
//   		LinkUrl: jsii.String("linkUrl"),
//   		LinkVisibility: jsii.String("linkVisibility"),
//   		Title: jsii.String("title"),
//   		TitleVisibility: jsii.String("titleVisibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessages.html
//
type CfnAnalysisPropsMixin_VisualMessagesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessages.html#cfn-quicksight-analysis-visualmessages-nodatamessage
	//
	NoDataMessage interface{} `field:"optional" json:"noDataMessage" yaml:"noDataMessage"`
}


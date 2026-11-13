package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   visualMessageConfigurationProperty := &VisualMessageConfigurationProperty{
//   	Description: jsii.String("description"),
//   	DescriptionVisibility: jsii.String("descriptionVisibility"),
//   	Enabled: jsii.Boolean(false),
//   	LinkText: jsii.String("linkText"),
//   	LinkUrl: jsii.String("linkUrl"),
//   	LinkVisibility: jsii.String("linkVisibility"),
//   	Title: jsii.String("title"),
//   	TitleVisibility: jsii.String("titleVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html
//
type CfnAnalysisPropsMixin_VisualMessageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-descriptionvisibility
	//
	DescriptionVisibility *string `field:"optional" json:"descriptionVisibility" yaml:"descriptionVisibility"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-linktext
	//
	LinkText *string `field:"optional" json:"linkText" yaml:"linkText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-linkurl
	//
	LinkUrl *string `field:"optional" json:"linkUrl" yaml:"linkUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-linkvisibility
	//
	LinkVisibility *string `field:"optional" json:"linkVisibility" yaml:"linkVisibility"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-visualmessageconfiguration.html#cfn-quicksight-analysis-visualmessageconfiguration-titlevisibility
	//
	TitleVisibility *string `field:"optional" json:"titleVisibility" yaml:"titleVisibility"`
}


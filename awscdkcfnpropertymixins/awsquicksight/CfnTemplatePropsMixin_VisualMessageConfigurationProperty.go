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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html
//
type CfnTemplatePropsMixin_VisualMessageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-descriptionvisibility
	//
	DescriptionVisibility *string `field:"optional" json:"descriptionVisibility" yaml:"descriptionVisibility"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-linktext
	//
	LinkText *string `field:"optional" json:"linkText" yaml:"linkText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-linkurl
	//
	LinkUrl *string `field:"optional" json:"linkUrl" yaml:"linkUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-linkvisibility
	//
	LinkVisibility *string `field:"optional" json:"linkVisibility" yaml:"linkVisibility"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualmessageconfiguration.html#cfn-quicksight-template-visualmessageconfiguration-titlevisibility
	//
	TitleVisibility *string `field:"optional" json:"titleVisibility" yaml:"titleVisibility"`
}


package previewawsquicksightmixins


// The key value pair of the persisted property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pluginVisualPropertyProperty := &PluginVisualPropertyProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualproperty.html
//
type CfnTemplatePropsMixin_PluginVisualPropertyProperty struct {
	// The name of the plugin visual property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualproperty.html#cfn-quicksight-template-pluginvisualproperty-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the plugin visual property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualproperty.html#cfn-quicksight-template-pluginvisualproperty-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


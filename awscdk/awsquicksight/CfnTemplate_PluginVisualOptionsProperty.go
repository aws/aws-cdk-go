package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginVisualOptionsProperty := &PluginVisualOptionsProperty{
//   	VisualProperties: []interface{}{
//   		&PluginVisualPropertyProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualoptions.html
//
type CfnTemplate_PluginVisualOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualoptions.html#cfn-quicksight-template-pluginvisualoptions-visualproperties
	//
	VisualProperties interface{} `field:"optional" json:"visualProperties" yaml:"visualProperties"`
}


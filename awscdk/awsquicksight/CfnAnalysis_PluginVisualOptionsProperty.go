package awsquicksight


// The options and persisted properties for the plugin visual.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualoptions.html
//
type CfnAnalysis_PluginVisualOptionsProperty struct {
	// The persisted properties and their values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualoptions.html#cfn-quicksight-analysis-pluginvisualoptions-visualproperties
	//
	VisualProperties interface{} `field:"optional" json:"visualProperties" yaml:"visualProperties"`
}


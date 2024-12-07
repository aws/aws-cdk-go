package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginVisualItemsLimitConfigurationProperty := &PluginVisualItemsLimitConfigurationProperty{
//   	ItemsLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualitemslimitconfiguration.html
//
type CfnAnalysis_PluginVisualItemsLimitConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualitemslimitconfiguration.html#cfn-quicksight-analysis-pluginvisualitemslimitconfiguration-itemslimit
	//
	ItemsLimit *float64 `field:"optional" json:"itemsLimit" yaml:"itemsLimit"`
}


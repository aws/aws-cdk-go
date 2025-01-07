package awsquicksight


// A query limits configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginVisualItemsLimitConfigurationProperty := &PluginVisualItemsLimitConfigurationProperty{
//   	ItemsLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualitemslimitconfiguration.html
//
type CfnTemplate_PluginVisualItemsLimitConfigurationProperty struct {
	// Determines how many values are be fetched at once.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualitemslimitconfiguration.html#cfn-quicksight-template-pluginvisualitemslimitconfiguration-itemslimit
	//
	ItemsLimit *float64 `field:"optional" json:"itemsLimit" yaml:"itemsLimit"`
}


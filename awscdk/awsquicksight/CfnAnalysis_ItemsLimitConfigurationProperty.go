package awsquicksight


// The limit configuration of the visual display for an axis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   itemsLimitConfigurationProperty := &ItemsLimitConfigurationProperty{
//   	ItemsLimit: jsii.Number(123),
//   	OtherCategories: jsii.String("otherCategories"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-itemslimitconfiguration.html
//
type CfnAnalysis_ItemsLimitConfigurationProperty struct {
	// The limit on how many items of a field are showed in the chart.
	//
	// For example, the number of slices that are displayed in a pie chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-itemslimitconfiguration.html#cfn-quicksight-analysis-itemslimitconfiguration-itemslimit
	//
	ItemsLimit *float64 `field:"optional" json:"itemsLimit" yaml:"itemsLimit"`
	// The `Show other` of an axis in the chart. Choose one of the following options:.
	//
	// - `INCLUDE`
	// - `EXCLUDE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-itemslimitconfiguration.html#cfn-quicksight-analysis-itemslimitconfiguration-othercategories
	//
	OtherCategories *string `field:"optional" json:"otherCategories" yaml:"otherCategories"`
}


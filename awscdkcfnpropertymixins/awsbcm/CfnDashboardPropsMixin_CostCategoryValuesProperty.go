package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   costCategoryValuesProperty := &CostCategoryValuesProperty{
//   	Key: jsii.String("key"),
//   	MatchOptions: []*string{
//   		jsii.String("matchOptions"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costcategoryvalues.html
//
type CfnDashboardPropsMixin_CostCategoryValuesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costcategoryvalues.html#cfn-bcm-dashboard-costcategoryvalues-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costcategoryvalues.html#cfn-bcm-dashboard-costcategoryvalues-matchoptions
	//
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-costcategoryvalues.html#cfn-bcm-dashboard-costcategoryvalues-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


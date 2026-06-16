package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dimensionValuesProperty := &DimensionValuesProperty{
//   	Key: jsii.String("key"),
//   	MatchOptions: []*string{
//   		jsii.String("matchOptions"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-dimensionvalues.html
//
type CfnDashboardPropsMixin_DimensionValuesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-dimensionvalues.html#cfn-bcm-dashboard-dimensionvalues-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-dimensionvalues.html#cfn-bcm-dashboard-dimensionvalues-matchoptions
	//
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-dimensionvalues.html#cfn-bcm-dashboard-dimensionvalues-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


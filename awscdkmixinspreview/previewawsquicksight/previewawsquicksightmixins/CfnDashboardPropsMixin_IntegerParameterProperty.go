package previewawsquicksightmixins


// An integer parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   integerParameterProperty := &IntegerParameterProperty{
//   	Name: jsii.String("name"),
//   	Values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameter.html
//
type CfnDashboardPropsMixin_IntegerParameterProperty struct {
	// The name of the integer parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameter.html#cfn-quicksight-dashboard-integerparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The values for the integer parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameter.html#cfn-quicksight-dashboard-integerparameter-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}


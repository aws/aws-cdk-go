package awsquicksight


// A decimal parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalParameterProperty := &DecimalParameterProperty{
//   	Name: jsii.String("name"),
//   	Values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameter.html
//
type CfnDashboard_DecimalParameterProperty struct {
	// A display name for the decimal parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameter.html#cfn-quicksight-dashboard-decimalparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the decimal parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameter.html#cfn-quicksight-dashboard-decimalparameter-values
	//
	Values interface{} `field:"required" json:"values" yaml:"values"`
}


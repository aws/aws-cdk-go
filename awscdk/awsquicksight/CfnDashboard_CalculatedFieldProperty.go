package awsquicksight


// The calculated field of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedFieldProperty := &CalculatedFieldProperty{
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	Expression: jsii.String("expression"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-calculatedfield.html
//
type CfnDashboard_CalculatedFieldProperty struct {
	// The data set that is used in this calculated field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-calculatedfield.html#cfn-quicksight-dashboard-calculatedfield-datasetidentifier
	//
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The expression of the calculated field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-calculatedfield.html#cfn-quicksight-dashboard-calculatedfield-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the calculated field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-calculatedfield.html#cfn-quicksight-dashboard-calculatedfield-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}


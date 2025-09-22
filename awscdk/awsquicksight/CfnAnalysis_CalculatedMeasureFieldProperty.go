package awsquicksight


// The table calculation measure field for pivot tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedMeasureFieldProperty := &CalculatedMeasureFieldProperty{
//   	Expression: jsii.String("expression"),
//   	FieldId: jsii.String("fieldId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedmeasurefield.html
//
type CfnAnalysis_CalculatedMeasureFieldProperty struct {
	// The expression in the table calculation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedmeasurefield.html#cfn-quicksight-analysis-calculatedmeasurefield-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The custom field ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedmeasurefield.html#cfn-quicksight-analysis-calculatedmeasurefield-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}


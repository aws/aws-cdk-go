package awsquicksight


// The default values of the `DecimalParameterDeclaration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalDefaultValuesProperty := &DecimalDefaultValuesProperty{
//   	DynamicValue: &DynamicDefaultValueProperty{
//   		DefaultValueColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//
//   		// the properties below are optional
//   		GroupNameColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		UserNameColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
//   	StaticValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimaldefaultvalues.html
//
type CfnAnalysis_DecimalDefaultValuesProperty struct {
	// The dynamic value of the `DecimalDefaultValues` .
	//
	// Different defaults are displayed according to users, groups, and values mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimaldefaultvalues.html#cfn-quicksight-analysis-decimaldefaultvalues-dynamicvalue
	//
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// The static values of the `DecimalDefaultValues` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimaldefaultvalues.html#cfn-quicksight-analysis-decimaldefaultvalues-staticvalues
	//
	StaticValues interface{} `field:"optional" json:"staticValues" yaml:"staticValues"`
}


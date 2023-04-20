package awsquicksight


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
type CfnAnalysis_DecimalDefaultValuesProperty struct {
	// `CfnAnalysis.DecimalDefaultValuesProperty.DynamicValue`.
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// `CfnAnalysis.DecimalDefaultValuesProperty.StaticValues`.
	StaticValues interface{} `field:"optional" json:"staticValues" yaml:"staticValues"`
}


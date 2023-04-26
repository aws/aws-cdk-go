package awsquicksight


// The default values of the `IntegerParameterDeclaration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerDefaultValuesProperty := &IntegerDefaultValuesProperty{
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
type CfnDashboard_IntegerDefaultValuesProperty struct {
	// The dynamic value of the `IntegerDefaultValues` .
	//
	// Different defaults are displayed according to users, groups, and values mapping.
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// The static values of the `IntegerDefaultValues` .
	StaticValues interface{} `field:"optional" json:"staticValues" yaml:"staticValues"`
}


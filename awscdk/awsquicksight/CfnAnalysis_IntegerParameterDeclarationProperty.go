package awsquicksight


// A parameter declaration for the `Integer` data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerParameterDeclarationProperty := &IntegerParameterDeclarationProperty{
//   	Name: jsii.String("name"),
//   	ParameterValueType: jsii.String("parameterValueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &IntegerDefaultValuesProperty{
//   		DynamicValue: &DynamicDefaultValueProperty{
//   			DefaultValueColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//
//   			// the properties below are optional
//   			GroupNameColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			UserNameColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   		StaticValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	MappedDataSetParameters: []interface{}{
//   		&MappedDataSetParameterProperty{
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			DataSetParameterName: jsii.String("dataSetParameterName"),
//   		},
//   	},
//   	ValueWhenUnset: &IntegerValueWhenUnsetConfigurationProperty{
//   		CustomValue: jsii.Number(123),
//   		ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   	},
//   }
//
type CfnAnalysis_IntegerParameterDeclarationProperty struct {
	// The name of the parameter that is being declared.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value type determines whether the parameter is a single-value or multi-value parameter.
	ParameterValueType *string `field:"required" json:"parameterValueType" yaml:"parameterValueType"`
	// The default values of a parameter.
	//
	// If the parameter is a single-value parameter, a maximum of one default value can be provided.
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// `CfnAnalysis.IntegerParameterDeclarationProperty.MappedDataSetParameters`.
	MappedDataSetParameters interface{} `field:"optional" json:"mappedDataSetParameters" yaml:"mappedDataSetParameters"`
	// A parameter declaration for the `Integer` data type.
	ValueWhenUnset interface{} `field:"optional" json:"valueWhenUnset" yaml:"valueWhenUnset"`
}


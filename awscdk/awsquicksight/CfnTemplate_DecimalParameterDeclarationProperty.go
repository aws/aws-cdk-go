package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalParameterDeclarationProperty := &DecimalParameterDeclarationProperty{
//   	Name: jsii.String("name"),
//   	ParameterValueType: jsii.String("parameterValueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &DecimalDefaultValuesProperty{
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
//   	ValueWhenUnset: &DecimalValueWhenUnsetConfigurationProperty{
//   		CustomValue: jsii.Number(123),
//   		ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   	},
//   }
//
type CfnTemplate_DecimalParameterDeclarationProperty struct {
	// `CfnTemplate.DecimalParameterDeclarationProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnTemplate.DecimalParameterDeclarationProperty.ParameterValueType`.
	ParameterValueType *string `field:"required" json:"parameterValueType" yaml:"parameterValueType"`
	// `CfnTemplate.DecimalParameterDeclarationProperty.DefaultValues`.
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// `CfnTemplate.DecimalParameterDeclarationProperty.MappedDataSetParameters`.
	MappedDataSetParameters interface{} `field:"optional" json:"mappedDataSetParameters" yaml:"mappedDataSetParameters"`
	// `CfnTemplate.DecimalParameterDeclarationProperty.ValueWhenUnset`.
	ValueWhenUnset interface{} `field:"optional" json:"valueWhenUnset" yaml:"valueWhenUnset"`
}


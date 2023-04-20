package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringParameterDeclarationProperty := &StringParameterDeclarationProperty{
//   	Name: jsii.String("name"),
//   	ParameterValueType: jsii.String("parameterValueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &StringDefaultValuesProperty{
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
//   		StaticValues: []*string{
//   			jsii.String("staticValues"),
//   		},
//   	},
//   	MappedDataSetParameters: []interface{}{
//   		&MappedDataSetParameterProperty{
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			DataSetParameterName: jsii.String("dataSetParameterName"),
//   		},
//   	},
//   	ValueWhenUnset: &StringValueWhenUnsetConfigurationProperty{
//   		CustomValue: jsii.String("customValue"),
//   		ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   	},
//   }
//
type CfnTemplate_StringParameterDeclarationProperty struct {
	// `CfnTemplate.StringParameterDeclarationProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnTemplate.StringParameterDeclarationProperty.ParameterValueType`.
	ParameterValueType *string `field:"required" json:"parameterValueType" yaml:"parameterValueType"`
	// `CfnTemplate.StringParameterDeclarationProperty.DefaultValues`.
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// `CfnTemplate.StringParameterDeclarationProperty.MappedDataSetParameters`.
	MappedDataSetParameters interface{} `field:"optional" json:"mappedDataSetParameters" yaml:"mappedDataSetParameters"`
	// `CfnTemplate.StringParameterDeclarationProperty.ValueWhenUnset`.
	ValueWhenUnset interface{} `field:"optional" json:"valueWhenUnset" yaml:"valueWhenUnset"`
}


package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterDeclarationProperty := &ParameterDeclarationProperty{
//   	DateTimeParameterDeclaration: &DateTimeParameterDeclarationProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		DefaultValues: &DateTimeDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				GroupNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				UserNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   			},
//   			RollingDate: &RollingDateConfigurationProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			StaticValues: []*string{
//   				jsii.String("staticValues"),
//   			},
//   		},
//   		MappedDataSetParameters: []interface{}{
//   			&MappedDataSetParameterProperty{
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				DataSetParameterName: jsii.String("dataSetParameterName"),
//   			},
//   		},
//   		TimeGranularity: jsii.String("timeGranularity"),
//   		ValueWhenUnset: &DateTimeValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.String("customValue"),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   	DecimalParameterDeclaration: &DecimalParameterDeclarationProperty{
//   		Name: jsii.String("name"),
//   		ParameterValueType: jsii.String("parameterValueType"),
//
//   		// the properties below are optional
//   		DefaultValues: &DecimalDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				GroupNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				UserNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   			},
//   			StaticValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   		MappedDataSetParameters: []interface{}{
//   			&MappedDataSetParameterProperty{
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				DataSetParameterName: jsii.String("dataSetParameterName"),
//   			},
//   		},
//   		ValueWhenUnset: &DecimalValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.Number(123),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   	IntegerParameterDeclaration: &IntegerParameterDeclarationProperty{
//   		Name: jsii.String("name"),
//   		ParameterValueType: jsii.String("parameterValueType"),
//
//   		// the properties below are optional
//   		DefaultValues: &IntegerDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				GroupNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				UserNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   			},
//   			StaticValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   		MappedDataSetParameters: []interface{}{
//   			&MappedDataSetParameterProperty{
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				DataSetParameterName: jsii.String("dataSetParameterName"),
//   			},
//   		},
//   		ValueWhenUnset: &IntegerValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.Number(123),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   	StringParameterDeclaration: &StringParameterDeclarationProperty{
//   		Name: jsii.String("name"),
//   		ParameterValueType: jsii.String("parameterValueType"),
//
//   		// the properties below are optional
//   		DefaultValues: &StringDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				GroupNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				UserNameColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   			},
//   			StaticValues: []*string{
//   				jsii.String("staticValues"),
//   			},
//   		},
//   		MappedDataSetParameters: []interface{}{
//   			&MappedDataSetParameterProperty{
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				DataSetParameterName: jsii.String("dataSetParameterName"),
//   			},
//   		},
//   		ValueWhenUnset: &StringValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.String("customValue"),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   }
//
type CfnAnalysis_ParameterDeclarationProperty struct {
	// `CfnAnalysis.ParameterDeclarationProperty.DateTimeParameterDeclaration`.
	DateTimeParameterDeclaration interface{} `field:"optional" json:"dateTimeParameterDeclaration" yaml:"dateTimeParameterDeclaration"`
	// `CfnAnalysis.ParameterDeclarationProperty.DecimalParameterDeclaration`.
	DecimalParameterDeclaration interface{} `field:"optional" json:"decimalParameterDeclaration" yaml:"decimalParameterDeclaration"`
	// `CfnAnalysis.ParameterDeclarationProperty.IntegerParameterDeclaration`.
	IntegerParameterDeclaration interface{} `field:"optional" json:"integerParameterDeclaration" yaml:"integerParameterDeclaration"`
	// `CfnAnalysis.ParameterDeclarationProperty.StringParameterDeclaration`.
	StringParameterDeclaration interface{} `field:"optional" json:"stringParameterDeclaration" yaml:"stringParameterDeclaration"`
}


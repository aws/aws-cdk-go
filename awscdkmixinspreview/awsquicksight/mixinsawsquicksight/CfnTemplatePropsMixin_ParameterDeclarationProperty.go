package mixinsawsquicksight


// The declaration definition of a parameter.
//
// For more information, see [Parameters in Amazon Quick Sight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the *Amazon Quick Suite User Guide* .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterDeclarationProperty := &ParameterDeclarationProperty{
//   	DateTimeParameterDeclaration: &DateTimeParameterDeclarationProperty{
//   		DefaultValues: &DateTimeDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
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
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				Expression: jsii.String("expression"),
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
//   		Name: jsii.String("name"),
//   		TimeGranularity: jsii.String("timeGranularity"),
//   		ValueWhenUnset: &DateTimeValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.String("customValue"),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   	DecimalParameterDeclaration: &DecimalParameterDeclarationProperty{
//   		DefaultValues: &DecimalDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
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
//   		Name: jsii.String("name"),
//   		ParameterValueType: jsii.String("parameterValueType"),
//   		ValueWhenUnset: &DecimalValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.Number(123),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   	IntegerParameterDeclaration: &IntegerParameterDeclarationProperty{
//   		DefaultValues: &IntegerDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
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
//   		Name: jsii.String("name"),
//   		ParameterValueType: jsii.String("parameterValueType"),
//   		ValueWhenUnset: &IntegerValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.Number(123),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   	StringParameterDeclaration: &StringParameterDeclarationProperty{
//   		DefaultValues: &StringDefaultValuesProperty{
//   			DynamicValue: &DynamicDefaultValueProperty{
//   				DefaultValueColumn: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
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
//   		Name: jsii.String("name"),
//   		ParameterValueType: jsii.String("parameterValueType"),
//   		ValueWhenUnset: &StringValueWhenUnsetConfigurationProperty{
//   			CustomValue: jsii.String("customValue"),
//   			ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdeclaration.html
//
type CfnTemplatePropsMixin_ParameterDeclarationProperty struct {
	// A parameter declaration for the `DateTime` data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdeclaration.html#cfn-quicksight-template-parameterdeclaration-datetimeparameterdeclaration
	//
	DateTimeParameterDeclaration interface{} `field:"optional" json:"dateTimeParameterDeclaration" yaml:"dateTimeParameterDeclaration"`
	// A parameter declaration for the `Decimal` data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdeclaration.html#cfn-quicksight-template-parameterdeclaration-decimalparameterdeclaration
	//
	DecimalParameterDeclaration interface{} `field:"optional" json:"decimalParameterDeclaration" yaml:"decimalParameterDeclaration"`
	// A parameter declaration for the `Integer` data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdeclaration.html#cfn-quicksight-template-parameterdeclaration-integerparameterdeclaration
	//
	IntegerParameterDeclaration interface{} `field:"optional" json:"integerParameterDeclaration" yaml:"integerParameterDeclaration"`
	// A parameter declaration for the `String` data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-parameterdeclaration.html#cfn-quicksight-template-parameterdeclaration-stringparameterdeclaration
	//
	StringParameterDeclaration interface{} `field:"optional" json:"stringParameterDeclaration" yaml:"stringParameterDeclaration"`
}


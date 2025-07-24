package awsquicksight


// A parameter declaration for the `Decimal` data type.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalparameterdeclaration.html
//
type CfnTemplate_DecimalParameterDeclarationProperty struct {
	// The name of the parameter that is being declared.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalparameterdeclaration.html#cfn-quicksight-template-decimalparameterdeclaration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value type determines whether the parameter is a single-value or multi-value parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalparameterdeclaration.html#cfn-quicksight-template-decimalparameterdeclaration-parametervaluetype
	//
	ParameterValueType *string `field:"required" json:"parameterValueType" yaml:"parameterValueType"`
	// The default values of a parameter.
	//
	// If the parameter is a single-value parameter, a maximum of one default value can be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalparameterdeclaration.html#cfn-quicksight-template-decimalparameterdeclaration-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalparameterdeclaration.html#cfn-quicksight-template-decimalparameterdeclaration-mappeddatasetparameters
	//
	MappedDataSetParameters interface{} `field:"optional" json:"mappedDataSetParameters" yaml:"mappedDataSetParameters"`
	// The configuration that defines the default value of a `Decimal` parameter when a value has not been set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalparameterdeclaration.html#cfn-quicksight-template-decimalparameterdeclaration-valuewhenunset
	//
	ValueWhenUnset interface{} `field:"optional" json:"valueWhenUnset" yaml:"valueWhenUnset"`
}


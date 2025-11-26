package previewawsquicksightmixins


// A parameter declaration for the `String` data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stringParameterDeclarationProperty := &StringParameterDeclarationProperty{
//   	DefaultValues: &StringDefaultValuesProperty{
//   		DynamicValue: &DynamicDefaultValueProperty{
//   			DefaultValueColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
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
//   	Name: jsii.String("name"),
//   	ParameterValueType: jsii.String("parameterValueType"),
//   	ValueWhenUnset: &StringValueWhenUnsetConfigurationProperty{
//   		CustomValue: jsii.String("customValue"),
//   		ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameterdeclaration.html
//
type CfnDashboardPropsMixin_StringParameterDeclarationProperty struct {
	// The default values of a parameter.
	//
	// If the parameter is a single-value parameter, a maximum of one default value can be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameterdeclaration.html#cfn-quicksight-dashboard-stringparameterdeclaration-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameterdeclaration.html#cfn-quicksight-dashboard-stringparameterdeclaration-mappeddatasetparameters
	//
	MappedDataSetParameters interface{} `field:"optional" json:"mappedDataSetParameters" yaml:"mappedDataSetParameters"`
	// The name of the parameter that is being declared.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameterdeclaration.html#cfn-quicksight-dashboard-stringparameterdeclaration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value type determines whether the parameter is a single-value or multi-value parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameterdeclaration.html#cfn-quicksight-dashboard-stringparameterdeclaration-parametervaluetype
	//
	ParameterValueType *string `field:"optional" json:"parameterValueType" yaml:"parameterValueType"`
	// The configuration that defines the default value of a `String` parameter when a value has not been set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-stringparameterdeclaration.html#cfn-quicksight-dashboard-stringparameterdeclaration-valuewhenunset
	//
	ValueWhenUnset interface{} `field:"optional" json:"valueWhenUnset" yaml:"valueWhenUnset"`
}


package awsquicksight


// A parameter declaration for the `DateTime` data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeParameterDeclarationProperty := &DateTimeParameterDeclarationProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DefaultValues: &DateTimeDefaultValuesProperty{
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
//   		RollingDate: &RollingDateConfigurationProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
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
//   	TimeGranularity: jsii.String("timeGranularity"),
//   	ValueWhenUnset: &DateTimeValueWhenUnsetConfigurationProperty{
//   		CustomValue: jsii.String("customValue"),
//   		ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameterdeclaration.html
//
type CfnAnalysis_DateTimeParameterDeclarationProperty struct {
	// The name of the parameter that is being declared.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameterdeclaration.html#cfn-quicksight-analysis-datetimeparameterdeclaration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The default values of a parameter.
	//
	// If the parameter is a single-value parameter, a maximum of one default value can be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameterdeclaration.html#cfn-quicksight-analysis-datetimeparameterdeclaration-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameterdeclaration.html#cfn-quicksight-analysis-datetimeparameterdeclaration-mappeddatasetparameters
	//
	MappedDataSetParameters interface{} `field:"optional" json:"mappedDataSetParameters" yaml:"mappedDataSetParameters"`
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameterdeclaration.html#cfn-quicksight-analysis-datetimeparameterdeclaration-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
	// The configuration that defines the default value of a `DateTime` parameter when a value has not been set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datetimeparameterdeclaration.html#cfn-quicksight-analysis-datetimeparameterdeclaration-valuewhenunset
	//
	ValueWhenUnset interface{} `field:"optional" json:"valueWhenUnset" yaml:"valueWhenUnset"`
}


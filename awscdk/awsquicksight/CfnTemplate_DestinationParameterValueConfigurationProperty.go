package awsquicksight


// The configuration of destination parameter values.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationParameterValueConfigurationProperty := &DestinationParameterValueConfigurationProperty{
//   	CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   		CustomValues: &CustomParameterValuesProperty{
//   			DateTimeValues: []*string{
//   				jsii.String("dateTimeValues"),
//   			},
//   			DecimalValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			IntegerValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			StringValues: []*string{
//   				jsii.String("stringValues"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		IncludeNullValue: jsii.Boolean(false),
//   	},
//   	SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   	SourceColumn: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	SourceField: jsii.String("sourceField"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-destinationparametervalueconfiguration.html
//
type CfnTemplate_DestinationParameterValueConfigurationProperty struct {
	// The configuration of custom values for destination parameter in `DestinationParameterValueConfiguration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-destinationparametervalueconfiguration.html#cfn-quicksight-template-destinationparametervalueconfiguration-customvaluesconfiguration
	//
	CustomValuesConfiguration interface{} `field:"optional" json:"customValuesConfiguration" yaml:"customValuesConfiguration"`
	// The configuration that selects all options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-destinationparametervalueconfiguration.html#cfn-quicksight-template-destinationparametervalueconfiguration-selectallvalueoptions
	//
	SelectAllValueOptions *string `field:"optional" json:"selectAllValueOptions" yaml:"selectAllValueOptions"`
	// A column of a data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-destinationparametervalueconfiguration.html#cfn-quicksight-template-destinationparametervalueconfiguration-sourcecolumn
	//
	SourceColumn interface{} `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// The source field ID of the destination parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-destinationparametervalueconfiguration.html#cfn-quicksight-template-destinationparametervalueconfiguration-sourcefield
	//
	SourceField *string `field:"optional" json:"sourceField" yaml:"sourceField"`
	// The source parameter name of the destination parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-destinationparametervalueconfiguration.html#cfn-quicksight-template-destinationparametervalueconfiguration-sourceparametername
	//
	SourceParameterName *string `field:"optional" json:"sourceParameterName" yaml:"sourceParameterName"`
}


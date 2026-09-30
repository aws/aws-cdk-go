package awsquicksight


// The default values of the `StringParameterDeclaration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringDefaultValuesProperty := &StringDefaultValuesProperty{
//   	DynamicValue: &DynamicDefaultValueProperty{
//   		DefaultValueColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//
//   			// the properties below are optional
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			TopicIdentifier: jsii.String("topicIdentifier"),
//   		},
//
//   		// the properties below are optional
//   		GroupNameColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//
//   			// the properties below are optional
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			TopicIdentifier: jsii.String("topicIdentifier"),
//   		},
//   		UserNameColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//
//   			// the properties below are optional
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			TopicIdentifier: jsii.String("topicIdentifier"),
//   		},
//   	},
//   	StaticValues: []*string{
//   		jsii.String("staticValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringdefaultvalues.html
//
type CfnAnalysis_StringDefaultValuesProperty struct {
	// The dynamic value of the `StringDefaultValues` .
	//
	// Different defaults displayed according to users, groups, and values mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringdefaultvalues.html#cfn-quicksight-analysis-stringdefaultvalues-dynamicvalue
	//
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// The static values of the `DecimalDefaultValues` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringdefaultvalues.html#cfn-quicksight-analysis-stringdefaultvalues-staticvalues
	//
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


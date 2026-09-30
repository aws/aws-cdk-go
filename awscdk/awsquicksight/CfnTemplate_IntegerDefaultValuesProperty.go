package awsquicksight


// The default values of the `IntegerParameterDeclaration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerDefaultValuesProperty := &IntegerDefaultValuesProperty{
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
//   	StaticValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-integerdefaultvalues.html
//
type CfnTemplate_IntegerDefaultValuesProperty struct {
	// The dynamic value of the `IntegerDefaultValues` .
	//
	// Different defaults are displayed according to users, groups, and values mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-integerdefaultvalues.html#cfn-quicksight-template-integerdefaultvalues-dynamicvalue
	//
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// The static values of the `IntegerDefaultValues` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-integerdefaultvalues.html#cfn-quicksight-template-integerdefaultvalues-staticvalues
	//
	StaticValues interface{} `field:"optional" json:"staticValues" yaml:"staticValues"`
}


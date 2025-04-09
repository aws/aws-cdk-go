package awsquicksight


// The default values of the `DateTimeParameterDeclaration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeDefaultValuesProperty := &DateTimeDefaultValuesProperty{
//   	DynamicValue: &DynamicDefaultValueProperty{
//   		DefaultValueColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//
//   		// the properties below are optional
//   		GroupNameColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		UserNameColumn: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
//   	RollingDate: &RollingDateConfigurationProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	StaticValues: []*string{
//   		jsii.String("staticValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimedefaultvalues.html
//
type CfnTemplate_DateTimeDefaultValuesProperty struct {
	// The dynamic value of the `DataTimeDefaultValues` .
	//
	// Different defaults are displayed according to users, groups, and values mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimedefaultvalues.html#cfn-quicksight-template-datetimedefaultvalues-dynamicvalue
	//
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// The rolling date of the `DataTimeDefaultValues` .
	//
	// The date is determined from the dataset based on input expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimedefaultvalues.html#cfn-quicksight-template-datetimedefaultvalues-rollingdate
	//
	RollingDate interface{} `field:"optional" json:"rollingDate" yaml:"rollingDate"`
	// The static values of the `DataTimeDefaultValues` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimedefaultvalues.html#cfn-quicksight-template-datetimedefaultvalues-staticvalues
	//
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


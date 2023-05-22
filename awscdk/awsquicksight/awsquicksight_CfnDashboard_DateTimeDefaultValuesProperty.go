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
type CfnDashboard_DateTimeDefaultValuesProperty struct {
	// The dynamic value of the `DataTimeDefaultValues` .
	//
	// Different defaults are displayed according to users, groups, and values mapping.
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// The rolling date of the `DataTimeDefaultValues` .
	//
	// The date is determined from the dataset based on input expression.
	RollingDate interface{} `field:"optional" json:"rollingDate" yaml:"rollingDate"`
	// The static values of the `DataTimeDefaultValues` .
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


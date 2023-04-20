package awsquicksight


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
type CfnTemplate_DateTimeDefaultValuesProperty struct {
	// `CfnTemplate.DateTimeDefaultValuesProperty.DynamicValue`.
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// `CfnTemplate.DateTimeDefaultValuesProperty.RollingDate`.
	RollingDate interface{} `field:"optional" json:"rollingDate" yaml:"rollingDate"`
	// `CfnTemplate.DateTimeDefaultValuesProperty.StaticValues`.
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


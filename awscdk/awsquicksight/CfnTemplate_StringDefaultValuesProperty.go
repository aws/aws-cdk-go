package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringDefaultValuesProperty := &StringDefaultValuesProperty{
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
//   	StaticValues: []*string{
//   		jsii.String("staticValues"),
//   	},
//   }
//
type CfnTemplate_StringDefaultValuesProperty struct {
	// `CfnTemplate.StringDefaultValuesProperty.DynamicValue`.
	DynamicValue interface{} `field:"optional" json:"dynamicValue" yaml:"dynamicValue"`
	// `CfnTemplate.StringDefaultValuesProperty.StaticValues`.
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


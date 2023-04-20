package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicDefaultValueProperty := &DynamicDefaultValueProperty{
//   	DefaultValueColumn: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	GroupNameColumn: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	UserNameColumn: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   }
//
type CfnDashboard_DynamicDefaultValueProperty struct {
	// `CfnDashboard.DynamicDefaultValueProperty.DefaultValueColumn`.
	DefaultValueColumn interface{} `field:"required" json:"defaultValueColumn" yaml:"defaultValueColumn"`
	// `CfnDashboard.DynamicDefaultValueProperty.GroupNameColumn`.
	GroupNameColumn interface{} `field:"optional" json:"groupNameColumn" yaml:"groupNameColumn"`
	// `CfnDashboard.DynamicDefaultValueProperty.UserNameColumn`.
	UserNameColumn interface{} `field:"optional" json:"userNameColumn" yaml:"userNameColumn"`
}


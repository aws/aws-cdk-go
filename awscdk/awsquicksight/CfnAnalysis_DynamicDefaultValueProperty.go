package awsquicksight


// Defines different defaults to the users or groups based on mapping.
//
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
type CfnAnalysis_DynamicDefaultValueProperty struct {
	// The column that contains the default value of each user or group.
	DefaultValueColumn interface{} `field:"required" json:"defaultValueColumn" yaml:"defaultValueColumn"`
	// The column that contains the group name.
	GroupNameColumn interface{} `field:"optional" json:"groupNameColumn" yaml:"groupNameColumn"`
	// The column that contains the username.
	UserNameColumn interface{} `field:"optional" json:"userNameColumn" yaml:"userNameColumn"`
}


package mixinsawsquicksight


// Defines different defaults to the users or groups based on mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynamicDefaultValueProperty := &DynamicDefaultValueProperty{
//   	DefaultValueColumn: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-dynamicdefaultvalue.html
//
type CfnAnalysisPropsMixin_DynamicDefaultValueProperty struct {
	// The column that contains the default value of each user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-dynamicdefaultvalue.html#cfn-quicksight-analysis-dynamicdefaultvalue-defaultvaluecolumn
	//
	DefaultValueColumn interface{} `field:"optional" json:"defaultValueColumn" yaml:"defaultValueColumn"`
	// The column that contains the group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-dynamicdefaultvalue.html#cfn-quicksight-analysis-dynamicdefaultvalue-groupnamecolumn
	//
	GroupNameColumn interface{} `field:"optional" json:"groupNameColumn" yaml:"groupNameColumn"`
	// The column that contains the username.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-dynamicdefaultvalue.html#cfn-quicksight-analysis-dynamicdefaultvalue-usernamecolumn
	//
	UserNameColumn interface{} `field:"optional" json:"userNameColumn" yaml:"userNameColumn"`
}


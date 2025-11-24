package mixinsawsquicksight


// Configuration for row level security.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tagRuleConfigurations interface{}
//
//   rowLevelPermissionConfigurationProperty := &RowLevelPermissionConfigurationProperty{
//   	RowLevelPermissionDataSet: &RowLevelPermissionDataSetProperty{
//   		Arn: jsii.String("arn"),
//   		FormatVersion: jsii.String("formatVersion"),
//   		Namespace: jsii.String("namespace"),
//   		PermissionPolicy: jsii.String("permissionPolicy"),
//   		Status: jsii.String("status"),
//   	},
//   	TagConfiguration: &RowLevelPermissionTagConfigurationProperty{
//   		Status: jsii.String("status"),
//   		TagRuleConfigurations: tagRuleConfigurations,
//   		TagRules: []interface{}{
//   			&RowLevelPermissionTagRuleProperty{
//   				ColumnName: jsii.String("columnName"),
//   				MatchAllValue: jsii.String("matchAllValue"),
//   				TagKey: jsii.String("tagKey"),
//   				TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissionconfiguration.html
//
type CfnDataSetPropsMixin_RowLevelPermissionConfigurationProperty struct {
	// <p>Information about a dataset that contains permissions for row-level security (RLS).
	//
	// The permissions dataset maps fields to users or groups. For more information, see
	//             <a href="https://docs.aws.amazon.com/quicksight/latest/user/restrict-access-to-a-data-set-using-row-level-security.html">Using Row-Level Security (RLS) to Restrict Access to a Dataset</a> in the <i>Amazon QuickSight User
	//                 Guide</i>.</p>
	//          <p>The option to deny permissions by setting <code>PermissionPolicy</code> to <code>DENY_ACCESS</code> is
	//             not supported for new RLS datasets.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissionconfiguration.html#cfn-quicksight-dataset-rowlevelpermissionconfiguration-rowlevelpermissiondataset
	//
	RowLevelPermissionDataSet interface{} `field:"optional" json:"rowLevelPermissionDataSet" yaml:"rowLevelPermissionDataSet"`
	// <p>The configuration of tags on a dataset to set row-level security.
	//
	// </p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissionconfiguration.html#cfn-quicksight-dataset-rowlevelpermissionconfiguration-tagconfiguration
	//
	TagConfiguration interface{} `field:"optional" json:"tagConfiguration" yaml:"tagConfiguration"`
}


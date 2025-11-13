package awsquicksight


// The element you can use to define tags for row-level security.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tagRuleConfigurations interface{}
//
//   rowLevelPermissionTagConfigurationProperty := &RowLevelPermissionTagConfigurationProperty{
//   	TagRules: []interface{}{
//   		&RowLevelPermissionTagRuleProperty{
//   			ColumnName: jsii.String("columnName"),
//   			TagKey: jsii.String("tagKey"),
//
//   			// the properties below are optional
//   			MatchAllValue: jsii.String("matchAllValue"),
//   			TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   	TagRuleConfigurations: tagRuleConfigurations,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagconfiguration.html
//
type CfnDataSet_RowLevelPermissionTagConfigurationProperty struct {
	// A set of rules associated with row-level security, such as the tag names and columns that they are assigned to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagconfiguration.html#cfn-quicksight-dataset-rowlevelpermissiontagconfiguration-tagrules
	//
	TagRules interface{} `field:"required" json:"tagRules" yaml:"tagRules"`
	// The status of row-level security tags.
	//
	// If enabled, the status is `ENABLED` . If disabled, the status is `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagconfiguration.html#cfn-quicksight-dataset-rowlevelpermissiontagconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The configuration of tags on a dataset to set row-level security.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagconfiguration.html#cfn-quicksight-dataset-rowlevelpermissiontagconfiguration-tagruleconfigurations
	//
	TagRuleConfigurations interface{} `field:"optional" json:"tagRuleConfigurations" yaml:"tagRuleConfigurations"`
}


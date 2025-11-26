package previewawsquicksightmixins


// The element you can use to define tags for row-level security.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tagRuleConfigurations interface{}
//
//   rowLevelPermissionTagConfigurationProperty := &RowLevelPermissionTagConfigurationProperty{
//   	Status: jsii.String("status"),
//   	TagRuleConfigurations: tagRuleConfigurations,
//   	TagRules: []interface{}{
//   		&RowLevelPermissionTagRuleProperty{
//   			ColumnName: jsii.String("columnName"),
//   			MatchAllValue: jsii.String("matchAllValue"),
//   			TagKey: jsii.String("tagKey"),
//   			TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagconfiguration.html
//
type CfnDataSetPropsMixin_RowLevelPermissionTagConfigurationProperty struct {
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
	// A set of rules associated with row-level security, such as the tag names and columns that they are assigned to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagconfiguration.html#cfn-quicksight-dataset-rowlevelpermissiontagconfiguration-tagrules
	//
	TagRules interface{} `field:"optional" json:"tagRules" yaml:"tagRules"`
}


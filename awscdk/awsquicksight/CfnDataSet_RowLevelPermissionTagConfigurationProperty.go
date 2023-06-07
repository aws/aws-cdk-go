package awsquicksight


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
type CfnDataSet_RowLevelPermissionTagConfigurationProperty struct {
	// `CfnDataSet.RowLevelPermissionTagConfigurationProperty.TagRules`.
	TagRules interface{} `field:"required" json:"tagRules" yaml:"tagRules"`
	// `CfnDataSet.RowLevelPermissionTagConfigurationProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `CfnDataSet.RowLevelPermissionTagConfigurationProperty.TagRuleConfigurations`.
	TagRuleConfigurations interface{} `field:"optional" json:"tagRuleConfigurations" yaml:"tagRuleConfigurations"`
}


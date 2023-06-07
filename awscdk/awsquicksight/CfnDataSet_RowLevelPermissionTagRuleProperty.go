package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rowLevelPermissionTagRuleProperty := &RowLevelPermissionTagRuleProperty{
//   	ColumnName: jsii.String("columnName"),
//   	TagKey: jsii.String("tagKey"),
//
//   	// the properties below are optional
//   	MatchAllValue: jsii.String("matchAllValue"),
//   	TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   }
//
type CfnDataSet_RowLevelPermissionTagRuleProperty struct {
	// `CfnDataSet.RowLevelPermissionTagRuleProperty.ColumnName`.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// `CfnDataSet.RowLevelPermissionTagRuleProperty.TagKey`.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// `CfnDataSet.RowLevelPermissionTagRuleProperty.MatchAllValue`.
	MatchAllValue *string `field:"optional" json:"matchAllValue" yaml:"matchAllValue"`
	// `CfnDataSet.RowLevelPermissionTagRuleProperty.TagMultiValueDelimiter`.
	TagMultiValueDelimiter *string `field:"optional" json:"tagMultiValueDelimiter" yaml:"tagMultiValueDelimiter"`
}


package previewawsquicksightmixins


// A set of rules associated with a tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rowLevelPermissionTagRuleProperty := &RowLevelPermissionTagRuleProperty{
//   	ColumnName: jsii.String("columnName"),
//   	MatchAllValue: jsii.String("matchAllValue"),
//   	TagKey: jsii.String("tagKey"),
//   	TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagrule.html
//
type CfnDataSetPropsMixin_RowLevelPermissionTagRuleProperty struct {
	// The column name that a tag key is assigned to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagrule.html#cfn-quicksight-dataset-rowlevelpermissiontagrule-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// A string that you want to use to filter by all the values in a column in the dataset and don’t want to list the values one by one.
	//
	// For example, you can use an asterisk as your match all value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagrule.html#cfn-quicksight-dataset-rowlevelpermissiontagrule-matchallvalue
	//
	MatchAllValue *string `field:"optional" json:"matchAllValue" yaml:"matchAllValue"`
	// The unique key for a tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagrule.html#cfn-quicksight-dataset-rowlevelpermissiontagrule-tagkey
	//
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// A string that you want to use to delimit the values when you pass the values at run time.
	//
	// For example, you can delimit the values with a comma.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiontagrule.html#cfn-quicksight-dataset-rowlevelpermissiontagrule-tagmultivaluedelimiter
	//
	TagMultiValueDelimiter *string `field:"optional" json:"tagMultiValueDelimiter" yaml:"tagMultiValueDelimiter"`
}


package previewawsquicksightmixins


// A rule defined to grant access on one or more restricted columns.
//
// Each dataset can have multiple rules. To create a restricted column, you add it to one or more rules. Each rule must contain at least one column and at least one user or group. To be able to see a restricted column, a user or group needs to be added to a rule for that column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnLevelPermissionRuleProperty := &ColumnLevelPermissionRuleProperty{
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	Principals: []*string{
//   		jsii.String("principals"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnlevelpermissionrule.html
//
type CfnDataSetPropsMixin_ColumnLevelPermissionRuleProperty struct {
	// An array of column names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnlevelpermissionrule.html#cfn-quicksight-dataset-columnlevelpermissionrule-columnnames
	//
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// An array of Amazon Resource Names (ARNs) for Quick Suite users or groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnlevelpermissionrule.html#cfn-quicksight-dataset-columnlevelpermissionrule-principals
	//
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}


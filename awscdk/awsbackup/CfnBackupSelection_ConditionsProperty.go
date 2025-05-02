package awsbackup


// Contains information about which resources to include or exclude from a backup plan using their tags.
//
// Conditions are case sensitive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionsProperty := &ConditionsProperty{
//   	StringEquals: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	StringLike: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	StringNotEquals: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   	StringNotLike: []interface{}{
//   		&ConditionParameterProperty{
//   			ConditionKey: jsii.String("conditionKey"),
//   			ConditionValue: jsii.String("conditionValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html
//
type CfnBackupSelection_ConditionsProperty struct {
	// Filters the values of your tagged resources for only those resources that you tagged with the same value.
	//
	// Also called "exact matching."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringequals
	//
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// Filters the values of your tagged resources for matching tag values with the use of a wildcard character (*) anywhere in the string.
	//
	// For example, "prod*" or "*rod*" matches the tag value "production".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringlike
	//
	StringLike interface{} `field:"optional" json:"stringLike" yaml:"stringLike"`
	// Filters the values of your tagged resources for only those resources that you tagged that do not have the same value.
	//
	// Also called "negated matching."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringnotequals
	//
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
	// Filters the values of your tagged resources for non-matching tag values with the use of a wildcard character (*) anywhere in the string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditions.html#cfn-backup-backupselection-conditions-stringnotlike
	//
	StringNotLike interface{} `field:"optional" json:"stringNotLike" yaml:"stringNotLike"`
}


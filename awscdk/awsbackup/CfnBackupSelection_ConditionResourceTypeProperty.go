package awsbackup


// Specifies an object that contains an array of triplets made up of a condition type (such as `STRINGEQUALS` ), a key, and a value.
//
// Conditions are used to filter resources in a selection that is assigned to a backup plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionResourceTypeProperty := &ConditionResourceTypeProperty{
//   	ConditionKey: jsii.String("conditionKey"),
//   	ConditionType: jsii.String("conditionType"),
//   	ConditionValue: jsii.String("conditionValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionresourcetype.html
//
type CfnBackupSelection_ConditionResourceTypeProperty struct {
	// The key in a key-value pair.
	//
	// For example, in `"Department": "accounting"` , `"Department"` is the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionresourcetype.html#cfn-backup-backupselection-conditionresourcetype-conditionkey
	//
	ConditionKey *string `field:"required" json:"conditionKey" yaml:"conditionKey"`
	// An operation, such as `STRINGEQUALS` , that is applied to a key-value pair used to filter resources in a selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionresourcetype.html#cfn-backup-backupselection-conditionresourcetype-conditiontype
	//
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
	// The value in a key-value pair.
	//
	// For example, in `"Department": "accounting"` , `"accounting"` is the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionresourcetype.html#cfn-backup-backupselection-conditionresourcetype-conditionvalue
	//
	ConditionValue *string `field:"required" json:"conditionValue" yaml:"conditionValue"`
}


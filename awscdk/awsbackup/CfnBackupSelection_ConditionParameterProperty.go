package awsbackup


// Includes information about tags you define to assign tagged resources to a backup plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionParameterProperty := &ConditionParameterProperty{
//   	ConditionKey: jsii.String("conditionKey"),
//   	ConditionValue: jsii.String("conditionValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionparameter.html
//
type CfnBackupSelection_ConditionParameterProperty struct {
	// The key in a key-value pair.
	//
	// For example, in the tag `Department: Accounting` , `Department` is the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionparameter.html#cfn-backup-backupselection-conditionparameter-conditionkey
	//
	ConditionKey *string `field:"optional" json:"conditionKey" yaml:"conditionKey"`
	// The value in a key-value pair.
	//
	// For example, in the tag `Department: Accounting` , `Accounting` is the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionparameter.html#cfn-backup-backupselection-conditionparameter-conditionvalue
	//
	ConditionValue *string `field:"optional" json:"conditionValue" yaml:"conditionValue"`
}


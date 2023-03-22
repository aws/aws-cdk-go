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
//   conditionResourceTypeProperty := &conditionResourceTypeProperty{
//   	conditionKey: jsii.String("conditionKey"),
//   	conditionType: jsii.String("conditionType"),
//   	conditionValue: jsii.String("conditionValue"),
//   }
//
type CfnBackupSelection_ConditionResourceTypeProperty struct {
	// The key in a key-value pair.
	//
	// For example, in `"Department": "accounting"` , `"Department"` is the key.
	ConditionKey *string `field:"required" json:"conditionKey" yaml:"conditionKey"`
	// An operation, such as `STRINGEQUALS` , that is applied to a key-value pair used to filter resources in a selection.
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
	// The value in a key-value pair.
	//
	// For example, in `"Department": "accounting"` , `"accounting"` is the value.
	ConditionValue *string `field:"required" json:"conditionValue" yaml:"conditionValue"`
}


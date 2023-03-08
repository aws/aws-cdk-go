package awsbackup


// Properties for defining a `CfnBackupSelection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//
//   cfnBackupSelectionProps := &CfnBackupSelectionProps{
//   	BackupPlanId: jsii.String("backupPlanId"),
//   	BackupSelection: &BackupSelectionResourceTypeProperty{
//   		IamRoleArn: jsii.String("iamRoleArn"),
//   		SelectionName: jsii.String("selectionName"),
//
//   		// the properties below are optional
//   		Conditions: conditions,
//   		ListOfTags: []interface{}{
//   			&ConditionResourceTypeProperty{
//   				ConditionKey: jsii.String("conditionKey"),
//   				ConditionType: jsii.String("conditionType"),
//   				ConditionValue: jsii.String("conditionValue"),
//   			},
//   		},
//   		NotResources: []*string{
//   			jsii.String("notResources"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   	},
//   }
//
type CfnBackupSelectionProps struct {
	// Uniquely identifies a backup plan.
	BackupPlanId *string `field:"required" json:"backupPlanId" yaml:"backupPlanId"`
	// Specifies the body of a request to assign a set of resources to a backup plan.
	//
	// It includes an array of resources, an optional array of patterns to exclude resources, an optional role to provide access to the AWS service the resource belongs to, and an optional array of tags used to identify a set of resources.
	BackupSelection interface{} `field:"required" json:"backupSelection" yaml:"backupSelection"`
}


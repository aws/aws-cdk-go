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
//   cfnBackupSelectionProps := &cfnBackupSelectionProps{
//   	backupPlanId: jsii.String("backupPlanId"),
//   	backupSelection: &backupSelectionResourceTypeProperty{
//   		iamRoleArn: jsii.String("iamRoleArn"),
//   		selectionName: jsii.String("selectionName"),
//
//   		// the properties below are optional
//   		conditions: conditions,
//   		listOfTags: []interface{}{
//   			&conditionResourceTypeProperty{
//   				conditionKey: jsii.String("conditionKey"),
//   				conditionType: jsii.String("conditionType"),
//   				conditionValue: jsii.String("conditionValue"),
//   			},
//   		},
//   		notResources: []*string{
//   			jsii.String("notResources"),
//   		},
//   		resources: []*string{
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


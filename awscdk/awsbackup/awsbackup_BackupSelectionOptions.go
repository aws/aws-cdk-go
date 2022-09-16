package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Options for a BackupSelection.
//
// Example:
//   var plan backupPlan
//
//   myTable := dynamodb.table.fromTableName(this, jsii.String("Table"), jsii.String("myTableName"))
//   myCoolConstruct := constructs.NewConstruct(this, jsii.String("MyCoolConstruct"))
//
//   plan.addSelection(jsii.String("Selection"), &backupSelectionOptions{
//   	resources: []backupResource{
//   		backup.*backupResource.fromDynamoDbTable(myTable),
//   		backup.*backupResource.fromTag(jsii.String("stage"), jsii.String("prod")),
//   		backup.*backupResource.fromConstruct(myCoolConstruct),
//   	},
//   })
//
// Experimental.
type BackupSelectionOptions struct {
	// The resources to backup.
	//
	// Use the helper static methods defined on `BackupResource`.
	// Experimental.
	Resources *[]BackupResource `field:"required" json:"resources" yaml:"resources"`
	// Whether to automatically give restores permissions to the role that AWS Backup uses.
	//
	// If `true`, the `AWSBackupServiceRolePolicyForRestores` managed
	// policy will be attached to the role.
	// Experimental.
	AllowRestores *bool `field:"optional" json:"allowRestores" yaml:"allowRestores"`
	// The name for this selection.
	// Experimental.
	BackupSelectionName *string `field:"optional" json:"backupSelectionName" yaml:"backupSelectionName"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for a BackupSelection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupPlan backupPlan
//   var backupResource backupResource
//   var role role
//
//   backupSelectionProps := &backupSelectionProps{
//   	backupPlan: backupPlan,
//   	resources: []*backupResource{
//   		backupResource,
//   	},
//
//   	// the properties below are optional
//   	allowRestores: jsii.Boolean(false),
//   	backupSelectionName: jsii.String("backupSelectionName"),
//   	role: role,
//   }
//
// Experimental.
type BackupSelectionProps struct {
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
	// The backup plan for this selection.
	// Experimental.
	BackupPlan IBackupPlan `field:"required" json:"backupPlan" yaml:"backupPlan"`
}


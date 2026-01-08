package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup"
)

// Properties for a BackupSelection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupPlanRef IBackupPlanRef
//   var backupResource BackupResource
//   var role Role
//
//   backupSelectionProps := &BackupSelectionProps{
//   	BackupPlan: backupPlanRef,
//   	Resources: []BackupResource{
//   		backupResource,
//   	},
//
//   	// the properties below are optional
//   	AllowRestores: jsii.Boolean(false),
//   	BackupSelectionName: jsii.String("backupSelectionName"),
//   	DisableDefaultBackupPolicy: jsii.Boolean(false),
//   	Role: role,
//   }
//
type BackupSelectionProps struct {
	// The resources to backup.
	//
	// Use the helper static methods defined on `BackupResource`.
	Resources *[]BackupResource `field:"required" json:"resources" yaml:"resources"`
	// Whether to automatically give restores permissions to the role that AWS Backup uses.
	//
	// If `true`, the `AWSBackupServiceRolePolicyForRestores` managed
	// policy will be attached to the role.
	// Default: false.
	//
	AllowRestores *bool `field:"optional" json:"allowRestores" yaml:"allowRestores"`
	// The name for this selection.
	// Default: - a CDK generated name.
	//
	BackupSelectionName *string `field:"optional" json:"backupSelectionName" yaml:"backupSelectionName"`
	// Whether to disable automatically assigning default backup permissions to the role that AWS Backup uses.
	//
	// If `false`, the `AWSBackupServiceRolePolicyForBackup` managed policy will be
	// attached to the role.
	// Default: false.
	//
	DisableDefaultBackupPolicy *bool `field:"optional" json:"disableDefaultBackupPolicy" yaml:"disableDefaultBackupPolicy"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role unless `disableDefaultBackupPolicy`
	// is set to `true`.
	// Default: - a new role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The backup plan for this selection.
	BackupPlan interfacesawsbackup.IBackupPlanRef `field:"required" json:"backupPlan" yaml:"backupPlan"`
}


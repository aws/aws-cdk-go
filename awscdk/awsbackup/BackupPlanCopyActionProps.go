package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a BackupPlanCopyAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupVault BackupVault
//
//   backupPlanCopyActionProps := &BackupPlanCopyActionProps{
//   	DestinationBackupVault: backupVault,
//
//   	// the properties below are optional
//   	DeleteAfter: cdk.Duration_Minutes(jsii.Number(30)),
//   	MoveToColdStorageAfter: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type BackupPlanCopyActionProps struct {
	// Destination Vault for recovery points to be copied into.
	DestinationBackupVault IBackupVault `field:"required" json:"destinationBackupVault" yaml:"destinationBackupVault"`
	// Specifies the duration after creation that a copied recovery point is deleted from the destination vault.
	//
	// Must be at least 90 days greater than `moveToColdStorageAfter`, if specified.
	// Default: - recovery point is never deleted.
	//
	DeleteAfter awscdk.Duration `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
	// Specifies the duration after creation that a copied recovery point is moved to cold storage.
	// Default: - recovery point is never moved to cold storage.
	//
	MoveToColdStorageAfter awscdk.Duration `field:"optional" json:"moveToColdStorageAfter" yaml:"moveToColdStorageAfter"`
}


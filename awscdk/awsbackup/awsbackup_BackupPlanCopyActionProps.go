package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a BackupPlanCopyAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupVault backupVault
//
//   backupPlanCopyActionProps := &backupPlanCopyActionProps{
//   	destinationBackupVault: backupVault,
//
//   	// the properties below are optional
//   	deleteAfter: cdk.duration.minutes(jsii.Number(30)),
//   	moveToColdStorageAfter: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type BackupPlanCopyActionProps struct {
	// Destination Vault for recovery points to be copied into.
	DestinationBackupVault IBackupVault `field:"required" json:"destinationBackupVault" yaml:"destinationBackupVault"`
	// Specifies the duration after creation that a copied recovery point is deleted from the destination vault.
	//
	// Must be at least 90 days greater than `moveToColdStorageAfter`, if specified.
	DeleteAfter awscdk.Duration `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
	// Specifies the duration after creation that a copied recovery point is moved to cold storage.
	MoveToColdStorageAfter awscdk.Duration `field:"optional" json:"moveToColdStorageAfter" yaml:"moveToColdStorageAfter"`
}


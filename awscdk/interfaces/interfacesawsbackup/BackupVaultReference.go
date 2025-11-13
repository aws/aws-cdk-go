package interfacesawsbackup


// A reference to a BackupVault resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupVaultReference := &BackupVaultReference{
//   	BackupVaultArn: jsii.String("backupVaultArn"),
//   	BackupVaultName: jsii.String("backupVaultName"),
//   }
//
type BackupVaultReference struct {
	// The ARN of the BackupVault resource.
	BackupVaultArn *string `field:"required" json:"backupVaultArn" yaml:"backupVaultArn"`
	// The BackupVaultName of the BackupVault resource.
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
}


package awscdkelasticachealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Backup configuration for ServerlessCache.
//
// Example:
//   var vpc vpc
//
//
//   serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
//   	Engine: elasticache.CacheEngine_VALKEY_LATEST,
//   	Backup: &BackupSettings{
//   		// set a backup name before deleting a cache
//   		BackupNameBeforeDeletion: jsii.String("my-final-backup-name"),
//   	},
//   	Vpc: Vpc,
//   })
//
// Experimental.
type BackupSettings struct {
	// ARNs of backups from which to restore data into the new cache.
	// Default: - Create a new cache with no existing data.
	//
	// Experimental.
	BackupArnsToRestore *[]*string `field:"optional" json:"backupArnsToRestore" yaml:"backupArnsToRestore"`
	// Name for the final backup taken before deletion.
	// Default: - No final backup.
	//
	// Experimental.
	BackupNameBeforeDeletion *string `field:"optional" json:"backupNameBeforeDeletion" yaml:"backupNameBeforeDeletion"`
	// Number of days to retain backups (1-35).
	// Default: - Backups are not retained.
	//
	// Experimental.
	BackupRetentionLimit *float64 `field:"optional" json:"backupRetentionLimit" yaml:"backupRetentionLimit"`
	// Automated daily backup UTC time.
	// Default: - No automated backups.
	//
	// Experimental.
	BackupTime awsevents.Schedule `field:"optional" json:"backupTime" yaml:"backupTime"`
}


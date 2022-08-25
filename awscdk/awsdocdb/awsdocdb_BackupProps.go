package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Backup configuration for DocumentDB databases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupProps := &backupProps{
//   	retention: cdk.duration.minutes(jsii.Number(30)),
//
//   	// the properties below are optional
//   	preferredWindow: jsii.String("preferredWindow"),
//   }
//
// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/backup-restore.db-cluster-snapshots.html#backup-restore.backup-window
//
type BackupProps struct {
	// How many days to retain the backup.
	Retention awscdk.Duration `field:"required" json:"retention" yaml:"retention"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'.
	PreferredWindow *string `field:"optional" json:"preferredWindow" yaml:"preferredWindow"`
}


package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Backup configuration for DocumentDB databases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   backupProps := &backupProps{
//   	retention: duration,
//
//   	// the properties below are optional
//   	preferredWindow: jsii.String("preferredWindow"),
//   }
//
// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/backup-restore.db-cluster-snapshots.html#backup-restore.backup-window
//
// Experimental.
type BackupProps struct {
	// How many days to retain the backup.
	// Experimental.
	Retention awscdk.Duration `field:"required" json:"retention" yaml:"retention"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'.
	// Experimental.
	PreferredWindow *string `field:"optional" json:"preferredWindow" yaml:"preferredWindow"`
}


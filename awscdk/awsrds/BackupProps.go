package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Backup configuration for RDS databases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupProps := &BackupProps{
//   	Retention: cdk.Duration_Minutes(jsii.Number(30)),
//
//   	// the properties below are optional
//   	PreferredWindow: jsii.String("preferredWindow"),
//   }
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
//
// Default: - The retention period for automated backups is 1 day.
// The preferred backup window will be a 30-minute window selected at random
// from an 8-hour block of time for each AWS Region.
//
type BackupProps struct {
	// How many days to retain the backup.
	Retention awscdk.Duration `field:"required" json:"retention" yaml:"retention"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'.
	// Default: - a 30-minute window selected at random from an 8-hour block of
	// time for each AWS Region. To see the time blocks available, see
	// https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
	//
	PreferredWindow *string `field:"optional" json:"preferredWindow" yaml:"preferredWindow"`
}


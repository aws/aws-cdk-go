package awsfsx


// Properties required for setting up a daily automatic backup time.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   lustreConfiguration := map[string]interface{}{
//   	// ...
//   	"automaticBackupRetention": cdk.Duration_days(jsii.Number(3)),
//   	 // backup retention
//   	"copyTagsToBackups": jsii.Boolean(true),
//   	 // if true, tags are copied to backups
//   	"dailyAutomaticBackupStartTime": fsx.NewDailyAutomaticBackupStartTime(&DailyAutomaticBackupStartTimeProps{
//   		"hour": jsii.Number(11),
//   		"minute": jsii.Number(30),
//   	}),
//   }
//
type DailyAutomaticBackupStartTimeProps struct {
	// The hour of the day (from 0-23) for automatic backup starts.
	Hour *float64 `field:"required" json:"hour" yaml:"hour"`
	// The minute of the hour (from 0-59) for automatic backup starts.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
}


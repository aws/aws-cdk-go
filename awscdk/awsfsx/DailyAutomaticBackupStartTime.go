package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Class for scheduling a daily automatic backup time.
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
type DailyAutomaticBackupStartTime interface {
	// Converts an hour, and minute into HH:MM string.
	ToTimestamp() *string
}

// The jsii proxy struct for DailyAutomaticBackupStartTime
type jsiiProxy_DailyAutomaticBackupStartTime struct {
	_ byte // padding
}

func NewDailyAutomaticBackupStartTime(props *DailyAutomaticBackupStartTimeProps) DailyAutomaticBackupStartTime {
	_init_.Initialize()

	if err := validateNewDailyAutomaticBackupStartTimeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DailyAutomaticBackupStartTime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.DailyAutomaticBackupStartTime",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewDailyAutomaticBackupStartTime_Override(d DailyAutomaticBackupStartTime, props *DailyAutomaticBackupStartTimeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.DailyAutomaticBackupStartTime",
		[]interface{}{props},
		d,
	)
}

func (d *jsiiProxy_DailyAutomaticBackupStartTime) ToTimestamp() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toTimestamp",
		nil, // no parameters
		&returns,
	)

	return returns
}


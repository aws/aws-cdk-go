package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Class for scheduling a weekly manitenance time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lustreMaintenanceTime := awscdk.Aws_fsx.NewLustreMaintenanceTime(&lustreMaintenanceTimeProps{
//   	day: awscdk.*Aws_fsx.weekday_MONDAY,
//   	hour: jsii.Number(123),
//   	minute: jsii.Number(123),
//   })
//
type LustreMaintenanceTime interface {
	// Converts a day, hour, and minute into a timestamp as used by FSx for Lustre's weeklyMaintenanceStartTime field.
	ToTimestamp() *string
}

// The jsii proxy struct for LustreMaintenanceTime
type jsiiProxy_LustreMaintenanceTime struct {
	_ byte // padding
}

func NewLustreMaintenanceTime(props *LustreMaintenanceTimeProps) LustreMaintenanceTime {
	_init_.Initialize()

	if err := validateNewLustreMaintenanceTimeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LustreMaintenanceTime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreMaintenanceTime",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLustreMaintenanceTime_Override(l LustreMaintenanceTime, props *LustreMaintenanceTimeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreMaintenanceTime",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_LustreMaintenanceTime) ToTimestamp() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toTimestamp",
		nil, // no parameters
		&returns,
	)

	return returns
}


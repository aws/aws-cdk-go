package awscdkiotalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The day of the month on which the scheduled audit takes place.
//
// Example:
//   var config AccountAuditConfiguration
//
//
//   // Daily audit
//   dailyAudit := iot.NewScheduledAudit(this, jsii.String("DailyAudit"), &ScheduledAuditProps{
//   	AccountAuditConfiguration: config,
//   	Frequency: iot.Frequency_DAILY,
//   	AuditChecks: []AuditCheck{
//   		iot.AuditCheck_AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK,
//   	},
//   })
//
//   // Weekly audit
//   weeklyAudit := iot.NewScheduledAudit(this, jsii.String("WeeklyAudit"), &ScheduledAuditProps{
//   	AccountAuditConfiguration: config,
//   	Frequency: iot.Frequency_WEEKLY,
//   	DayOfWeek: iot.DayOfWeek_SUNDAY,
//   	AuditChecks: []AuditCheck{
//   		iot.AuditCheck_CA_CERTIFICATE_EXPIRING_CHECK,
//   	},
//   })
//
//   // Monthly audit
//   monthlyAudit := iot.NewScheduledAudit(this, jsii.String("MonthlyAudit"), &ScheduledAuditProps{
//   	AccountAuditConfiguration: config,
//   	Frequency: iot.Frequency_MONTHLY,
//   	DayOfMonth: iot.DayOfMonth_Of(jsii.Number(1)),
//   	AuditChecks: []AuditCheck{
//   		iot.AuditCheck_CA_CERTIFICATE_KEY_QUALITY_CHECK,
//   	},
//   })
//
// Experimental.
type DayOfMonth interface {
	// The day of the month.
	// Experimental.
	Day() *string
}

// The jsii proxy struct for DayOfMonth
type jsiiProxy_DayOfMonth struct {
	_ byte // padding
}

func (j *jsiiProxy_DayOfMonth) Day() *string {
	var returns *string
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}


// Custom day of the month.
// Experimental.
func DayOfMonth_Of(day *float64) DayOfMonth {
	_init_.Initialize()

	if err := validateDayOfMonth_OfParameters(day); err != nil {
		panic(err)
	}
	var returns DayOfMonth

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.DayOfMonth",
		"of",
		[]interface{}{day},
		&returns,
	)

	return returns
}

func DayOfMonth_LAST_DAY() DayOfMonth {
	_init_.Initialize()
	var returns DayOfMonth
	_jsii_.StaticGet(
		"@aws-cdk/aws-iot-alpha.DayOfMonth",
		"LAST_DAY",
		&returns,
	)
	return returns
}


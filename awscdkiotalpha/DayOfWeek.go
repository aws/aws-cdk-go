package awscdkiotalpha


// The day of the week on which the scheduled audit takes place.
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
type DayOfWeek string

const (
	// Sunday.
	// Experimental.
	DayOfWeek_SUNDAY DayOfWeek = "SUNDAY"
	// Monday.
	// Experimental.
	DayOfWeek_MONDAY DayOfWeek = "MONDAY"
	// Tuesday.
	// Experimental.
	DayOfWeek_TUESDAY DayOfWeek = "TUESDAY"
	// Wednesday.
	// Experimental.
	DayOfWeek_WEDNESDAY DayOfWeek = "WEDNESDAY"
	// Thursday.
	// Experimental.
	DayOfWeek_THURSDAY DayOfWeek = "THURSDAY"
	// Friday.
	// Experimental.
	DayOfWeek_FRIDAY DayOfWeek = "FRIDAY"
	// Saturday.
	// Experimental.
	DayOfWeek_SATURDAY DayOfWeek = "SATURDAY"
	// UNSET_VALUE.
	// Experimental.
	DayOfWeek_UNSET_VALUE DayOfWeek = "UNSET_VALUE"
)


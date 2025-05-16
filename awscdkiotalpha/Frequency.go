package awscdkiotalpha


// The frequency at which the scheduled audit takes place.
//
// Example:
//   var config accountAuditConfiguration
//
//
//   // Daily audit
//   dailyAudit := iot.NewScheduledAudit(this, jsii.String("DailyAudit"), &ScheduledAuditProps{
//   	AccountAuditConfiguration: config,
//   	Frequency: iot.Frequency_DAILY,
//   	AuditChecks: []auditCheck{
//   		iot.*auditCheck_AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK,
//   	},
//   })
//
//   // Weekly audit
//   weeklyAudit := iot.NewScheduledAudit(this, jsii.String("WeeklyAudit"), &ScheduledAuditProps{
//   	AccountAuditConfiguration: config,
//   	Frequency: iot.Frequency_WEEKLY,
//   	DayOfWeek: iot.DayOfWeek_SUNDAY,
//   	AuditChecks: []*auditCheck{
//   		iot.*auditCheck_CA_CERTIFICATE_EXPIRING_CHECK,
//   	},
//   })
//
//   // Monthly audit
//   monthlyAudit := iot.NewScheduledAudit(this, jsii.String("MonthlyAudit"), &ScheduledAuditProps{
//   	AccountAuditConfiguration: config,
//   	Frequency: iot.Frequency_MONTHLY,
//   	DayOfMonth: iot.DayOfMonth_Of(jsii.Number(1)),
//   	AuditChecks: []*auditCheck{
//   		iot.*auditCheck_CA_CERTIFICATE_KEY_QUALITY_CHECK,
//   	},
//   })
//
// Experimental.
type Frequency string

const (
	// Daily.
	// Experimental.
	Frequency_DAILY Frequency = "DAILY"
	// Weekly.
	// Experimental.
	Frequency_WEEKLY Frequency = "WEEKLY"
	// Bi-weekly.
	// Experimental.
	Frequency_BI_WEEKLY Frequency = "BI_WEEKLY"
	// Monthly.
	// Experimental.
	Frequency_MONTHLY Frequency = "MONTHLY"
)


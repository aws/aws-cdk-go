package awscdkiotalpha


// Properties for defining AWS IoT Scheduled Audit.
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
type ScheduledAuditProps struct {
	// Account audit configuration.
	//
	// The audit checks specified in `auditChecks` must be enabled in this configuration.
	// Experimental.
	AccountAuditConfiguration IAccountAuditConfiguration `field:"required" json:"accountAuditConfiguration" yaml:"accountAuditConfiguration"`
	// Which checks are performed during the scheduled audit.
	//
	// Checks must be enabled for your account.
	// Experimental.
	AuditChecks *[]AuditCheck `field:"required" json:"auditChecks" yaml:"auditChecks"`
	// How often the scheduled audit occurs.
	// Experimental.
	Frequency Frequency `field:"required" json:"frequency" yaml:"frequency"`
	// The day of the month on which the scheduled audit is run (if the frequency is "MONTHLY").
	//
	// If days 29-31 are specified, and the month does not have that many days, the audit takes place on the "LAST" day of the month.
	// Default: - required if frequency is "MONTHLY", not allowed otherwise.
	//
	// Experimental.
	DayOfMonth DayOfMonth `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// The day of the week on which the scheduled audit is run (if the frequency is "WEEKLY" or "BIWEEKLY").
	// Default: - required if frequency is "WEEKLY" or "BIWEEKLY", not allowed otherwise.
	//
	// Experimental.
	DayOfWeek DayOfWeek `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The name of the scheduled audit.
	// Default: - auto generated name.
	//
	// Experimental.
	ScheduledAuditName *string `field:"optional" json:"scheduledAuditName" yaml:"scheduledAuditName"`
}


package awscdkiotalpha


// The AWS IoT Device Defender audit checks.
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
// See: https://docs.aws.amazon.com/iot-device-defender/latest/devguide/device-defender-audit-checks.html
//
// Experimental.
type AuditCheck string

const (
	// Checks the permissiveness of an authenticated Amazon Cognito identity pool role.
	//
	// For this check, AWS IoT Device Defender audits all Amazon Cognito identity pools that have been used to connect to the AWS IoT message broker
	// during the 31 days before the audit is performed.
	// Experimental.
	AuditCheck_AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK AuditCheck = "AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK"
	// Checks if a CA certificate is expiring.
	//
	// This check applies to CA certificates expiring within 30 days or that have expired.
	// Experimental.
	AuditCheck_CA_CERTIFICATE_EXPIRING_CHECK AuditCheck = "CA_CERTIFICATE_EXPIRING_CHECK"
	// Checks the quality of the CA certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, and if the key meets a minimum required size.
	//
	// This check applies to CA certificates that are ACTIVE or PENDING_TRANSFER.
	// Experimental.
	AuditCheck_CA_CERTIFICATE_KEY_QUALITY_CHECK AuditCheck = "CA_CERTIFICATE_KEY_QUALITY_CHECK"
	// Checks if multiple devices connect using the same client ID.
	// Experimental.
	AuditCheck_CONFLICTING_CLIENT_IDS_CHECK AuditCheck = "CONFLICTING_CLIENT_IDS_CHECK"
	// Checks if a device certificate is expiring.
	//
	// This check applies to device certificates expiring within 30 days or that have expired.
	// Experimental.
	AuditCheck_DEVICE_CERTIFICATE_EXPIRING_CHECK AuditCheck = "DEVICE_CERTIFICATE_EXPIRING_CHECK"
	// Checks the quality of the device certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, signed by a registered certificate authority,
	// and if the key meets a minimum required size.
	// Experimental.
	AuditCheck_DEVICE_CERTIFICATE_KEY_QUALITY_CHECK AuditCheck = "DEVICE_CERTIFICATE_KEY_QUALITY_CHECK"
	// Checks if multiple concurrent connections use the same X.509 certificate to authenticate with AWS IoT.
	// Experimental.
	AuditCheck_DEVICE_CERTIFICATE_SHARED_CHECK AuditCheck = "DEVICE_CERTIFICATE_SHARED_CHECK"
	// Checks the permissiveness of a policy attached to an authenticated Amazon Cognito identity pool role.
	// Experimental.
	AuditCheck_IOT_POLICY_OVERLY_PERMISSIVE_CHECK AuditCheck = "IOT_POLICY_OVERLY_PERMISSIVE_CHECK"
	// Checks if a role alias has access to services that haven't been used for the AWS IoT device in the last year.
	// Experimental.
	AuditCheck_IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK AuditCheck = "IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK"
	// Checks if the temporary credentials provided by AWS IoT role aliases are overly permissive.
	// Experimental.
	AuditCheck_IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK AuditCheck = "IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK"
	// Checks if AWS IoT logs are disabled.
	// Experimental.
	AuditCheck_LOGGING_DISABLED_CHECK AuditCheck = "LOGGING_DISABLED_CHECK"
	// Checks if a revoked CA certificate is still active.
	// Experimental.
	AuditCheck_REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK AuditCheck = "REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK"
	// Checks if a revoked device certificate is still active.
	// Experimental.
	AuditCheck_REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK AuditCheck = "REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK"
	// Checks if policy attached to an unauthenticated Amazon Cognito identity pool role is too permissive.
	// Experimental.
	AuditCheck_UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK AuditCheck = "UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK"
)


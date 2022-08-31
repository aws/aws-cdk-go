package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnScheduledAudit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledAuditProps := &cfnScheduledAuditProps{
//   	frequency: jsii.String("frequency"),
//   	targetCheckNames: []*string{
//   		jsii.String("targetCheckNames"),
//   	},
//
//   	// the properties below are optional
//   	dayOfMonth: jsii.String("dayOfMonth"),
//   	dayOfWeek: jsii.String("dayOfWeek"),
//   	scheduledAuditName: jsii.String("scheduledAuditName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnScheduledAuditProps struct {
	// How often the scheduled audit occurs.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Which checks are performed during the scheduled audit.
	//
	// Checks must be enabled for your account. (Use `DescribeAccountAuditConfiguration` to see the list of all checks, including those that are enabled or use `UpdateAccountAuditConfiguration` to select which checks are enabled.)
	//
	// The following checks are currently aviable:
	//
	// - `AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`
	// - `CA_CERTIFICATE_EXPIRING_CHECK`
	// - `CA_CERTIFICATE_KEY_QUALITY_CHECK`
	// - `CONFLICTING_CLIENT_IDS_CHECK`
	// - `DEVICE_CERTIFICATE_EXPIRING_CHECK`
	// - `DEVICE_CERTIFICATE_KEY_QUALITY_CHECK`
	// - `DEVICE_CERTIFICATE_SHARED_CHECK`
	// - `IOT_POLICY_OVERLY_PERMISSIVE_CHECK`
	// - `IOT_ROLE_ALIAS_ALLOWS_ACCESS_TO_UNUSED_SERVICES_CHECK`
	// - `IOT_ROLE_ALIAS_OVERLY_PERMISSIVE_CHECK`
	// - `LOGGING_DISABLED_CHECK`
	// - `REVOKED_CA_CERTIFICATE_STILL_ACTIVE_CHECK`
	// - `REVOKED_DEVICE_CERTIFICATE_STILL_ACTIVE_CHECK`
	// - `UNAUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK`.
	TargetCheckNames *[]*string `field:"required" json:"targetCheckNames" yaml:"targetCheckNames"`
	// The day of the month on which the scheduled audit is run (if the `frequency` is "MONTHLY").
	//
	// If days 29-31 are specified, and the month does not have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// The day of the week on which the scheduled audit is run (if the `frequency` is "WEEKLY" or "BIWEEKLY").
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The name of the scheduled audit.
	ScheduledAuditName *string `field:"optional" json:"scheduledAuditName" yaml:"scheduledAuditName"`
	// Metadata that can be used to manage the scheduled audit.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


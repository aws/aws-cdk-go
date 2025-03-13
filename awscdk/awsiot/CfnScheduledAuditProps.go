package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnScheduledAudit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledAuditProps := &CfnScheduledAuditProps{
//   	Frequency: jsii.String("frequency"),
//   	TargetCheckNames: []*string{
//   		jsii.String("targetCheckNames"),
//   	},
//
//   	// the properties below are optional
//   	DayOfMonth: jsii.String("dayOfMonth"),
//   	DayOfWeek: jsii.String("dayOfWeek"),
//   	ScheduledAuditName: jsii.String("scheduledAuditName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-scheduledaudit.html
//
type CfnScheduledAuditProps struct {
	// How often the scheduled audit occurs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-scheduledaudit.html#cfn-iot-scheduledaudit-frequency
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-scheduledaudit.html#cfn-iot-scheduledaudit-targetchecknames
	//
	TargetCheckNames *[]*string `field:"required" json:"targetCheckNames" yaml:"targetCheckNames"`
	// The day of the month on which the scheduled audit is run (if the `frequency` is "MONTHLY").
	//
	// If days 29-31 are specified, and the month does not have that many days, the audit takes place on the "LAST" day of the month.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-scheduledaudit.html#cfn-iot-scheduledaudit-dayofmonth
	//
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// The day of the week on which the scheduled audit is run (if the `frequency` is "WEEKLY" or "BIWEEKLY").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-scheduledaudit.html#cfn-iot-scheduledaudit-dayofweek
	//
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The name of the scheduled audit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-scheduledaudit.html#cfn-iot-scheduledaudit-scheduledauditname
	//
	ScheduledAuditName *string `field:"optional" json:"scheduledAuditName" yaml:"scheduledAuditName"`
	// Metadata that can be used to manage the scheduled audit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-scheduledaudit.html#cfn-iot-scheduledaudit-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


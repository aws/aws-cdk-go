# AWS IoT Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

AWS IoT Core lets you connect billions of IoT devices and route trillions of
messages to AWS services without managing infrastructure.

## `TopicRule`

Create a topic rule that give your devices the ability to interact with AWS services.
You can create a topic rule with an action that invoke the Lambda action as following:

```go
func := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String(`
	    exports.handler = (event) => {
	      console.log("It is test for lambda action of AWS IoT Rule.", event);
	    };`)),
})

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	TopicRuleName: jsii.String("MyTopicRule"),
	 // optional
	Description: jsii.String("invokes the lambda function"),
	 // optional
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	Actions: []IAction{
		actions.NewLambdaFunctionAction(func),
	},
})
```

Or, you can add an action after constructing the `TopicRule` instance as following:

```go
var func Function


topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
})
topicRule.AddAction(actions.NewLambdaFunctionAction(func))
```

You can also supply `errorAction` as following,
and the IoT Rule will trigger it if a rule's action is unable to perform:

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	ErrorAction: actions.NewCloudWatchLogsAction(logGroup),
})
```

If you wanna make the topic rule disable, add property `enabled: false` as following:

```go
iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	Enabled: jsii.Boolean(false),
})
```

See also [@aws-cdk/aws-iot-actions-alpha](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-iot-actions-alpha-readme.html) for other actions.

## Logging

AWS IoT provides a [logging feature](https://docs.aws.amazon.com/iot/latest/developerguide/configure-logging.html) that allows you to monitor and log AWS IoT activity.

You can enable IoT logging with the following code:

```go
iot.NewLogging(this, jsii.String("Logging"), &LoggingProps{
	LogLevel: iot.LogLevel_INFO,
})
```

**Note**: All logs are forwarded to the `AWSIotLogsV2` log group in CloudWatch.

## Audit

An [AWS IoT Device Defender audit looks](https://docs.aws.amazon.com/iot-device-defender/latest/devguide/device-defender-audit.html) at account- and device-related settings and policies to ensure security measures are in place.
An audit can help you detect any drifts from security best practices or access policies.

### Account Audit Configuration

The IoT audit includes [various audit checks](https://docs.aws.amazon.com/iot-device-defender/latest/devguide/device-defender-audit-checks.html), and it is necessary to configure settings to enable those checks.

You can enable an account audit configuration with the following code:

```go
// Audit notification are sent to the SNS topic
var targetTopic ITopic


iot.NewAccountAuditConfiguration(this, jsii.String("AuditConfiguration"), &AccountAuditConfigurationProps{
	TargetTopic: TargetTopic,
})
```

By default, all audit checks are enabled, but it is also possible to enable only specific audit checks.

```go
iot.NewAccountAuditConfiguration(this, jsii.String("AuditConfiguration"), &AccountAuditConfigurationProps{
	CheckConfiguration: &CheckConfiguration{
		// enabled
		AuthenticatedCognitoRoleOverlyPermissiveCheck: jsii.Boolean(true),
		// enabled by default
		CaCertificateExpiringCheck: undefined,
		// disabled
		CaCertificateKeyQualityCheck: jsii.Boolean(false),
		ConflictingClientIdsCheck: jsii.Boolean(false),
		DeviceCertificateAgeCheck: jsii.Boolean(false),
		DeviceCertificateExpiringCheck: jsii.Boolean(false),
		DeviceCertificateKeyQualityCheck: jsii.Boolean(false),
		DeviceCertificateSharedCheck: jsii.Boolean(false),
		IntermediateCaRevokedForActiveDeviceCertificatesCheck: jsii.Boolean(false),
		IoTPolicyPotentialMisConfigurationCheck: jsii.Boolean(false),
		IotPolicyOverlyPermissiveCheck: jsii.Boolean(false),
		IotRoleAliasAllowsAccessToUnusedServicesCheck: jsii.Boolean(false),
		IotRoleAliasOverlyPermissiveCheck: jsii.Boolean(false),
		LoggingDisabledCheck: jsii.Boolean(false),
		RevokedCaCertificateStillActiveCheck: jsii.Boolean(false),
		RevokedDeviceCertificateStillActiveCheck: jsii.Boolean(false),
		UnauthenticatedCognitoRoleOverlyPermissiveCheck: jsii.Boolean(false),
	},
})
```

To configure [the device certificate age check](https://docs.aws.amazon.com/iot-device-defender/latest/devguide/device-certificate-age-check.html), you can specify the duration for the check:

```go
import "github.com/aws/aws-cdk-go/awscdk"


iot.NewAccountAuditConfiguration(this, jsii.String("AuditConfiguration"), &AccountAuditConfigurationProps{
	CheckConfiguration: &CheckConfiguration{
		DeviceCertificateAgeCheck: jsii.Boolean(true),
		// The default value is 365 days
		// Valid values range from 30 days (minimum) to 3650 days (10 years, maximum)
		DeviceCertificateAgeCheckDuration: awscdk.Duration_Days(jsii.Number(365)),
	},
})
```

### Scheduled Audit

You can create a [scheduled audit](https://docs.aws.amazon.com/iot-device-defender/latest/devguide/AuditCommands.html#device-defender-AuditCommandsManageSchedules) that is run at a specified time interval. Checks must be enabled for your account by creating `AccountAuditConfiguration`.

```go
var config AccountAuditConfiguration


// Daily audit
dailyAudit := iot.NewScheduledAudit(this, jsii.String("DailyAudit"), &ScheduledAuditProps{
	AccountAuditConfiguration: config,
	Frequency: iot.Frequency_DAILY,
	AuditChecks: []AuditCheck{
		iot.AuditCheck_AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK,
	},
})

// Weekly audit
weeklyAudit := iot.NewScheduledAudit(this, jsii.String("WeeklyAudit"), &ScheduledAuditProps{
	AccountAuditConfiguration: config,
	Frequency: iot.Frequency_WEEKLY,
	DayOfWeek: iot.DayOfWeek_SUNDAY,
	AuditChecks: []AuditCheck{
		iot.AuditCheck_CA_CERTIFICATE_EXPIRING_CHECK,
	},
})

// Monthly audit
monthlyAudit := iot.NewScheduledAudit(this, jsii.String("MonthlyAudit"), &ScheduledAuditProps{
	AccountAuditConfiguration: config,
	Frequency: iot.Frequency_MONTHLY,
	DayOfMonth: iot.DayOfMonth_Of(jsii.Number(1)),
	AuditChecks: []AuditCheck{
		iot.AuditCheck_CA_CERTIFICATE_KEY_QUALITY_CHECK,
	},
})
```

# Amazon Simple Email Service Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Email receiving

Create a receipt rule set with rules and actions (actions can be found in the
`aws-cdk-lib/aws-ses-actions` package):

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


bucket := s3.NewBucket(this, jsii.String("Bucket"))
topic := sns.NewTopic(this, jsii.String("Topic"))

ses.NewReceiptRuleSet(this, jsii.String("RuleSet"), &ReceiptRuleSetProps{
	Rules: []receiptRuleOptions{
		&receiptRuleOptions{
			Recipients: []*string{
				jsii.String("hello@aws.com"),
			},
			Actions: []iReceiptRuleAction{
				actions.NewAddHeader(&AddHeaderProps{
					Name: jsii.String("X-Special-Header"),
					Value: jsii.String("aws"),
				}),
				actions.NewS3(&S3Props{
					Bucket: *Bucket,
					ObjectKeyPrefix: jsii.String("emails/"),
					Topic: *Topic,
				}),
			},
		},
		&receiptRuleOptions{
			Recipients: []*string{
				jsii.String("aws.com"),
			},
			Actions: []*iReceiptRuleAction{
				actions.NewSns(&SnsProps{
					Topic: *Topic,
				}),
			},
		},
	},
})
```

Alternatively, rules can be added to a rule set:

```go
ruleSet := ses.NewReceiptRuleSet(this, jsii.String("RuleSet"))

awsRule := ruleSet.addRule(jsii.String("Aws"), &ReceiptRuleOptions{
	Recipients: []*string{
		jsii.String("aws.com"),
	},
})
```

And actions to rules:

```go
import actions "github.com/aws/aws-cdk-go/awscdk"

var awsRule receiptRule
var topic topic

awsRule.AddAction(actions.NewSns(&SnsProps{
	Topic: Topic,
}))
```

When using `addRule`, the new rule is added after the last added rule unless `after` is specified.

### Drop spams

A rule to drop spam can be added by setting `dropSpam` to `true`:

```go
ses.NewReceiptRuleSet(this, jsii.String("RuleSet"), &ReceiptRuleSetProps{
	DropSpam: jsii.Boolean(true),
})
```

This will add a rule at the top of the rule set with a Lambda action that stops processing messages that have at least one spam indicator. See [Lambda Function Examples](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-lambda-example-functions.html).

### Receipt filter

Create a receipt filter:

```go
ses.NewReceiptFilter(this, jsii.String("Filter"), &ReceiptFilterProps{
	Ip: jsii.String("1.2.3.4/16"),
})
```

An allow list filter is also available:

```go
ses.NewAllowListReceiptFilter(this, jsii.String("AllowList"), &AllowListReceiptFilterProps{
	Ips: []*string{
		jsii.String("10.0.0.0/16"),
		jsii.String("1.2.3.4/16"),
	},
})
```

This will first create a block all filter and then create allow filters for the listed ip addresses.

## Email sending

### Dedicated IP pools

When you create a new Amazon SES account, your emails are sent from IP addresses that are shared with other
Amazon SES users. For [an additional monthly charge](https://aws.amazon.com/ses/pricing/), you can lease
dedicated IP addresses that are reserved for your exclusive use.

Use the DedicatedIpPool construct to create a pool of dedicated IP addresses. When specifying a name for your dedicated IP pool, ensure that it adheres to the following naming convention:

* The name must include only lowercase letters (a-z), numbers (0-9), underscores (_), and hyphens (-).
* The name must not exceed 64 characters in length.

```go
ses.NewDedicatedIpPool(this, jsii.String("Pool"), &DedicatedIpPoolProps{
	DedicatedIpPoolName: jsii.String("mypool"),
	ScalingMode: ses.ScalingMode_STANDARD,
})
```

The pool can then be used in a configuration set. If the provided dedicatedIpPoolName does not follow the specified naming convention, an error will be thrown.

### Configuration sets

Configuration sets are groups of rules that you can apply to your verified identities. A verified identity is
a domain, subdomain, or email address you use to send email through Amazon SES. When you apply a configuration
set to an email, all of the rules in that configuration set are applied to the email.

Use the `ConfigurationSet` construct to create a configuration set:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var myPool iDedicatedIpPool


ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
	CustomTrackingRedirectDomain: jsii.String("track.cdk.dev"),
	TlsPolicy: ses.ConfigurationSetTlsPolicy_REQUIRE,
	DedicatedIpPool: myPool,
	// Specify maximum delivery time
	// This configuration can be useful in such cases as time-sensitive emails (like those containing a one-time-password),
	// transactional emails, and email that you want to ensure isn't delivered during non-business hours.
	MaxDeliveryDuration: awscdk.Duration_Minutes(jsii.Number(10)),
})
```

Use `addEventDestination()` to publish email sending events to Amazon SNS, Amazon CloudWatch, Amazon Data Firehose or Amazon EventBridge:

```go
var myConfigurationSet configurationSet
var myTopic topic


myConfigurationSet.AddEventDestination(jsii.String("ToSns"), &ConfigurationSetEventDestinationOptions{
	Destination: ses.EventDestination_SnsTopic(myTopic),
})
```

**Note**: For EventBridge, you must specify the default EventBus:

```go
import events "github.com/aws/aws-cdk-go/awscdk"

var myConfigurationSet configurationSet


bus := events.EventBus_FromEventBusName(this, jsii.String("EventBus"), jsii.String("default"))

myConfigurationSet.AddEventDestination(jsii.String("ToEventBus"), &ConfigurationSetEventDestinationOptions{
	Destination: ses.EventDestination_EventBus(bus),
})
```

For Firehose, if you don't specify IAM Role ARN for Amazon SES to send events. An IAM Role will be created automatically following https://docs.aws.amazon.com/ses/latest/dg/event-publishing-add-event-destination-firehose.html.

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
import firehose "github.com/aws/aws-cdk-go/awscdk"

var myConfigurationSet configurationSet
var firehoseDeliveryStream iDeliveryStream
var iamRole iRole


// Create IAM Role automatically
myConfigurationSet.AddEventDestination(jsii.String("ToFirehose"), &ConfigurationSetEventDestinationOptions{
	Destination: ses.EventDestination_FirehoseDeliveryStream(&FirehoseDeliveryStreamDestination{
		DeliveryStream: firehoseDeliveryStream,
	}),
})

// Specify your IAM Role
myConfigurationSet.AddEventDestination(jsii.String("ToFirehose"), &ConfigurationSetEventDestinationOptions{
	Destination: ses.EventDestination_*FirehoseDeliveryStream(&FirehoseDeliveryStreamDestination{
		DeliveryStream: firehoseDeliveryStream,
		Role: iamRole,
	}),
})
```

### Override account-level suppression list settings

You can customize account-level suppression list separately for different configuration sets by overriding it
with configuration set-level suppression.

For details, see [Using configuration set-level suppression to override your account-level suppression list](https://docs.aws.amazon.com/ses/latest/dg/sending-email-suppression-list-config-level.html).

By default, the configuration set uses your account-level suppression list settings.

You can disable account-level suppression list by specifying `disableSuppressionList` to true. Email sent with this configuration set will not use any suppression settings at all.

```go
ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
	DisableSuppressionList: jsii.Boolean(true),
})
```

You can also override account level settings with configuration set-level suppression enabled. Email sent with this configuration set will only use the suppression conditions you enabled for it (bounces, complaints, or bounces and complaints) - regardless of what your account-level suppression list settings are, it will override them.

```go
// Only bounces will be suppressed.
// Only bounces will be suppressed.
ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
	SuppressionReasons: ses.SuppressionReasons_BOUNCES_ONLY,
})

// Only complaints will be suppressed.
// Only complaints will be suppressed.
ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
	SuppressionReasons: ses.SuppressionReasons_COMPLAINTS_ONLY,
})

// Both bounces and complaints will be suppressed.
// Both bounces and complaints will be suppressed.
ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
	SuppressionReasons: ses.SuppressionReasons_BOUNCES_AND_COMPLAINTS,
})
```

### Email identity

In Amazon SES, a verified identity is a domain or email address that you use to send or receive email. Before you
can send an email using Amazon SES, you must create and verify each identity that you're going to use as a `From`,
`Source`, `Sender`, or `Return-Path` address. Verifying an identity with Amazon SES confirms that you own it and
helps prevent unauthorized use.

To verify an identity for a hosted zone, you create an `EmailIdentity`:

```go
var myHostedZone iPublicHostedZone


identity := ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
	Identity: ses.Identity_PublicHostedZone(myHostedZone),
	MailFromDomain: jsii.String("mail.cdk.dev"),
})
```

By default, [Easy DKIM](https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html) with
a 2048-bit DKIM key is used.

You can instead configure DKIM authentication by using your own public-private key pair. This process is known
as [Bring Your Own DKIM (BYODKIM)](https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-bring-your-own.html):

```go
var myHostedZone iPublicHostedZone


ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
	Identity: ses.Identity_PublicHostedZone(myHostedZone),
	DkimIdentity: ses.DkimIdentity_ByoDkim(&ByoDkimOptions{
		PrivateKey: awscdk.SecretValue_SecretsManager(jsii.String("dkim-private-key")),
		PublicKey: jsii.String("...base64-encoded-public-key..."),
		Selector: jsii.String("selector"),
	}),
})
```

When using `publicHostedZone()` for the identity, all necessary Amazon Route 53 records are created automatically:

* CNAME records for Easy DKIM
* TXT record for BYOD DKIM
* MX and TXT records for the custom MAIL FROM

When working with `domain()`, records must be created manually:

```go
identity := ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
	Identity: ses.Identity_Domain(jsii.String("cdk.dev")),
})

for _, record := range identity.DkimRecords {}
```

#### Grants

To grant a specific action to a principal use the `grant` method.
For sending emails, `grantSendEmail` can be used instead:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
var user user


identity := ses.NewEmailIdentity(this, jsii.String("Identity"), &EmailIdentityProps{
	Identity: ses.Identity_Domain(jsii.String("cdk.dev")),
})

identity.grantSendEmail(user)
```

You can also reference an existing email identity using its ARN and grant permissions to it:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
var user user


// Imports an existing email identity using its ARN.
// This is one way to reference an existing identity; another option is using its name via fromEmailIdentityName.
importedIdentity := ses.EmailIdentity_FromEmailIdentityArn(this, jsii.String("ImportedIdentity"), jsii.String("arn:aws:ses:us-east-1:123456789012:identity/example.com"))

// Grant send email permission to the imported identity
importedIdentity.GrantSendEmail(user)
```

### Virtual Deliverability Manager (VDM)

Virtual Deliverability Manager is an Amazon SES feature that helps you enhance email deliverability,
like increasing inbox deliverability and email conversions, by providing insights into your sending
and delivery data, and giving advice on how to fix the issues that are negatively affecting your
delivery success rate and reputation.

Use the `VdmAttributes` construct to configure the Virtual Deliverability Manager for your account:

```go
// Enables engagement tracking and optimized shared delivery by default
// Enables engagement tracking and optimized shared delivery by default
ses.NewVdmAttributes(this, jsii.String("Vdm"))
```

If you want to override the VDM settings in the specified configuration set, use `vdmOptions` in the `ConfigurationSet` construct.

> **Note:** The configuration set level settings need to be used together with the account level settings. (To set the account level settings using CDK, use the `VdmAttributes` Construct.)
> If you enable only the configuration set level settings, VDM will not be enabled until the account level settings are configured.
> For more information, see [Virtual Deliverability Manager settings](https://docs.aws.amazon.com/ses/latest/dg/vdm-settings.html).

```go
ses.NewConfigurationSet(this, jsii.String("ConfigurationSetWithVdmOptions"), &ConfigurationSetProps{
	VdmOptions: &VdmOptions{
		EngagementMetrics: jsii.Boolean(true),
		OptimizedSharedDelivery: jsii.Boolean(true),
	},
})
```

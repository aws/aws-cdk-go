# Amazon Simple Email Service Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Email receiving

Create a receipt rule set with rules and actions (actions can be found in the
`@aws-cdk/aws-ses-actions` package):

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"
import actions "github.com/aws/aws-cdk-go/awscdk"


bucket := s3.NewBucket(this, jsii.String("Bucket"))
topic := sns.NewTopic(this, jsii.String("Topic"))

ses.NewReceiptRuleSet(this, jsii.String("RuleSet"), &receiptRuleSetProps{
	rules: []receiptRuleOptions{
		&receiptRuleOptions{
			recipients: []*string{
				jsii.String("hello@aws.com"),
			},
			actions: []iReceiptRuleAction{
				actions.NewAddHeader(&addHeaderProps{
					name: jsii.String("X-Special-Header"),
					value: jsii.String("aws"),
				}),
				actions.NewS3(&s3Props{
					bucket: bucket,
					objectKeyPrefix: jsii.String("emails/"),
					topic: topic,
				}),
			},
		},
		&receiptRuleOptions{
			recipients: []*string{
				jsii.String("aws.com"),
			},
			actions: []*iReceiptRuleAction{
				actions.NewSns(&snsProps{
					topic: topic,
				}),
			},
		},
	},
})
```

Alternatively, rules can be added to a rule set:

```go
ruleSet := ses.NewReceiptRuleSet(this, jsii.String("RuleSet"))

awsRule := ruleSet.addRule(jsii.String("Aws"), &receiptRuleOptions{
	recipients: []*string{
		jsii.String("aws.com"),
	},
})
```

And actions to rules:

```go
import actions "github.com/aws/aws-cdk-go/awscdk"

var awsRule receiptRule
var topic topic

awsRule.addAction(actions.NewSns(&snsProps{
	topic: topic,
}))
```

When using `addRule`, the new rule is added after the last added rule unless `after` is specified.

### Drop spams

A rule to drop spam can be added by setting `dropSpam` to `true`:

```go
ses.NewReceiptRuleSet(this, jsii.String("RuleSet"), &receiptRuleSetProps{
	dropSpam: jsii.Boolean(true),
})
```

This will add a rule at the top of the rule set with a Lambda action that stops processing messages that have at least one spam indicator. See [Lambda Function Examples](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-lambda-example-functions.html).

### Receipt filter

Create a receipt filter:

```go
ses.NewReceiptFilter(this, jsii.String("Filter"), &receiptFilterProps{
	ip: jsii.String("1.2.3.4/16"),
})
```

An allow list filter is also available:

```go
ses.NewAllowListReceiptFilter(this, jsii.String("AllowList"), &allowListReceiptFilterProps{
	ips: []*string{
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

Use the `DedicatedIpPool` construct to create a pool of dedicated IP addresses:

```go
ses.NewDedicatedIpPool(this, jsii.String("Pool"))
```

The pool can then be used in a configuration set.

### Configuration sets

Configuration sets are groups of rules that you can apply to your verified identities. A verified identity is
a domain, subdomain, or email address you use to send email through Amazon SES. When you apply a configuration
set to an email, all of the rules in that configuration set are applied to the email.

Use the `ConfigurationSet` construct to create a configuration set:

```go
var myPool iDedicatedIpPool


ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &configurationSetProps{
	customTrackingRedirectDomain: jsii.String("track.cdk.dev"),
	suppressionReasons: ses.suppressionReasons_COMPLAINTS_ONLY,
	tlsPolicy: ses.configurationSetTlsPolicy_REQUIRE,
	dedicatedIpPool: myPool,
})
```

### Email identity

In Amazon SES, a verified identity is a domain or email address that you use to send or receive email. Before you
can send an email using Amazon SES, you must create and verify each identity that you're going to use as a `From`,
`Source`, `Sender`, or `Return-Path` address. Verifying an identity with Amazon SES confirms that you own it and
helps prevent unauthorized use.

To verify an identity for a hosted zone, you create an `EmailIdentity`:

```go
// Example automatically generated from non-compiling source. May contain errors.
var myHostedZone iPublicHostedZone


identity := ses.NewEmailIdentity(stack, jsii.String("Identity"), &emailIdentityProps{
	identity: ses.identity.publicHostedZone(myHostedZone),
	mailFromDomain: jsii.String("mail.cdk.dev"),
})
```

By default, [Easy DKIM](https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html) with
a 2048-bit DKIM key is used.

You can instead configure DKIM authentication by using your own public-private key pair. This process is known
as [Bring Your Own DKIM (BYODKIM)](https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-bring-your-own.html):

```go
// Example automatically generated from non-compiling source. May contain errors.
var myHostedZone iPublicHostedZone


ses.NewEmailIdentity(stack, jsii.String("Identity"), &emailIdentityProps{
	identity: ses.identity.publicHostedZone(myHostedZone),
	dkimIdentity: dkimIdentity_ByoDkim(map[string]interface{}{
		"privateKey": SecretValue_secretsManager(jsii.String("dkim-private-key")),
		"publicKey": jsii.String("...base64-encoded-public-key..."),
		"selector": jsii.String("selector"),
	}),
})
```

When using `publicHostedZone()` for the identity, all necessary Amazon Route 53 records are created automatically:

* CNAME records for Easy DKIM
* TXT record for BYOD DKIM
* MX and TXT records for the custom MAIL FROM

When working with `domain()`, records must be created manually:

```go
// Example automatically generated from non-compiling source. May contain errors.
identity := ses.NewEmailIdentity(stack, jsii.String("Identity"), &emailIdentityProps{
	identity: ses.identity.domain(jsii.String("cdk.dev")),
})

for _, record := range identity.dkimRecords {}
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

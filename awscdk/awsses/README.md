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

## Receipt filter

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

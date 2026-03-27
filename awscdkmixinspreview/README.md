# CDK Mixins

> **Note**: The core Mixins mechanism  is now GA  and available in `constructs` and `aws-cdk-lib` (`Mixins`, `Mixin`, `IMixin`, `MixinApplicator`, `ConstructSelector`).
> All service Mixins are now available in `aws-cdk-lib`.
> Please update your imports.
>
> This package continues to provide **Logs Delivery Mixins** and **EventBridge Event Facades**, which are still experimental.

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

CDK Mixins provide a new, advanced way to add functionality through composable abstractions.
Unlike traditional L2 constructs that bundle all features together, Mixins allow you to pick and choose exactly the capabilities you need for constructs.
Mixins can be applied during or after construct construction.
Mixins are an *addition*, *not* a replacement for construct properties.
By itself, they cannot change optionality of properties or change defaults.

## Usage and documentation

See the [documentation for CDK Mixins](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib-readme.html#mixins) in  `aws-cdk-lib`.

### Built-in Mixins

### Logs Delivery

Configures vended logs delivery for supported resources to various destinations:

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"

// Create CloudFront distribution
var origin IBucket

distribution := cloudfront.NewDistribution(scope, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(origin),
	},
})

// Create log destination
logGroup := logs.NewLogGroup(scope, jsii.String("DeliveryLogGroup"))

// Configure log delivery using the mixin
distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToLogGroup(logGroup, &CfnDistributionConnectionLogsLogGroupProps{
	OutputFormat: cloudfrontMixins.CfnDistributionConnectionLogsOutputFormat.LogGroup_JSON,
	RecordFields: []CfnDistributionConnectionLogsRecordFields{
		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_CONNECTIONSTATUS,
		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_CLIENTIP,
		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_SERVERIP,
		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_TLSPROTOCOL,
	},
}))
```

Configures vended logs delivery for supported resources when a pre-created destination is provided:

```go
import cloudfrontMixins "github.com/aws/aws-cdk-go/awscdkmixinspreview"

// Create CloudFront distribution
var origin IBucket

distribution := cloudfront.NewDistribution(scope, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(origin),
	},
})

// Create destination bucket
destBucket := s3.NewBucket(scope, jsii.String("DeliveryBucket"))
// Add permissions to bucket to facilitate log delivery
bucketPolicy := s3.NewBucketPolicy(scope, jsii.String("DeliveryBucketPolicy"), &BucketPolicyProps{
	Bucket: destBucket,
	Document: iam.NewPolicyDocument(),
})
// Create S3 delivery destination for logs
destination := logs.NewCfnDeliveryDestination(scope, jsii.String("Destination"), &CfnDeliveryDestinationProps{
	DestinationResourceArn: destBucket.bucketArn,
	Name: jsii.String("unique-destination-name"),
	DeliveryDestinationType: jsii.String("S3"),
})

distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToDestination(destination))
```

Vended Logs Configuration for Cross Account delivery (only supported for S3 and Firehose destinations)

```go
import logDestinations "github.com/aws/aws-cdk-go/awscdkmixinspreview"
import cloudfrontMixins "github.com/aws/aws-cdk-go/awscdkmixinspreview"

// Create CloudFront distribution
var origin IBucket


destinationAccount := "123456789012"
sourceAccount := "234567890123"
region := "us-east-1"

app := awscdk.NewApp()

destStack := awscdk.NewStack(app, jsii.String("destination-stack"), &StackProps{
	Env: &Environment{
		Account: destinationAccount,
		Region: jsii.String(*Region),
	},
})

// Create destination bucket
destBucket := s3.NewBucket(destStack, jsii.String("DeliveryBucket"))
logDestinations.NewS3DeliveryDestination(destStack, jsii.String("Destination"), &S3DeliveryDestinationProps{
	Bucket: destBucket,
	SourceAccountId: sourceAccount,
})

sourceStack := awscdk.NewStack(app, jsii.String("source-stack"), &StackProps{
	Env: &Environment{
		Account: sourceAccount,
		Region: jsii.String(*Region),
	},
})
distribution := cloudfront.NewDistribution(sourceStack, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(origin),
	},
})

destination := logs.CfnDeliveryDestination_FromDeliveryDestinationArn(sourceStack, jsii.String("Destination"), jsii.String("arn of Delivery Destination in destinationAccount"))

distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToDestination(destination))
```

---


## EventBridge Event Patterns

CDK Mixins automatically generates typed EventBridge event patterns for AWS resources. These patterns come in two flavors: **resource-specific** and **standalone**.

### Resource-Specific Event Patterns

Resource-specific patterns are created by attaching a resource reference (e.g. an S3 bucket). The resource identifier is automatically injected into the event pattern, so calling a pattern method with no arguments still filters events to that specific resource. For example, an S3 `objectCreatedPattern()` will automatically include the bucket name in the pattern, meaning it only matches events from that particular bucket.

#### Event Patterns Basic Usage

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
import "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"
var fn Function


// Works with L2 constructs
myBucket := s3.NewBucket(scope, jsii.String("Bucket"))
bucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(myBucket)

events.NewRule(scope, jsii.String("Rule"), &RuleProps{
	EventPattern: bucketEvents.ObjectCreatedPattern(&ObjectCreatedProps{
		Object: &ObjectType{
			Key: events.Match_Wildcard(jsii.String("uploads/*")),
		},
	}),
	Targets: []IRuleTarget{
		targets.NewLambdaFunction(fn),
	},
})

// Also works with L1 constructs
cfnBucket := s3.NewCfnBucket(scope, jsii.String("CfnBucket"))
cfnBucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(cfnBucket)

events.NewCfnRule(scope, jsii.String("CfnRule"), &CfnRuleProps{
	State: jsii.String("ENABLED"),
	EventPattern: cfnBucketEvents.*ObjectCreatedPattern(&ObjectCreatedProps{
		Object: &ObjectType{
			Key: events.Match_*Wildcard(jsii.String("uploads/*")),
		},
	}),
	Targets: []interface{}{
		&TargetProperty{
			Arn: fn.functionArn,
			Id: jsii.String("L1"),
		},
	},
})
```

#### Event Pattern Features

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"


bucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(bucket)

// Bucket name is automatically injected from the bucket reference
pattern := bucketEvents.ObjectCreatedPattern()
```

**Event Metadata Support**: Control EventBridge pattern metadata

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
import events "github.com/aws/aws-cdk-go/awscdk"


bucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(bucket)

pattern := bucketEvents.ObjectCreatedPattern(&ObjectCreatedProps{
	EventMetadata: &AWSEventMetadataProps{
		Region: events.Match_Prefix(jsii.String("us-")),
		Version: []*string{
			jsii.String("0"),
		},
	},
})
```

### Standalone Event Patterns

Standalone patterns are not tied to any specific resource. They match events across all resources of that type. For example, a standalone `awsAPICallViaCloudTrailPattern()` will match CloudTrail API calls for all S3 buckets in the account, not just a specific one.

#### Event Patterns Basic Usage

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
import "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var fn Function


// Works with L2 Rule
// Works with L2 Rule
events.NewRule(scope, jsii.String("Rule"), &RuleProps{
	EventPattern: awscdkmixinspreview.AWSAPICallViaCloudTrail_AwsAPICallViaCloudTrailPattern(&AWSAPICallViaCloudTrailProps{
		TlsDetails: &TlsDetails{
			TlsVersion: []*string{
				jsii.String("TLSv1.3"),
			},
		},
		EventMetadata: &AWSEventMetadataProps{
			Region: []*string{
				jsii.String("us-east-1"),
			},
		},
	}),
	Targets: []IRuleTarget{
		targets.NewLambdaFunction(fn),
	},
})

// Also works with L1 CfnRule
// Also works with L1 CfnRule
events.NewCfnRule(scope, jsii.String("CfnRule"), &CfnRuleProps{
	State: jsii.String("ENABLED"),
	EventPattern: awscdkmixinspreview.AWSAPICallViaCloudTrail_*AwsAPICallViaCloudTrailPattern(&AWSAPICallViaCloudTrailProps{
		TlsDetails: &TlsDetails{
			TlsVersion: []*string{
				jsii.String("TLSv1.3"),
			},
		},
		EventMetadata: &AWSEventMetadataProps{
			Region: []*string{
				jsii.String("us-east-1"),
			},
		},
	}),
	Targets: []interface{}{
		&TargetProperty{
			Arn: fn.functionArn,
			Id: jsii.String("L1"),
		},
	},
})
```

#### Event Pattern Features

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"


// Matches CloudTrail API calls across ALL S3 buckets
pattern := awscdkmixinspreview.AWSAPICallViaCloudTrail_AwsAPICallViaCloudTrailPattern()
```

**Event Metadata Support**: Control EventBridge pattern metadata

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
import events "github.com/aws/aws-cdk-go/awscdk"


pattern := awscdkmixinspreview.AWSAPICallViaCloudTrail_AwsAPICallViaCloudTrailPattern(&AWSAPICallViaCloudTrailProps{
	EventMetadata: &AWSEventMetadataProps{
		Region: events.Match_Prefix(jsii.String("us-")),
		Version: []*string{
			jsii.String("0"),
		},
	},
})
```

### Available Events

Event patterns are generated for EventBridge events available in the AWS Event Schema Registry. Common examples:

**S3 Events**:

* `objectCreatedPattern()` - Object creation events
* `objectDeletedPattern()` - Object deletion events
* `objectTagsAddedPattern()` - Object tagging events
* `awsAPICallViaCloudTrailPattern()` - CloudTrail API calls

Import events from service-specific modules:

```go
// Resource-specific (filters to a specific bucket)
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"

// Standalone (matches across all buckets)
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
```

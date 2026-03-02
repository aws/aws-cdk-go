# CDK Mixins (Preview)

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

> **Note**: The core Mixins mechanism (`Mixins`, `Mixin`, `IMixin`, `MixinApplicator`, `ConstructSelector`) is now available in `constructs` and `aws-cdk-lib/core`. Please update your imports.
> This package continues to provide additional preview features until they move to their final destinations.

This package provides two main features:

1. **Mixins** - Composable abstractions for adding functionality to constructs
2. **EventBridge Event Patterns** - Type-safe event patterns for AWS resources

---


CDK Mixins provide a new, advanced way to add functionality through composable abstractions.
Unlike traditional L2 constructs that bundle all features together, Mixins allow you to pick and choose exactly the capabilities you need for constructs.

## Key Benefits

CDK Mixins offer a well-defined way to build self-contained constructs features.
Mixins are applied during or after construct construction.

* **Universal Compatibility**: Apply the same abstractions to L1 constructs, L2 constructs, or custom constructs
* **Composable Design**: Mix and match features without being locked into specific implementations
* **Cross-Service Abstractions**: Use common patterns like encryption across different AWS services
* **Escape Hatch Freedom**: Customize resources in a safe, typed way while keeping the abstractions you want

Mixins are an *addition*, *not* a replacement for construct properties.
By itself, they cannot change optionality of properties or change defaults.

### Usage and documentation

See the [documentation for `aws-cdk-lib`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib-readme.html#mixins).

### Built-in Mixins

#### Cross-Service Mixins

**EncryptionAtRest**: Applies encryption to supported AWS resources

```go
// Works across different resource types
myBucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
awscdk.Mixins_Of(myBucket).Apply(NewEncryptionAtRest())

myLogGroup := logs.NewCfnLogGroup(scope, jsii.String("LogGroup"))
awscdk.Mixins_Of(myLogGroup).Apply(NewEncryptionAtRest())
```

#### S3-Specific Mixins

**AutoDeleteObjects**: Configures automatic object deletion for S3 buckets

```go
myBucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
awscdk.Mixins_Of(myBucket).Apply(awscdkmixinspreview.NewAutoDeleteObjects())
```

**BucketVersioning**: Enables versioning on S3 buckets

```go
myBucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
awscdk.Mixins_Of(myBucket).Apply(awscdkmixinspreview.NewBucketVersioning())
```

**BucketBlockPublicAccess**: Enables blocking public-access on S3 buckets

```go
myBucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
awscdk.Mixins_Of(myBucket).Apply(awscdkmixinspreview.NewBucketBlockPublicAccess())
```

**BucketPolicyStatementsMixin**: Adds IAM policy statements to a bucket policy

```go
bucketPolicy := s3.NewCfnBucketPolicy(scope, jsii.String("BucketPolicy"), &CfnBucketPolicyProps{
	Bucket: bucket,
	PolicyDocument: iam.NewPolicyDocument(),
})
awscdk.Mixins_Of(bucketPolicy).Apply(awscdkmixinspreview.NewBucketPolicyStatementsMixin([]PolicyStatement{
	iam.NewPolicyStatement(&PolicyStatementProps{
		Actions: []*string{
			jsii.String("s3:GetObject"),
		},
		Resources: []*string{
			jsii.String("*"),
		},
		Principals: []IPrincipal{
			iam.NewAnyPrincipal(),
		},
	}),
}))
```

#### ECS-Specific Mixins

**ClusterSettings**: Applies one or more cluster settings to ECS clusters

```go
import ecs "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"


cluster := ecs.NewCfnCluster(scope, jsii.String("Cluster"))
awscdk.Mixins_Of(cluster).Apply(awscdkmixinspreview.NewClusterSettings([]ClusterSettingsProperty{
	&ClusterSettingsProperty{
		Name: jsii.String("containerInsights"),
		Value: jsii.String("enhanced"),
	},
}))
```

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
import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
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
import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
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

### L1 Property Mixins

For every CloudFormation resource, CDK Mixins automatically generates type-safe property mixins. These allow you to apply L1 properties with full TypeScript support:

```go
import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"


s3.NewBucket(scope, jsii.String("Bucket")).With(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	VersioningConfiguration: &VersioningConfigurationProperty{
		Status: jsii.String("Enabled"),
	},
	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
		BlockPublicAcls: jsii.Boolean(true),
		BlockPublicPolicy: jsii.Boolean(true),
	},
}))
```

Property mixins support two merge strategies:

```go
// MERGE (default): Deep merges properties with existing values
awscdk.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	VersioningConfiguration: &VersioningConfigurationProperty{
		Status: jsii.String("Enabled"),
	},
}, &CfnPropertyMixinOptions{
	Strategy: awscdkmixinspreview.PropertyMergeStrategy_MERGE,
}))

// OVERRIDE: Replaces existing property values
awscdk.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	VersioningConfiguration: &VersioningConfigurationProperty{
		Status: jsii.String("Enabled"),
	},
}, &CfnPropertyMixinOptions{
	Strategy: awscdkmixinspreview.PropertyMergeStrategy_OVERRIDE,
}))
```

Property mixins are available for all AWS services:

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
```

### Error Handling

Mixins provide comprehensive error handling:

```go
// Graceful handling of unsupported constructs
awscdk.Mixins_Of(scope).Apply(NewEncryptionAtRest()) // Skips unsupported constructs

// Strict application that requires all constructs to match
awscdk.Mixins_Of(scope).RequireAll().Apply(NewEncryptionAtRest())
```

---


## EventBridge Event Patterns

CDK Mixins automatically generates typed EventBridge event patterns for AWS resources. These patterns work with both L1 and L2 constructs, providing a consistent interface for creating EventBridge rules.

### Event Patterns Basic Usage

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

### Event Pattern Features

**Automatic Resource Injection**: Resource identifiers are automatically included in patterns

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

### Available Events

Event patterns are generated for EventBridge events available in the AWS Event Schema Registry. Common examples:

**S3 Events**:

* `objectCreatedPattern()` - Object creation events
* `objectDeletedPattern()` - Object deletion events
* `objectTagsAddedPattern()` - Object tagging events
* `awsAPICallViaCloudTrailPattern()` - CloudTrail API calls

Import events from service-specific modules:

```go
import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
```

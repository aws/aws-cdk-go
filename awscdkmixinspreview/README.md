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

CDK Mixins provide a new, advanced way to add functionality through composable abstractions.
Unlike traditional L2 constructs that bundle all features together, Mixins allow you to pick and choose exactly the capabilities you need for constructs.

## Key Benefits

* **Universal Compatibility**: Apply the same abstractions to L1 constructs, L2 constructs, or custom constructs
* **Composable Design**: Mix and match features without being locked into specific implementations
* **Cross-Service Abstractions**: Use common patterns like encryption across different AWS services
* **Escape Hatch Freedom**: Customize resources in a safe, typed way while keeping the abstractions you want

## Basic Usage

Mixins use `Mixins.of()` as the fundamental API for applying abstractions to constructs:

```go
// Base form: apply mixins to any construct
bucket := s3.NewCfnBucket(scope, jsii.String("MyBucket"))
awscdkmixinspreview.Mixins_Of(bucket).Apply(NewEncryptionAtRest()).Apply(awscdkmixinspreview.NewAutoDeleteObjects())
```

### Fluent Syntax with `.with()`

For convenience, you can use the `.with()` method for a more fluent syntax:

```go
import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"


bucket := s3.NewCfnBucket(scope, jsii.String("MyBucket")).with(awscdkmixinspreview.NewEnableVersioning()).with(awscdkmixinspreview.NewAutoDeleteObjects())
```

The `.with()` method is available after importing `@aws-cdk/mixins-preview/with`, which augments all constructs with this method. It provides the same functionality as `Mixins.of().apply()` but with a more chainable API.

> **Note**: The `.with()` fluent syntax is only available in JavaScript and TypeScript. Other jsii languages (Python, Java, C#, and Go) should use the `Mixins.of(...).mustApply()` syntax instead. The import requirement is temporary during the preview phase. Once the API is stable, the `.with()` method will be available by default on all constructs and in all languages.

## Creating Custom Mixins

Mixins are simple classes that implement the `IMixin` interface (usually by extending the abstract `Mixin` class:

```go
// Simple mixin that enables versioning
type customVersioningMixin struct {
	Mixin
}

func (this *customVersioningMixin) supports(construct interface{}) *bool {
	return jsii.Boolean(*construct instanceof s3.CfnBucket)
}

func (this *customVersioningMixin) applyTo(bucket interface{}) interface{} {
	*bucket.versioningConfiguration = map[string]*string{
		"status": jsii.String("Enabled"),
	}
	return *bucket
}

// Usage
bucket := s3.NewCfnBucket(scope, jsii.String("MyBucket"))
awscdkmixinspreview.Mixins_Of(bucket).Apply(NewCustomVersioningMixin())
```

## Construct Selection

Mixins operate on construct trees and can be applied selectively:

```go
// Apply to all constructs in a scope
awscdkmixinspreview.Mixins_Of(scope).Apply(NewEncryptionAtRest())

// Apply to specific resource types
awscdkmixinspreview.Mixins_Of(scope, awscdkmixinspreview.ConstructSelector_ResourcesOfType(s3.CfnBucket)).Apply(NewEncryptionAtRest())

// Apply to constructs matching a pattern
awscdkmixinspreview.Mixins_Of(scope, awscdkmixinspreview.ConstructSelector_ById(/.*-prod-.*/)).Apply(NewProductionSecurityMixin())
```

## Built-in Mixins

### Cross-Service Mixins

**EncryptionAtRest**: Applies encryption to supported AWS resources

```go
// Works across different resource types
bucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
awscdkmixinspreview.Mixins_Of(bucket).Apply(NewEncryptionAtRest())

logGroup := logs.NewCfnLogGroup(scope, jsii.String("LogGroup"))
awscdkmixinspreview.Mixins_Of(logGroup).Apply(NewEncryptionAtRest())
```

### S3-Specific Mixins

**AutoDeleteObjects**: Configures automatic object deletion for S3 buckets

```go
bucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewAutoDeleteObjects())
```

**EnableVersioning**: Enables versioning on S3 buckets

```go
bucket := s3.NewCfnBucket(scope, jsii.String("Bucket"))
awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewEnableVersioning())
```

### L1 Property Mixins

For every CloudFormation resource, CDK Mixins automatically generates type-safe property mixins. These allow you to apply L1 properties with full TypeScript support:

```go
import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"


bucket := s3.NewBucket(scope, jsii.String("Bucket")).with(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
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
var bucket CfnBucket


// MERGE (default): Deep merges properties with existing values
awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	VersioningConfiguration: &VersioningConfigurationProperty{
		Status: jsii.String("Enabled"),
	},
}, &CfnPropertyMixinOptions{
	Strategy: awscdkmixinspreview.PropertyMergeStrategy_MERGE,
}))

// OVERWRITE: Replaces existing property values
awscdkmixinspreview.Mixins_Of(bucket).Apply(awscdkmixinspreview.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
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

## Error Handling

Mixins provide comprehensive error handling:

```go
// Graceful handling of unsupported constructs
awscdkmixinspreview.Mixins_Of(scope).Apply(NewEncryptionAtRest()) // Skips unsupported constructs

// Strict application that requires all constructs to match
awscdkmixinspreview.Mixins_Of(scope).MustApply(NewEncryptionAtRest())
```

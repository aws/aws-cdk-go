# CDK CloudFormation Property Mixins

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Stable](https://img.shields.io/badge/cdk--constructs-stable-success.svg?style=for-the-badge)

---
<!--END STABILITY BANNER-->

Auto-generated, type-safe CDK Mixins for every CloudFormation resource property.
These allow you to apply L1 properties to any construct (L1, L2, or custom) using the Mixins mechanism from `aws-cdk-lib`.

## Usage

For every CloudFormation resource, this package provides a `CfnXxxPropsMixin` class.
Apply it using `.with()` or `Mixins.of()`:

```go
s3.NewBucket(scope, jsii.String("MyBucket")).With(awscdkcfnpropertymixins.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	VersioningConfiguration: &VersioningConfigurationProperty{
		Status: jsii.String("Enabled"),
	},
	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
		BlockPublicAcls: jsii.Boolean(true),
		BlockPublicPolicy: jsii.Boolean(true),
	},
}))
```

### Cross-Service References

Deeply nested properties support cross-service references:

```go
myKey := kms.NewKey(scope, jsii.String("MyKey"))

s3.NewBucket(scope, jsii.String("EncryptedBucket")).With(awscdkcfnpropertymixins.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	BucketEncryption: &BucketEncryptionProperty{
		ServerSideEncryptionConfiguration: []interface{}{
			&ServerSideEncryptionRuleProperty{
				ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
					SseAlgorithm: jsii.String("aws:kms"),
					KmsMasterKeyId: myKey,
				},
			},
		},
	},
}))
```

### Merge Strategies

When a mixin is applied, its properties are merged onto the target resource using a merge strategy.
The strategy controls what happens when both the mixin and the existing resource define the same property.

There are two built-in strategies:

#### `PropertyMergeStrategy.combine()` (default)

Deep merges nested objects from the mixin into the target.
When both the existing and new value for a property are plain objects, their keys are merged recursively — existing keys are preserved and new keys are added.
Primitives, arrays, and mismatched types are replaced by the mixin value.

This is useful when you want to add configuration without losing what's already set:

```go
combineBucket := s3.NewCfnBucket(scope, jsii.String("CombineBucket"))
combineBucket.publicAccessBlockConfiguration = &PublicAccessBlockConfigurationProperty{
	BlockPublicAcls: jsii.Boolean(true),
}

// Adds blockPublicPolicy while preserving the existing blockPublicAcls
combineBucket.With(awscdkcfnpropertymixins.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
		BlockPublicPolicy: jsii.Boolean(true),
	},
}))
```

#### `PropertyMergeStrategy.override()`

Replaces existing property values with the mixin values.
Each property is copied as-is without inspecting nested objects.
Any previous value on the target is discarded.

This is useful when you want to fully replace a configuration:

```go
overrideBucket := s3.NewCfnBucket(scope, jsii.String("OverrideBucket"))
overrideBucket.publicAccessBlockConfiguration = &PublicAccessBlockConfigurationProperty{
	BlockPublicAcls: jsii.Boolean(true),
}

// Replaces the entire publicAccessBlockConfiguration
overrideBucket.With(awscdkcfnpropertymixins.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
		BlockPublicPolicy: jsii.Boolean(true),
	},
}, &CfnPropertyMixinOptions{
	Strategy: awscdk.PropertyMergeStrategy_Override(),
}))
```

#### Custom Strategies

You can implement `IMergeStrategy` to define your own merge behavior.
The `apply` method receives the target object, source object, and an allowlist of property keys:

```go
type arrayAppendStrategy struct {
}

func (this *arrayAppendStrategy) apply(target interface{}, source interface{}, allowedKeys []*string) {
	for _, key := range *allowedKeys {
		if key in *source {
			if array.isArray(*target[key]) {
				// append to target
				*target[key] = *target[key].concat(*source[key])
			} else {
				// override
				*target[key] = *source[key]
			}
		}
	}
}
```

### Deferred Values (Boxes)

Property mixins support **Box-backed values**. Most L2 constructs in `aws-cdk-lib` use Boxes internally to defer property computation until synthesis time. When a mixin encounters a Box on the target, the merge is automatically deferred — the merge strategy runs once the Box resolves, ensuring it operates on final values.

This means mixins work correctly with L2 constructs that use Boxes for properties like `replicas`, `rules`, or `tags`, without any special handling from the user:

```go
// TableV2 uses a Box internally for replicas.
// The mixin defers the merge and appends the new replica at synthesis time.
table := dynamodb.NewTableV2(scope, jsii.String("Table"), &TablePropsV2{
	PartitionKey: &Attribute{
		Name: jsii.String("pk"),
		Type: dynamodb.AttributeType_STRING,
	},
	// L2 prop: pointInTimeRecovery is a boolean
	Replicas: []ReplicaTableProps{
		&ReplicaTableProps{
			Region: jsii.String("us-east-1"),
			PointInTimeRecovery: jsii.Boolean(true),
		},
	},
})

// Mixins always use L1 (CloudFormation) property names and shapes,
// regardless of what the L2 API looks like.
table.With(awscdkcfnpropertymixins.NewCfnGlobalTablePropsMixin(&CfnGlobalTableMixinProps{
	Replicas: []interface{}{
		&ReplicaSpecificationProperty{
			Region: jsii.String("eu-west-1"),
			// L1 prop: pointInTimeRecoverySpecification is an object
			PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
				PointInTimeRecoveryEnabled: jsii.Boolean(true),
			},
		},
	},
}, &CfnPropertyMixinOptions{
	Strategy: awscdk.PropertyMergeStrategy_Combine(&CombineStrategyOptions{
		Arrays: awscdk.ArrayMergeStrategy_Append(),
	}),
}))
```

Most L2 constructs in `aws-cdk-lib` use Boxes or Lazys internally to defer property computation until synthesis time. Property mixins detect these automatically and defer the merge until the value resolves, so the merge strategy always operates on final values — no special handling is needed from the user. The only case where merging cannot be deferred is a raw `Token` that is not backed by a Box. This is very rare in the AWS Construct Library, but may occur in third-party packages. If you encounter a construct where merging doesn't work as expected, please [open an issue](https://github.com/aws/aws-cdk/issues/new?template=bug-report.yml) so we can investigate.

### CloudFormation Property Mixins for Every Service

This package provides auto-generated property mixins for every CloudFormation resource across all AWS services.
Each service has its own submodule that mirrors the `aws-cdk-lib` module structure.
Import the mixin for the resource you want to configure from the corresponding service submodule:

```go
import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
```

The naming convention follows a consistent pattern: for a CloudFormation resource `AWS::S3::Bucket`, the mixin class is `CfnBucketPropsMixin` and lives in the `aws-s3` submodule.
The mixin props interface is named `CfnBucketMixinProps` and all properties are optional, so you only need to specify the ones you want to set.

Property mixins work with any construct that has the target resource as its default child.
This means you can apply them to L1 constructs, L2 constructs, and custom constructs alike:

```go
// L1 construct
// L1 construct
s3.NewCfnBucket(scope, jsii.String("L1Bucket")).With(awscdkcfnpropertymixins.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	VersioningConfiguration: &VersioningConfigurationProperty{
		Status: jsii.String("Enabled"),
	},
}))

// L2 construct — the mixin finds the CfnBucket default child
// L2 construct — the mixin finds the CfnBucket default child
s3.NewBucket(scope, jsii.String("L2Bucket")).With(awscdkcfnpropertymixins.NewCfnBucketPropsMixin(&CfnBucketMixinProps{
	VersioningConfiguration: &VersioningConfigurationProperty{
		Status: jsii.String("Enabled"),
	},
}))
```

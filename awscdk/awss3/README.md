# Amazon S3 Construct Library

Define an unencrypted S3 bucket.

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyFirstBucket"))
```

`Bucket` constructs expose the following deploy-time attributes:

* `bucketArn` - the ARN of the bucket (i.e. `arn:aws:s3:::bucket_name`)
* `bucketName` - the name of the bucket (i.e. `bucket_name`)
* `bucketWebsiteUrl` - the Website URL of the bucket (i.e.
  `http://bucket_name.s3-website-us-west-1.amazonaws.com`)
* `bucketDomainName` - the URL of the bucket (i.e. `bucket_name.s3.amazonaws.com`)
* `bucketDualStackDomainName` - the dual-stack URL of the bucket (i.e.
  `bucket_name.s3.dualstack.eu-west-1.amazonaws.com`)
* `bucketRegionalDomainName` - the regional URL of the bucket (i.e.
  `bucket_name.s3.eu-west-1.amazonaws.com`)
* `arnForObjects(pattern)` - the ARN of an object or objects within the bucket (i.e.
  `arn:aws:s3:::bucket_name/exampleobject.png` or
  `arn:aws:s3:::bucket_name/Development/*`)
* `urlForObject(key)` - the HTTP URL of an object within the bucket (i.e.
  `https://s3.cn-north-1.amazonaws.com.cn/china-bucket/mykey`)
* `virtualHostedUrlForObject(key)` - the virtual-hosted style HTTP URL of an object
  within the bucket (i.e. `https://china-bucket-s3.cn-north-1.amazonaws.com.cn/mykey`)
* `s3UrlForObject(key)` - the S3 URL of an object within the bucket (i.e.
  `s3://bucket/mykey`)

## Encryption

Define a KMS-encrypted bucket:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &bucketProps{
	encryption: s3.bucketEncryption_KMS,
})

// you can access the encryption key:
assert(bucket.encryptionKey instanceof kms.key)
```

You can also supply your own key:

```go
// Example automatically generated from non-compiling source. May contain errors.
myKmsKey := kms.NewKey(this, jsii.String("MyKey"))

bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &bucketProps{
	encryption: s3.bucketEncryption_KMS,
	encryptionKey: myKmsKey,
})

assert(bucket.encryptionKey == myKmsKey)
```

Enable KMS-SSE encryption via [S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html):

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &bucketProps{
	encryption: s3.bucketEncryption_KMS,
	bucketKeyEnabled: jsii.Boolean(true),
})
```

Use `BucketEncryption.ManagedKms` to use the S3 master KMS key:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("Buck"), &bucketProps{
	encryption: s3.bucketEncryption_KMS_MANAGED,
})

assert(bucket.encryptionKey == nil)
```

## Permissions

A bucket policy will be automatically created for the bucket upon the first call to
`addToResourcePolicy(statement)`:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
result := bucket.addToResourcePolicy(iam.NewPolicyStatement(&policyStatementProps{
	actions: []*string{
		jsii.String("s3:GetObject"),
	},
	resources: []*string{
		bucket.arnForObjects(jsii.String("file.txt")),
	},
	principals: []iPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))
```

If you try to add a policy statement to an existing bucket, this method will
not do anything:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.bucket.fromBucketName(this, jsii.String("existingBucket"), jsii.String("bucket-name"))

// No policy statement will be added to the resource
result := bucket.addToResourcePolicy(iam.NewPolicyStatement(&policyStatementProps{
	actions: []*string{
		jsii.String("s3:GetObject"),
	},
	resources: []*string{
		bucket.arnForObjects(jsii.String("file.txt")),
	},
	principals: []iPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))
```

That's because it's not possible to tell whether the bucket
already has a policy attached, let alone to re-use that policy to add more
statements to it. We recommend that you always check the result of the call:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
result := bucket.addToResourcePolicy(iam.NewPolicyStatement(&policyStatementProps{
	actions: []*string{
		jsii.String("s3:GetObject"),
	},
	resources: []*string{
		bucket.arnForObjects(jsii.String("file.txt")),
	},
	principals: []iPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))

if !result.statementAdded {}
```

The bucket policy can be directly accessed after creation to add statements or
adjust the removal policy.

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.policy.applyRemovalPolicy(cdk.removalPolicy_RETAIN)
```

Most of the time, you won't have to manipulate the bucket policy directly.
Instead, buckets have "grant" methods called to give prepackaged sets of permissions
to other resources. For example:

```go
// Example automatically generated from non-compiling source. May contain errors.
var myLambda function


bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.grantReadWrite(myLambda)
```

Will give the Lambda's execution role permissions to read and write
from the bucket.

## AWS Foundational Security Best Practices

### Enforcing SSL

To require all requests use Secure Socket Layer (SSL):

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("Bucket"), &bucketProps{
	enforceSSL: jsii.Boolean(true),
})
```

## Sharing buckets between stacks

To use a bucket in a different stack in the same CDK application, pass the object to the other stack:

```go
/**
 * Stack that defines the bucket
 */
type producer struct {
	stack
	myBucket bucket
}

func newProducer(scope app, id *string, props stackProps) *producer {
	this := &producer{}
	cdk.NewStack_Override(this, scope, id, props)

	bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
		removalPolicy: cdk.removalPolicy_DESTROY,
	})
	this.myBucket = bucket
	return this
}

type consumerProps struct {
	stackProps
	userBucket iBucket
}

/**
 * Stack that consumes the bucket
 */
type consumer struct {
	stack
}

func newConsumer(scope app, id *string, props consumerProps) *consumer {
	this := &consumer{}
	cdk.NewStack_Override(this, scope, id, props)

	user := iam.NewUser(this, jsii.String("MyUser"))
	*props.userBucket.grantReadWrite(user)
	return this
}

producer := NewProducer(app, jsii.String("ProducerStack"))
NewConsumer(app, jsii.String("ConsumerStack"), &consumerProps{
	userBucket: producer.myBucket,
})
```

## Importing existing buckets

To import an existing bucket into your CDK application, use the `Bucket.fromBucketAttributes`
factory method. This method accepts `BucketAttributes` which describes the properties of an already
existing bucket:

```go
// Example automatically generated from non-compiling source. May contain errors.
var myLambda function

bucket := s3.bucket.fromBucketAttributes(this, jsii.String("ImportedBucket"), &bucketAttributes{
	bucketArn: jsii.String("arn:aws:s3:::my-bucket"),
})

// now you can just call methods on the bucket
bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &notificationKeyFilter{
	prefix: jsii.String("home/myusername/*"),
})
```

Alternatively, short-hand factories are available as `Bucket.fromBucketName` and
`Bucket.fromBucketArn`, which will derive all bucket attributes from the bucket
name or ARN respectively:

```go
// Example automatically generated from non-compiling source. May contain errors.
byName := s3.bucket.fromBucketName(this, jsii.String("BucketByName"), jsii.String("my-bucket"))
byArn := s3.bucket.fromBucketArn(this, jsii.String("BucketByArn"), jsii.String("arn:aws:s3:::my-bucket"))
```

The bucket's region defaults to the current stack's region, but can also be explicitly set in cases where one of the bucket's
regional properties needs to contain the correct values.

```go
// Example automatically generated from non-compiling source. May contain errors.
myCrossRegionBucket := s3.bucket.fromBucketAttributes(this, jsii.String("CrossRegionImport"), &bucketAttributes{
	bucketArn: jsii.String("arn:aws:s3:::my-bucket"),
	region: jsii.String("us-east-1"),
})
```

## Bucket Notifications

The Amazon S3 notification feature enables you to receive notifications when
certain events happen in your bucket as described under [S3 Bucket
Notifications] of the S3 Developer Guide.

To subscribe for bucket notifications, use the `bucket.addEventNotification` method. The
`bucket.addObjectCreatedNotification` and `bucket.addObjectRemovedNotification` can also be used for
these common use cases.

The following example will subscribe an SNS topic to be notified of all `s3:ObjectCreated:*` events:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
topic := sns.NewTopic(this, jsii.String("MyTopic"))
bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewSnsDestination(topic))
```

This call will also ensure that the topic policy can accept notifications for
this specific bucket.

Supported S3 notification targets are exposed by the `@aws-cdk/aws-s3-notifications` package.

It is also possible to specify S3 object key filters when subscribing. The
following example will notify `myQueue` when objects prefixed with `foo/` and
have the `.jpg` suffix are removed from the bucket.

```go
// Example automatically generated from non-compiling source. May contain errors.
var myQueue queue

bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.addEventNotification(s3.eventType_OBJECT_REMOVED,
s3n.NewSqsDestination(myQueue), &notificationKeyFilter{
	prefix: jsii.String("foo/"),
	suffix: jsii.String(".jpg"),
})
```

Adding notifications on existing buckets:

```go
// Example automatically generated from non-compiling source. May contain errors.
var topic topic

bucket := s3.bucket.fromBucketAttributes(this, jsii.String("ImportedBucket"), &bucketAttributes{
	bucketArn: jsii.String("arn:aws:s3:::my-bucket"),
})
bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewSnsDestination(topic))
```

When you add an event notification to a bucket, a custom resource is created to
manage the notifications. By default, a new role is created for the Lambda
function that implements this feature. If you want to use your own role instead,
you should provide it in the `Bucket` constructor:

```go
// Example automatically generated from non-compiling source. May contain errors.
var myRole iRole

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	notificationsHandlerRole: myRole,
})
```

Whatever role you provide, the CDK will try to modify it by adding the
permissions from `AWSLambdaBasicExecutionRole` (an AWS managed policy) as well
as the permissions `s3:PutBucketNotification` and `s3:GetBucketNotification`.
If you’re passing an imported role, and you don’t want this to happen, configure
it to be immutable:

```go
// Example automatically generated from non-compiling source. May contain errors.
importedRole := iam.role.fromRoleArn(this, jsii.String("role"), jsii.String("arn:aws:iam::123456789012:role/RoleName"), &fromRoleArnOptions{
	mutable: jsii.Boolean(false),
})
```

> If you provide an imported immutable role, make sure that it has at least all
> the permissions mentioned above. Otherwise, the deployment will fail!

### EventBridge notifications

Amazon S3 can send events to Amazon EventBridge whenever certain events happen in your bucket.
Unlike other destinations, you don't need to select which event types you want to deliver.

The following example will enable EventBridge notifications:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyEventBridgeBucket"), &bucketProps{
	eventBridgeEnabled: jsii.Boolean(true),
})
```

## Block Public Access

Use `blockPublicAccess` to specify [block public access settings](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html) on the bucket.

Enable all block public access settings:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &bucketProps{
	blockPublicAccess: s3.blockPublicAccess_BLOCK_ALL(),
})
```

Block and ignore public ACLs:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &bucketProps{
	blockPublicAccess: s3.blockPublicAccess_BLOCK_ACLS(),
})
```

Alternatively, specify the settings manually:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &bucketProps{
	blockPublicAccess: s3.NewBlockPublicAccess(&blockPublicAccessOptions{
		blockPublicPolicy: jsii.Boolean(true),
	}),
})
```

When `blockPublicPolicy` is set to `true`, `grantPublicRead()` throws an error.

## Logging configuration

Use `serverAccessLogsBucket` to describe where server access logs are to be stored.

```go
// Example automatically generated from non-compiling source. May contain errors.
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	serverAccessLogsBucket: accessLogsBucket,
})
```

It's also possible to specify a prefix for Amazon S3 to assign to all log object keys.

```go
// Example automatically generated from non-compiling source. May contain errors.
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	serverAccessLogsBucket: accessLogsBucket,
	serverAccessLogsPrefix: jsii.String("logs"),
})
```

## S3 Inventory

An [inventory](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) contains a list of the objects in the source bucket and metadata for each object. The inventory lists are stored in the destination bucket as a CSV file compressed with GZIP, as an Apache optimized row columnar (ORC) file compressed with ZLIB, or as an Apache Parquet (Parquet) file compressed with Snappy.

You can configure multiple inventory lists for a bucket. You can configure what object metadata to include in the inventory, whether to list all object versions or only current versions, where to store the inventory list file output, and whether to generate the inventory on a daily or weekly basis.

```go
// Example automatically generated from non-compiling source. May contain errors.
inventoryBucket := s3.NewBucket(this, jsii.String("InventoryBucket"))

dataBucket := s3.NewBucket(this, jsii.String("DataBucket"), &bucketProps{
	inventories: []inventory{
		&inventory{
			frequency: s3.inventoryFrequency_DAILY,
			includeObjectVersions: s3.inventoryObjectVersion_CURRENT,
			destination: &inventoryDestination{
				bucket: inventoryBucket,
			},
		},
		&inventory{
			frequency: s3.*inventoryFrequency_WEEKLY,
			includeObjectVersions: s3.*inventoryObjectVersion_ALL,
			destination: &inventoryDestination{
				bucket: inventoryBucket,
				prefix: jsii.String("with-all-versions"),
			},
		},
	},
})
```

If the destination bucket is created as part of the same CDK application, the necessary permissions will be automatically added to the bucket policy.
However, if you use an imported bucket (i.e `Bucket.fromXXX()`), you'll have to make sure it contains the following policy document:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "InventoryAndAnalyticsExamplePolicy",
      "Effect": "Allow",
      "Principal": { "Service": "s3.amazonaws.com" },
      "Action": "s3:PutObject",
      "Resource": ["arn:aws:s3:::destinationBucket/*"]
    }
  ]
}
```

## Website redirection

You can use the two following properties to specify the bucket [redirection policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html#advanced-conditional-redirects). Please note that these methods cannot both be applied to the same bucket.

### Static redirection

You can statically redirect a to a given Bucket URL or any other host name with `websiteRedirect`:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
	websiteRedirect: &redirectTarget{
		hostName: jsii.String("www.example.com"),
	},
})
```

### Routing rules

Alternatively, you can also define multiple `websiteRoutingRules`, to define complex, conditional redirections:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
	websiteRoutingRules: []routingRule{
		&routingRule{
			hostName: jsii.String("www.example.com"),
			httpRedirectCode: jsii.String("302"),
			protocol: s3.redirectProtocol_HTTPS,
			replaceKey: s3.replaceKey.prefixWith(jsii.String("test/")),
			condition: &routingRuleCondition{
				httpErrorCodeReturnedEquals: jsii.String("200"),
				keyPrefixEquals: jsii.String("prefix"),
			},
		},
	},
})
```

## Filling the bucket as part of deployment

To put files into a bucket as part of a deployment (for example, to host a
website), see the `@aws-cdk/aws-s3-deployment` package, which provides a
resource that can do just that.

## The URL for objects

S3 provides two types of URLs for accessing objects via HTTP(S). Path-Style and
[Virtual Hosted-Style](https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html)
URL. Path-Style is a classic way and will be
[deprecated](https://aws.amazon.com/jp/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story).
We recommend to use Virtual Hosted-Style URL for newly made bucket.

You can generate both of them.

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.urlForObject(jsii.String("objectname")) // Path-Style URL
bucket.virtualHostedUrlForObject(jsii.String("objectname")) // Virtual Hosted-Style URL
bucket.virtualHostedUrlForObject(jsii.String("objectname"), &virtualHostedStyleUrlOptions{
	regional: jsii.Boolean(false),
})
```

## Object Ownership

You can use one of following properties to specify the bucket [object Ownership](https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html).

### Object writer

The Uploading account will own the object.

```go
// Example automatically generated from non-compiling source. May contain errors.
s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	objectOwnership: s3.objectOwnership_OBJECT_WRITER,
})
```

### Bucket owner preferred

The bucket owner will own the object if the object is uploaded with the bucket-owner-full-control canned ACL. Without this setting and canned ACL, the object is uploaded and remains owned by the uploading account.

```go
// Example automatically generated from non-compiling source. May contain errors.
s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	objectOwnership: s3.objectOwnership_BUCKET_OWNER_PREFERRED,
})
```

### Bucket owner enforced (recommended)

ACLs are disabled, and the bucket owner automatically owns and has full control over every object in the bucket. ACLs no longer affect permissions to data in the S3 bucket. The bucket uses policies to define access control.

```go
// Example automatically generated from non-compiling source. May contain errors.
s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	objectOwnership: s3.objectOwnership_BUCKET_OWNER_ENFORCED,
})
```

## Bucket deletion

When a bucket is removed from a stack (or the stack is deleted), the S3
bucket will be removed according to its removal policy (which by default will
simply orphan the bucket and leave it in your AWS account). If the removal
policy is set to `RemovalPolicy.DESTROY`, the bucket will be deleted as long
as it does not contain any objects.

To override this and force all objects to get deleted during bucket deletion,
enable the`autoDeleteObjects` option.

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyTempFileBucket"), &bucketProps{
	removalPolicy: cdk.removalPolicy_DESTROY,
	autoDeleteObjects: jsii.Boolean(true),
})
```

**Warning** if you have deployed a bucket with `autoDeleteObjects: true`,
switching this to `false` in a CDK version *before* `1.126.0` will lead to
all objects in the bucket being deleted. Be sure to update your bucket resources
by deploying with CDK version `1.126.0` or later **before** switching this value to `false`.

## Transfer Acceleration

[Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/transfer-acceleration.html) can be configured to enable fast, easy, and secure transfers of files over long distances:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	transferAcceleration: jsii.Boolean(true),
})
```

To access the bucket that is enabled for Transfer Acceleration, you must use a special endpoint. The URL can be generated using method `transferAccelerationUrlForObject`:

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	transferAcceleration: jsii.Boolean(true),
})
bucket.transferAccelerationUrlForObject(jsii.String("objectname"))
```

## Intelligent Tiering

[Intelligent Tiering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html) can be configured to automatically move files to glacier:

```go
// Example automatically generated from non-compiling source. May contain errors.
s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	intelligentTieringConfigurations: []intelligentTieringConfiguration{
		&intelligentTieringConfiguration{
			name: jsii.String("foo"),
			prefix: jsii.String("folder/name"),
			archiveAccessTierTime: cdk.duration.days(jsii.Number(90)),
			deepArchiveAccessTierTime: cdk.*duration.days(jsii.Number(180)),
			tags: []tag{
				&tag{
					key: jsii.String("tagname"),
					value: jsii.String("tagvalue"),
				},
			},
		},
	},
})
```

## Lifecycle Rule

[Managing lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html) can be configured transition or expiration actions.

```go
// Example automatically generated from non-compiling source. May contain errors.
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
	lifecycleRules: []lifecycleRule{
		&lifecycleRule{
			abortIncompleteMultipartUploadAfter: cdk.duration.minutes(jsii.Number(30)),
			enabled: jsii.Boolean(false),
			expiration: cdk.*duration.days(jsii.Number(30)),
			expirationDate: NewDate(),
			expiredObjectDeleteMarker: jsii.Boolean(false),
			id: jsii.String("id"),
			noncurrentVersionExpiration: cdk.*duration.days(jsii.Number(30)),

			// the properties below are optional
			noncurrentVersionsToRetain: jsii.Number(123),
			noncurrentVersionTransitions: []noncurrentVersionTransition{
				&noncurrentVersionTransition{
					storageClass: s3.storageClass_GLACIER(),
					transitionAfter: cdk.*duration.days(jsii.Number(30)),

					// the properties below are optional
					noncurrentVersionsToRetain: jsii.Number(123),
				},
			},
			objectSizeGreaterThan: jsii.Number(500),
			prefix: jsii.String("prefix"),
			objectSizeLessThan: jsii.Number(10000),
			transitions: []transition{
				&transition{
					storageClass: s3.*storageClass_GLACIER(),

					// the properties below are optional
					transitionAfter: cdk.*duration.days(jsii.Number(30)),
					transitionDate: NewDate(),
				},
			},
		},
	},
})
```

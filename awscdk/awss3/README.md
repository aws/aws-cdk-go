# Amazon S3 Construct Library

Define an S3 bucket.

```go
bucket := s3.NewBucket(this, jsii.String("MyFirstBucket"))
```

`Bucket` constructs expose the following deploy-time attributes:

* `bucketArn` - the ARN of the bucket (i.e. `arn:aws:s3:::amzn-s3-demo-bucket`)
* `bucketName` - the name of the bucket (i.e. `amzn-s3-demo-bucket`)
* `bucketWebsiteUrl` - the Website URL of the bucket (i.e.
  `http://amzn-s3-demo-bucket.s3-website-us-west-1.amazonaws.com`)
* `bucketDomainName` - the URL of the bucket (i.e. `amzn-s3-demo-bucket.s3.amazonaws.com`)
* `bucketDualStackDomainName` - the dual-stack URL of the bucket (i.e.
  `amzn-s3-demo-bucket.s3.dualstack.eu-west-1.amazonaws.com`)
* `bucketRegionalDomainName` - the regional URL of the bucket (i.e.
  `amzn-s3-demo-bucket.s3.eu-west-1.amazonaws.com`)
* `arnForObjects(pattern)` - the ARN of an object or objects within the bucket (i.e.
  `arn:aws:s3:::amzn-s3-demo-bucket/exampleobject.png` or
  `arn:aws:s3:::amzn-s3-demo-bucket/Development/*`)
* `urlForObject(key)` - the HTTP URL of an object within the bucket (i.e.
  `https://s3.cn-north-1.amazonaws.com.cn/china-bucket/mykey`)
* `virtualHostedUrlForObject(key)` - the virtual-hosted style HTTP URL of an object
  within the bucket (i.e. `https://china-bucket-s3.cn-north-1.amazonaws.com.cn/mykey`)
* `s3UrlForObject(key)` - the S3 URL of an object within the bucket (i.e.
  `s3://bucket/mykey`)

## Encryption

Define a KMS-encrypted bucket:

```go
bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &BucketProps{
	Encryption: s3.BucketEncryption_KMS,
})

// you can access the encryption key:
assert(bucket.EncryptionKey instanceof kms.Key)
```

You can also supply your own key:

```go
myKmsKey := kms.NewKey(this, jsii.String("MyKey"))

bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &BucketProps{
	Encryption: s3.BucketEncryption_KMS,
	EncryptionKey: myKmsKey,
})

assert(bucket.EncryptionKey == myKmsKey)
```

Enable KMS-SSE encryption via [S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html):

```go
bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &BucketProps{
	Encryption: s3.BucketEncryption_KMS,
	BucketKeyEnabled: jsii.Boolean(true),
})
```

Use `BucketEncryption.ManagedKms` to use the S3 master KMS key:

```go
bucket := s3.NewBucket(this, jsii.String("Buck"), &BucketProps{
	Encryption: s3.BucketEncryption_KMS_MANAGED,
})

assert(bucket.EncryptionKey == nil)
```

Enable DSSE encryption:

```
const bucket = new s3.Bucket(stack, 'MyDSSEBucket', {
  encryption: s3.BucketEncryption.DSSE_MANAGED,
  bucketKeyEnabled: true,
});
```

## Permissions

A bucket policy will be automatically created for the bucket upon the first call to
`addToResourcePolicy(statement)`:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
result := bucket.AddToResourcePolicy(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("s3:GetObject"),
	},
	Resources: []*string{
		bucket.ArnForObjects(jsii.String("file.txt")),
	},
	Principals: []iPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))
```

If you try to add a policy statement to an existing bucket, this method will
not do anything:

```go
bucket := s3.Bucket_FromBucketName(this, jsii.String("existingBucket"), jsii.String("amzn-s3-demo-bucket"))

// No policy statement will be added to the resource
result := bucket.AddToResourcePolicy(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("s3:GetObject"),
	},
	Resources: []*string{
		bucket.ArnForObjects(jsii.String("file.txt")),
	},
	Principals: []iPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))
```

That's because it's not possible to tell whether the bucket
already has a policy attached, let alone to re-use that policy to add more
statements to it. We recommend that you always check the result of the call:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
result := bucket.AddToResourcePolicy(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("s3:GetObject"),
	},
	Resources: []*string{
		bucket.ArnForObjects(jsii.String("file.txt")),
	},
	Principals: []iPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))

if !result.StatementAdded {}
```

The bucket policy can be directly accessed after creation to add statements or
adjust the removal policy.

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.Policy.ApplyRemovalPolicy(cdk.RemovalPolicy_RETAIN)
```

Most of the time, you won't have to manipulate the bucket policy directly.
Instead, buckets have "grant" methods called to give prepackaged sets of permissions
to other resources. For example:

```go
var myLambda function


bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.GrantReadWrite(myLambda)
```

Will give the Lambda's execution role permissions to read and write
from the bucket.

### Understanding "grant" Methods

The S3 construct library provides several grant methods for the `Bucket` resource, but two of them have a special behavior. This two accept an `objectsKeyPattern` parameter to restrict granted permissions to specific resources:

* `grantRead`
* `grantReadWrite`

When examining the synthesized policy, you'll notice it includes both your specified object key patterns and the bucket itself.
This is by design. Some permissions (like `s3:ListBucket`) apply at the bucket level, while others (like `s3:GetObject`) apply to specific objects.

Specifically, the [`s3:ListBucket` action operates on bucket resources](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html#amazons3-bucket)
and requires the bucket ARN to work properly. This might be seen as a bug, giving the impression that more permissions were granted than the ones you intended, but the reality is that the policy does not ignore your `objectsKeyPattern` - object-specific actions like `s3:GetObject`
will still be limited to the resources defined in your pattern.

If you need to restrict the `s3:ListBucket` action to specific paths, you can add a `Condition` to your policy that limits the `objectsKeyPattern` to specific folders. For more details and examples, see the [AWS documentation on bucket policies](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html#example-bucket-policies-folders).

## AWS Foundational Security Best Practices

### Enforcing SSL

To require all requests use Secure Socket Layer (SSL):

```go
bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
	EnforceSSL: jsii.Boolean(true),
})
```

To require a minimum TLS version for all requests:

```go
bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
	EnforceSSL: jsii.Boolean(true),
	MinimumTLSVersion: jsii.Number(1.2),
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

func newProducer(scope construct, id *string, props stackProps) *producer {
	this := &producer{}
	newStack_Override(this, scope, id, props)

	bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
		RemovalPolicy: cdk.RemovalPolicy_DESTROY,
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

func newConsumer(scope construct, id *string, props consumerProps) *consumer {
	this := &consumer{}
	newStack_Override(this, scope, id, props)

	user := iam.NewUser(this, jsii.String("MyUser"))
	*props.userBucket.GrantReadWrite(user)
	return this
}

app := awscdk.NewApp()
producer := NewProducer(app, jsii.String("ProducerStack"))
NewConsumer(app, jsii.String("ConsumerStack"), &consumerProps{
	userBucket: producer.myBucket,
})
```

## Importing existing buckets

To import an existing bucket into your CDK application, use the `Bucket.fromBucketAttributes`
factory method. This method accepts `BucketAttributes` which describes the properties of an already
existing bucket:

Note that this method allows importing buckets with legacy names containing uppercase letters (`A-Z`) or underscores (`_`), which were
permitted for buckets created before March 1, 2018. For buckets created after this date, uppercase letters and underscores
are not allowed in the bucket name.

```go
var myLambda function

bucket := s3.Bucket_FromBucketAttributes(this, jsii.String("ImportedBucket"), &BucketAttributes{
	BucketArn: jsii.String("arn:aws:s3:::amzn-s3-demo-bucket"),
})

// now you can just call methods on the bucket
bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &NotificationKeyFilter{
	Prefix: jsii.String("home/myusername/*"),
})
```

Alternatively, short-hand factories are available as `Bucket.fromBucketName` and
`Bucket.fromBucketArn`, which will derive all bucket attributes from the bucket
name or ARN respectively:

```go
byName := s3.Bucket_FromBucketName(this, jsii.String("BucketByName"), jsii.String("amzn-s3-demo-bucket"))
byArn := s3.Bucket_FromBucketArn(this, jsii.String("BucketByArn"), jsii.String("arn:aws:s3:::amzn-s3-demo-bucket"))
```

The bucket's region defaults to the current stack's region, but can also be explicitly set in cases where one of the bucket's
regional properties needs to contain the correct values.

```go
myCrossRegionBucket := s3.Bucket_FromBucketAttributes(this, jsii.String("CrossRegionImport"), &BucketAttributes{
	BucketArn: jsii.String("arn:aws:s3:::amzn-s3-demo-bucket"),
	Region: jsii.String("us-east-1"),
})
```

## Bucket Notifications

The Amazon S3 notification feature enables you to receive notifications when
certain events happen in your bucket as described under [S3 Bucket
Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) of the S3 Developer Guide.

To subscribe for bucket notifications, use the `bucket.addEventNotification` method. The
`bucket.addObjectCreatedNotification` and `bucket.addObjectRemovedNotification` can also be used for
these common use cases.

The following example will subscribe an SNS topic to be notified of all `s3:ObjectCreated:*` events:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
topic := sns.NewTopic(this, jsii.String("MyTopic"))
bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewSnsDestination(topic))
```

This call will also ensure that the topic policy can accept notifications for
this specific bucket.

Supported S3 notification targets are exposed by the `aws-cdk-lib/aws-s3-notifications` package.

It is also possible to specify S3 object key filters when subscribing. The
following example will notify `myQueue` when objects prefixed with `foo/` and
have the `.jpg` suffix are removed from the bucket.

```go
var myQueue queue

bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.AddEventNotification(s3.EventType_OBJECT_REMOVED, s3n.NewSqsDestination(myQueue), &NotificationKeyFilter{
	Prefix: jsii.String("foo/"),
	Suffix: jsii.String(".jpg"),
})
```

Adding notifications on existing buckets:

```go
var topic topic

bucket := s3.Bucket_FromBucketAttributes(this, jsii.String("ImportedBucket"), &BucketAttributes{
	BucketArn: jsii.String("arn:aws:s3:::amzn-s3-demo-bucket"),
})
bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewSnsDestination(topic))
```

If you do not want for S3 to validate permissions of Amazon SQS, Amazon SNS, and Lambda destinations you can use the `notificationsSkipDestinationValidation` flag:

```go
var myQueue queue

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	NotificationsSkipDestinationValidation: jsii.Boolean(true),
})
bucket.AddEventNotification(s3.EventType_OBJECT_REMOVED, s3n.NewSqsDestination(myQueue))
```

When you add an event notification to a bucket, a custom resource is created to
manage the notifications. By default, a new role is created for the Lambda
function that implements this feature. If you want to use your own role instead,
you should provide it in the `Bucket` constructor:

```go
var myRole iRole

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	NotificationsHandlerRole: myRole,
})
```

Whatever role you provide, the CDK will try to modify it by adding the
permissions from `AWSLambdaBasicExecutionRole` (an AWS managed policy) as well
as the permissions `s3:PutBucketNotification` and `s3:GetBucketNotification`.
If you’re passing an imported role, and you don’t want this to happen, configure
it to be immutable:

```go
importedRole := iam.Role_FromRoleArn(this, jsii.String("role"), jsii.String("arn:aws:iam::123456789012:role/RoleName"), &FromRoleArnOptions{
	Mutable: jsii.Boolean(false),
})
```

> If you provide an imported immutable role, make sure that it has at least all
> the permissions mentioned above. Otherwise, the deployment will fail!

### EventBridge notifications

Amazon S3 can send events to Amazon EventBridge whenever certain events happen in your bucket.
Unlike other destinations, you don't need to select which event types you want to deliver.

The following example will enable EventBridge notifications:

```go
bucket := s3.NewBucket(this, jsii.String("MyEventBridgeBucket"), &BucketProps{
	EventBridgeEnabled: jsii.Boolean(true),
})
```

## Block Public Access

Use `blockPublicAccess` to specify [block public access settings](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html) on the bucket.

Enable all block public access settings:

```go
bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &BucketProps{
	BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
})
```

Block and ignore public ACLs (other options remain unblocked):

```go
bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &BucketProps{
	BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ACLS_ONLY(),
})
```

Alternatively, specify the settings manually (unspecified options will remain blocked):

```go
bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &BucketProps{
	BlockPublicAccess: s3.NewBlockPublicAccess(&BlockPublicAccessOptions{
		BlockPublicPolicy: jsii.Boolean(false),
	}),
})
```

When `blockPublicPolicy` is set to `true`, `grantPublicRead()` throws an error.

## Public Read Access

Use `publicReadAccess` to allow public read access to the bucket.

Note that to enable `publicReadAccess`, make sure both bucket-level and account-level block public access control is disabled.

Bucket-level block public access control can be configured through `blockPublicAccess` property. Account-level block public
access control can be configured on AWS Console -> S3 -> Block Public Access settings for this account (Navigation Panel).

```go
bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
	PublicReadAccess: jsii.Boolean(true),
	BlockPublicAccess: map[string]*bool{
		"blockPublicPolicy": jsii.Boolean(false),
		"blockPublicAcls": jsii.Boolean(false),
		"ignorePublicAcls": jsii.Boolean(false),
		"restrictPublicBuckets": jsii.Boolean(false),
	},
})
```

## Logging configuration

Use `serverAccessLogsBucket` to describe where server access logs are to be stored.

```go
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
})
```

It's also possible to specify a prefix for Amazon S3 to assign to all log object keys.

```go
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
	ServerAccessLogsPrefix: jsii.String("logs"),
})
```

You have two options for the log object key format.
`Non-date-based partitioning` is the default log object key format and appears as follows:

```txt
[DestinationPrefix][YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]
```

```go
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
	ServerAccessLogsPrefix: jsii.String("logs"),
	// You can use a simple prefix with `TargetObjectKeyFormat.simplePrefix()`, but it is the same even if you do not specify `targetObjectKeyFormat` property.
	TargetObjectKeyFormat: s3.TargetObjectKeyFormat_SimplePrefix(),
})
```

Another option is `Date-based partitioning`.
If you choose this format, you can select either the event time or the delivery time of the log file as the date source used in the log format.
This format appears as follows:

```txt
[DestinationPrefix][SourceAccountId]/[SourceRegion]/[SourceBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]
```

```go
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
	ServerAccessLogsPrefix: jsii.String("logs"),
	TargetObjectKeyFormat: s3.TargetObjectKeyFormat_PartitionedPrefix(s3.PartitionDateSource_EVENT_TIME),
})
```

### Allowing access log delivery using a Bucket Policy (recommended)

When possible, it is recommended to use a bucket policy to grant access instead of
using ACLs. When the `@aws-cdk/aws-s3:serverAccessLogsUseBucketPolicy` feature flag
is enabled, this is done by default for server access logs. If S3 Server Access Logs
are the only logs delivered to your bucket (or if all other services logging to the
bucket support using bucket policy instead of ACLs), you can set object ownership
to [bucket owner enforced](#bucket-owner-enforced-recommended), as is recommended.

```go
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
})

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
	ServerAccessLogsPrefix: jsii.String("logs"),
})
```

The above code will create a new bucket policy if none exists or update the
existing bucket policy to allow access log delivery.

However, there could be an edge case if the `accessLogsBucket` also defines a bucket
policy resource using the L1 Construct. Although the mixing of L1 and L2 Constructs is not
recommended, there are no mechanisms in place to prevent users from doing this at the moment.

```go
bucketName := "amzn-s3-demo-bucket"
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
	BucketName: jsii.String(BucketName),
})

// Creating a bucket policy using L1
bucketPolicy := s3.NewCfnBucketPolicy(this, jsii.String("BucketPolicy"), &CfnBucketPolicyProps{
	Bucket: bucketName,
	PolicyDocument: map[string]interface{}{
		"Statement": []map[string]interface{}{
			map[string]interface{}{
				"Action": jsii.String("s3:*"),
				"Effect": jsii.String("Deny"),
				"Principal": map[string]*string{
					"AWS": jsii.String("*"),
				},
				"Resource": []*string{
					accessLogsBucket.bucketArn,
					fmt.Sprintf("%v/*", accessLogsBucket.bucketArn),
				},
			},
		},
		"Version": jsii.String("2012-10-17"),
	},
})

// 'serverAccessLogsBucket' will create a new L2 bucket policy
// to allow log delivery and overwrite the L1 bucket policy.
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
	ServerAccessLogsPrefix: jsii.String("logs"),
})
```

The above example uses the L2 Bucket Construct with the L1 CfnBucketPolicy Construct. However,
when `serverAccessLogsBucket` is set, a new L2 Bucket Policy resource will be created
which overwrites the permissions defined in the L1 Bucket Policy causing unintended
behaviours.

As noted above, we highly discourage the mixed usage of L1 and L2 Constructs. The recommended
approach would to define the bucket policy using `addToResourcePolicy` method.

```go
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
})

accessLogsBucket.AddToResourcePolicy(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("s3:*"),
	},
	Resources: []*string{
		accessLogsBucket.BucketArn,
		accessLogsBucket.ArnForObjects(jsii.String("*")),
	},
	Principals: []iPrincipal{
		iam.NewAnyPrincipal(),
	},
}))

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
	ServerAccessLogsPrefix: jsii.String("logs"),
})
```

Alternatively, users can use the L2 Bucket Policy Construct
`BucketPolicy.fromCfnBucketPolicy` to wrap around `CfnBucketPolicy` Construct. This will allow the subsequent bucket policy generated by `serverAccessLogsBucket` usage to append to the existing bucket policy instead of overwriting.

```go
bucketName := "amzn-s3-demo-bucket"
accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
	BucketName: jsii.String(BucketName),
})

bucketPolicy := s3.NewCfnBucketPolicy(this, jsii.String("BucketPolicy"), &CfnBucketPolicyProps{
	Bucket: bucketName,
	PolicyDocument: map[string]interface{}{
		"Statement": []map[string]interface{}{
			map[string]interface{}{
				"Action": jsii.String("s3:*"),
				"Effect": jsii.String("Deny"),
				"Principal": map[string]*string{
					"AWS": jsii.String("*"),
				},
				"Resource": []*string{
					accessLogsBucket.bucketArn,
					fmt.Sprintf("%v/*", accessLogsBucket.bucketArn),
				},
			},
		},
		"Version": jsii.String("2012-10-17"),
	},
})

// Wrap L1 Construct with L2 Bucket Policy Construct. Subsequent
// generated bucket policy to allow access log delivery would append
// to the current policy.
s3.BucketPolicy_FromCfnBucketPolicy(bucketPolicy)

bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ServerAccessLogsBucket: accessLogsBucket,
	ServerAccessLogsPrefix: jsii.String("logs"),
})
```

## S3 Inventory

An [inventory](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) contains a list of the objects in the source bucket and metadata for each object. The inventory lists are stored in the destination bucket as a CSV file compressed with GZIP, as an Apache optimized row columnar (ORC) file compressed with ZLIB, or as an Apache Parquet (Parquet) file compressed with Snappy.

You can configure multiple inventory lists for a bucket. You can configure what object metadata to include in the inventory, whether to list all object versions or only current versions, where to store the inventory list file output, and whether to generate the inventory on a daily or weekly basis.

```go
inventoryBucket := s3.NewBucket(this, jsii.String("InventoryBucket"))

dataBucket := s3.NewBucket(this, jsii.String("DataBucket"), &BucketProps{
	Inventories: []inventory{
		&inventory{
			Frequency: s3.InventoryFrequency_DAILY,
			IncludeObjectVersions: s3.InventoryObjectVersion_CURRENT,
			Destination: &InventoryDestination{
				Bucket: inventoryBucket,
			},
		},
		&inventory{
			Frequency: s3.InventoryFrequency_WEEKLY,
			IncludeObjectVersions: s3.InventoryObjectVersion_ALL,
			Destination: &InventoryDestination{
				Bucket: inventoryBucket,
				Prefix: jsii.String("with-all-versions"),
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
      "Resource": ["arn:aws:s3:::amzn-s3-demo-destination-bucket/*"]
    }
  ]
}
```

## Website redirection

You can use the two following properties to specify the bucket [redirection policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html#advanced-conditional-redirects). Please note that these methods cannot both be applied to the same bucket.

### Static redirection

You can statically redirect a to a given Bucket URL or any other host name with `websiteRedirect`:

```go
bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &BucketProps{
	WebsiteRedirect: &RedirectTarget{
		HostName: jsii.String("www.example.com"),
	},
})
```

### Routing rules

Alternatively, you can also define multiple `websiteRoutingRules`, to define complex, conditional redirections:

```go
bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &BucketProps{
	WebsiteRoutingRules: []routingRule{
		&routingRule{
			HostName: jsii.String("www.example.com"),
			HttpRedirectCode: jsii.String("302"),
			Protocol: s3.RedirectProtocol_HTTPS,
			ReplaceKey: s3.ReplaceKey_PrefixWith(jsii.String("test/")),
			Condition: &RoutingRuleCondition{
				HttpErrorCodeReturnedEquals: jsii.String("200"),
				KeyPrefixEquals: jsii.String("prefix"),
			},
		},
	},
})
```

## Filling the bucket as part of deployment

To put files into a bucket as part of a deployment (for example, to host a
website), see the `aws-cdk-lib/aws-s3-deployment` package, which provides a
resource that can do just that.

## The URL for objects

S3 provides two types of URLs for accessing objects via HTTP(S). Path-Style and
[Virtual Hosted-Style](https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html)
URL. Path-Style is a classic way and will be
[deprecated](https://aws.amazon.com/jp/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story).
We recommend to use Virtual Hosted-Style URL for newly made bucket.

You can generate both of them.

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))
bucket.UrlForObject(jsii.String("objectname")) // Path-Style URL
bucket.VirtualHostedUrlForObject(jsii.String("objectname")) // Virtual Hosted-Style URL
bucket.VirtualHostedUrlForObject(jsii.String("objectname"), &VirtualHostedStyleUrlOptions{
	Regional: jsii.Boolean(false),
})
```

## Object Ownership

You can use one of following properties to specify the bucket [object Ownership](https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html).

### Object writer

The Uploading account will own the object.

```go
s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ObjectOwnership: s3.ObjectOwnership_OBJECT_WRITER,
})
```

### Bucket owner preferred

The bucket owner will own the object if the object is uploaded with the bucket-owner-full-control canned ACL. Without this setting and canned ACL, the object is uploaded and remains owned by the uploading account.

```go
s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_PREFERRED,
})
```

### Bucket owner enforced (recommended)

ACLs are disabled, and the bucket owner automatically owns and has full control
over every object in the bucket. ACLs no longer affect permissions to data in the
S3 bucket. The bucket uses policies to define access control.

```go
s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
})
```

Some services may not not support log delivery to buckets that have object ownership
set to bucket owner enforced, such as
[S3 buckets using ACLs](#allowing-access-log-delivery-using-a-bucket-policy-recommended)
or [CloudFront Distributions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html#AccessLogsBucketAndFileOwnership).

## Bucket deletion

When a bucket is removed from a stack (or the stack is deleted), the S3
bucket will be removed according to its removal policy (which by default will
simply orphan the bucket and leave it in your AWS account). If the removal
policy is set to `RemovalPolicy.DESTROY`, the bucket will be deleted as long
as it does not contain any objects.

To override this and force all objects to get deleted during bucket deletion,
enable the`autoDeleteObjects` option.

When `autoDeleteObjects` is enabled, `s3:PutBucketPolicy` is added to the bucket policy. This is done to allow the custom resource this feature is built on to add a deny policy for `s3:PutObject` to the bucket policy when a delete stack event occurs. Adding this deny policy prevents new objects from being written to the bucket. Doing this prevents race conditions with external bucket writers during the deletion process.

```go
bucket := s3.NewBucket(this, jsii.String("MyTempFileBucket"), &BucketProps{
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
	AutoDeleteObjects: jsii.Boolean(true),
})
```

**Warning** if you have deployed a bucket with `autoDeleteObjects: true`,
switching this to `false` in a CDK version *before* `1.126.0` will lead to
all objects in the bucket being deleted. Be sure to update your bucket resources
by deploying with CDK version `1.126.0` or later **before** switching this value to `false`.

## Transfer Acceleration

[Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/transfer-acceleration.html) can be configured to enable fast, easy, and secure transfers of files over long distances:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	TransferAcceleration: jsii.Boolean(true),
})
```

To access the bucket that is enabled for Transfer Acceleration, you must use a special endpoint. The URL can be generated using method `transferAccelerationUrlForObject`:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	TransferAcceleration: jsii.Boolean(true),
})
bucket.TransferAccelerationUrlForObject(jsii.String("objectname"))
```

## Intelligent Tiering

[Intelligent Tiering](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html) can be configured to automatically move files to glacier:

```go
s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	IntelligentTieringConfigurations: []intelligentTieringConfiguration{
		&intelligentTieringConfiguration{
			Name: jsii.String("foo"),
			Prefix: jsii.String("folder/name"),
			ArchiveAccessTierTime: awscdk.Duration_Days(jsii.Number(90)),
			DeepArchiveAccessTierTime: awscdk.Duration_*Days(jsii.Number(180)),
			Tags: []tag{
				&tag{
					Key: jsii.String("tagname"),
					Value: jsii.String("tagvalue"),
				},
			},
		},
	},
})
```

## Lifecycle Rule

[Managing lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html) can be configured transition or expiration actions.

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	LifecycleRules: []lifecycleRule{
		&lifecycleRule{
			AbortIncompleteMultipartUploadAfter: awscdk.Duration_Minutes(jsii.Number(30)),
			Enabled: jsii.Boolean(false),
			Expiration: awscdk.Duration_Days(jsii.Number(30)),
			ExpirationDate: NewDate(),
			ExpiredObjectDeleteMarker: jsii.Boolean(false),
			Id: jsii.String("id"),
			NoncurrentVersionExpiration: awscdk.Duration_*Days(jsii.Number(30)),

			// the properties below are optional
			NoncurrentVersionsToRetain: jsii.Number(123),
			NoncurrentVersionTransitions: []noncurrentVersionTransition{
				&noncurrentVersionTransition{
					StorageClass: s3.StorageClass_GLACIER(),
					TransitionAfter: awscdk.Duration_*Days(jsii.Number(30)),

					// the properties below are optional
					NoncurrentVersionsToRetain: jsii.Number(123),
				},
			},
			ObjectSizeGreaterThan: jsii.Number(500),
			Prefix: jsii.String("prefix"),
			ObjectSizeLessThan: jsii.Number(10000),
			Transitions: []transition{
				&transition{
					StorageClass: s3.StorageClass_GLACIER(),

					// exactly one of transitionAfter or transitionDate must be specified
					TransitionAfter: awscdk.Duration_*Days(jsii.Number(30)),
				},
			},
		},
	},
})
```

To indicate which default minimum object size behavior is applied to the lifecycle configuration, use the
`transitionDefaultMinimumObjectSize` property.

The default value of the property before September 2024 is `TransitionDefaultMinimumObjectSize.VARIES_BY_STORAGE_CLASS`
that allows objects smaller than 128 KB to be transitioned only to the S3 Glacier and S3 Glacier Deep Archive storage classes,
otherwise `TransitionDefaultMinimumObjectSize.ALL_STORAGE_CLASSES_128_K` that prevents objects smaller than 128 KB from being
transitioned to any storage class.

To customize the minimum object size for any transition you
can add a filter that specifies a custom `objectSizeGreaterThan` or `objectSizeLessThan` for `lifecycleRules`
property. Custom filters always take precedence over the default transition behavior.

```go
s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	TransitionDefaultMinimumObjectSize: s3.TransitionDefaultMinimumObjectSize_VARIES_BY_STORAGE_CLASS,
	LifecycleRules: []lifecycleRule{
		&lifecycleRule{
			Transitions: []transition{
				&transition{
					StorageClass: s3.StorageClass_DEEP_ARCHIVE(),
					TransitionAfter: awscdk.Duration_Days(jsii.Number(30)),
				},
			},
		},
		&lifecycleRule{
			ObjectSizeLessThan: jsii.Number(300000),
			ObjectSizeGreaterThan: jsii.Number(200000),
			Transitions: []*transition{
				&transition{
					StorageClass: s3.StorageClass_ONE_ZONE_INFREQUENT_ACCESS(),
					TransitionAfter: awscdk.Duration_*Days(jsii.Number(30)),
				},
			},
		},
	},
})
```

## Object Lock Configuration

[Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html)
can be configured to enable a write-once-read-many model for an S3 bucket. Object Lock must be
configured when a bucket is created; if a bucket is created without Object Lock, it cannot be
enabled later via the CDK.

Object Lock can be enabled on an S3 bucket by specifying:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	ObjectLockEnabled: jsii.Boolean(true),
})
```

Usually, it is desired to not just enable Object Lock for a bucket but to also configure a
[retention mode](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-retention-modes)
and a [retention period](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-retention-periods).
These can be specified by providing `objectLockDefaultRetention`:

```go
// Configure for governance mode with a duration of 7 years
// Configure for governance mode with a duration of 7 years
s3.NewBucket(this, jsii.String("Bucket1"), &BucketProps{
	ObjectLockDefaultRetention: s3.ObjectLockRetention_Governance(awscdk.Duration_Days(jsii.Number(7 * 365))),
})

// Configure for compliance mode with a duration of 1 year
// Configure for compliance mode with a duration of 1 year
s3.NewBucket(this, jsii.String("Bucket2"), &BucketProps{
	ObjectLockDefaultRetention: s3.ObjectLockRetention_Compliance(awscdk.Duration_*Days(jsii.Number(365))),
})
```

## Replicating Objects

You can use [replicating objects](https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication.html) to enable automatic, asynchronous copying of objects across Amazon S3 buckets.
Buckets that are configured for object replication can be owned by the same AWS account or by different accounts.
You can replicate objects to a single destination bucket or to multiple destination buckets.
The destination buckets can be in different AWS Regions or within the same Region as the source bucket.

To replicate objects to a destination bucket, you can specify the `replicationRules` property:

```go
var destinationBucket1 iBucket
var destinationBucket2 iBucket
var replicationRole iRole
var encryptionKey iKey
var destinationEncryptionKey iKey


sourceBucket := s3.NewBucket(this, jsii.String("SourceBucket"), &BucketProps{
	// Versioning must be enabled on both the source and destination bucket
	Versioned: jsii.Boolean(true),
	// Optional. Specify the KMS key to use for encrypts objects in the source bucket.
	EncryptionKey: EncryptionKey,
	// Optional. If not specified, a new role will be created.
	ReplicationRole: ReplicationRole,
	ReplicationRules: []replicationRule{
		&replicationRule{
			// The destination bucket for the replication rule.
			Destination: destinationBucket1,
			// The priority of the rule.
			// Amazon S3 will attempt to replicate objects according to all replication rules.
			// However, if there are two or more rules with the same destination bucket, then objects will be replicated according to the rule with the highest priority.
			// The higher the number, the higher the priority.
			// It is essential to specify priority explicitly when the replication configuration has multiple rules.
			Priority: jsii.Number(1),
		},
		&replicationRule{
			Destination: destinationBucket2,
			Priority: jsii.Number(2),
			// Whether to specify S3 Replication Time Control (S3 RTC).
			// S3 RTC replicates most objects that you upload to Amazon S3 in seconds,
			// and 99.99 percent of those objects within specified time.
			ReplicationTimeControl: s3.ReplicationTimeValue_FIFTEEN_MINUTES(),
			// Whether to enable replication metrics about S3 RTC.
			// If set, metrics will be output to indicate whether replication by S3 RTC took longer than the configured time.
			Metrics: s3.ReplicationTimeValue_FIFTEEN_MINUTES(),
			// The kms key to use for the destination bucket.
			KmsKey: destinationEncryptionKey,
			// The storage class to use for the destination bucket.
			StorageClass: s3.StorageClass_INFREQUENT_ACCESS(),
			// Whether to replicate objects with SSE-KMS encryption.
			SseKmsEncryptedObjects: jsii.Boolean(false),
			// Whether to replicate modifications on replicas.
			ReplicaModifications: jsii.Boolean(true),
			// Whether to replicate delete markers.
			// This property cannot be enabled if the replication rule has a tag filter.
			DeleteMarkerReplication: jsii.Boolean(false),
			// The ID of the rule.
			Id: jsii.String("full-settings-rule"),
			// The object filter for the rule.
			Filter: &Filter{
				// The prefix filter for the rule.
				Prefix: jsii.String("prefix"),
				// The tag filter for the rule.
				Tags: []tag{
					&tag{
						Key: jsii.String("tagKey"),
						Value: jsii.String("tagValue"),
					},
				},
			},
		},
	},
})

// Grant permissions to the replication role.
// This method is not required if you choose to use an auto-generated replication role or manually grant permissions.
sourceBucket.GrantReplicationPermission(replicationRole, &GrantReplicationPermissionProps{
	// Optional. Specify the KMS key to use for decrypting objects in the source bucket.
	SourceDecryptionKey: encryptionKey,
	Destinations: []grantReplicationPermissionDestinationProps{
		&grantReplicationPermissionDestinationProps{
			Bucket: destinationBucket1,
		},
		&grantReplicationPermissionDestinationProps{
			Bucket: destinationBucket2,
			EncryptionKey: destinationEncryptionKey,
		},
	},
})
```

### Cross Account Replication

You can also set a destination bucket from a different account as the replication destination.

In this case, the bucket policy for the destination bucket is required, to configure it through CDK use  `addReplicationPolicy()` method to add bucket policy on destination bucket.
In a cross-account scenario, where the source and destination buckets are owned by different AWS accounts, you can use a KMS key to encrypt object replicas. However, the KMS key owner must grant the source bucket owner permission to use the KMS key.
For more information, please refer to https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication-walkthrough-2.html .

> **NOTE:** AWS managed keys don't allow cross-account use, and therefore can't be used to perform cross-account replication.

If you need to override the bucket ownership to destination account pass the account value to the method to provide permissions to override bucket owner.
`addReplicationPolicy(bucket.replicationRoleArn, true, '11111111111')`;

However, if the destination bucket is a referenced bucket, CDK cannot set the bucket policy,
so you will need to [configure the necessary bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/userguide/replication-walkthrough-2.html) separately.

```go
// The destination bucket in a different account.
var destinationBucket iBucket
var replicationRole iRole

sourceBucket := s3.NewBucket(this, jsii.String("SourceBucket"), &BucketProps{
	Versioned: jsii.Boolean(true),
	// Optional. If not specified, a new role will be created.
	ReplicationRole: ReplicationRole,
	ReplicationRules: []replicationRule{
		&replicationRule{
			Destination: destinationBucket,
			Priority: jsii.Number(1),
			// Whether to want to change replica ownership to the AWS account that owns the destination bucket.
			// The replicas are owned by same AWS account that owns the source object by default.
			AccessControlTransition: jsii.Boolean(true),
		},
	},
})

//Add permissions to the destination after replication role is created
if sourceBucket.ReplicationRoleArn {
	destinationBucket.AddReplicationPolicy(sourceBucket.ReplicationRoleArn, jsii.Boolean(true), jsii.String("111111111111"))
}
```

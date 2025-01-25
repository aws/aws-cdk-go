# CloudFront Origins for the CDK CloudFront Library

This library contains convenience methods for defining origins for a CloudFront distribution. You can use this library to create origins from
S3 buckets, Elastic Load Balancing v2 load balancers, or any other domain name.

## S3 Bucket

An S3 bucket can be used as an origin. An S3 bucket origin can either be configured using a standard S3 bucket or using a S3 bucket that's configured as a website endpoint (see AWS docs for [Using an S3 Bucket](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistS3AndCustomOrigins.html#using-s3-as-origin)).

> Note: `S3Origin` has been deprecated. Use `S3BucketOrigin` for standard S3 origins and `S3StaticWebsiteOrigin` for static website S3 origins.

### Standard S3 Bucket

To set up an origin using a standard S3 bucket, use the `S3BucketOrigin` class. The bucket
is handled as a bucket origin and
CloudFront's redirect and error handling will be used. It is recommended to use `S3BucketOrigin.withOriginAccessControl()` to configure OAC for your origin.

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket),
	},
})
```

> Note: When you use CloudFront OAC with Amazon S3 bucket origins, you must set Amazon S3 Object Ownership to Bucket owner enforced (the default for new Amazon S3 buckets). If you require ACLs, use the Bucket owner preferred setting to maintain control over objects uploaded via CloudFront.

### S3 Bucket Configured as a Website Endpoint

To set up an origin using an S3 bucket configured as a website endpoint, use the `S3StaticWebsiteOrigin` class. When the bucket is configured as a
website endpoint, the bucket is treated as an HTTP origin,
and the distribution can use built-in S3 redirects and S3 custom error pages.

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3StaticWebsiteOrigin(myBucket),
	},
})
```

### Restricting access to a standard S3 Origin

CloudFront provides two ways to send authenticated requests to a standard Amazon S3 origin:

* origin access control (OAC) and
* origin access identity (OAI)

OAI is considered legacy due to limited functionality and regional
limitations, whereas OAC is recommended because it supports all Amazon S3
buckets in all AWS Regions, Amazon S3 server-side encryption with AWS KMS (SSE-KMS), and dynamic requests (PUT and DELETE) to Amazon S3. Additionally,
OAC provides stronger security posture with short term credentials,
and more frequent credential rotations as compared to OAI. OAI and OAC can be used in conjunction with a bucket that is not public to
require that your users access your content using CloudFront URLs and not S3 URLs directly.

See AWS docs on [Restricting access to an Amazon S3 Origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) for more details.

> Note: OAC and OAI can only be used with an regular S3 bucket origin (not a bucket configured as a website endpoint).

The `S3BucketOrigin` class supports creating a standard S3 origin with OAC, OAI, and no access control (using your bucket access settings) via
the `withOriginAccessControl()`, `withOriginAccessIdentity()`, and `withBucketDefaults()` methods respectively.

#### Setting up a new origin access control (OAC)

Setup a standard S3 origin with origin access control as follows:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket),
	},
})
```

When creating a standard S3 origin using `origins.S3BucketOrigin.withOriginAccessControl()`, an [Origin Access Control resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html) is automatically created with the origin type set to `s3` and signing behavior set to `always`.

You can grant read, list, write or delete access to the OAC using the `originAccessLevels` property:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
	OriginAccessLevels: []accessLevel{
		cloudfront.*accessLevel_READ,
		cloudfront.*accessLevel_WRITE,
		cloudfront.*accessLevel_DELETE,
	},
})
```

For details of list permission, see [Setting up OAC with LIST permission](#setting-up-oac-with-list-permission).

You can also pass in a custom S3 origin access control:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
oac := cloudfront.NewS3OriginAccessControl(this, jsii.String("MyOAC"), &S3OriginAccessControlProps{
	Signing: cloudfront.Signing_SIGV4_NO_OVERRIDE(),
})
s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
	OriginAccessControl: oac,
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: s3Origin,
	},
})
```

An existing S3 origin access control can be imported using the `fromOriginAccessControlId` method:

```go
importedOAC := cloudfront.S3OriginAccessControl_FromOriginAccessControlId(this, jsii.String("myImportedOAC"), jsii.String("ABC123ABC123AB"))
```

> [Note](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html): When you use OAC with S3
> bucket origins, the bucket's object ownership must be either set to Bucket owner enforced (default for new S3 buckets) or Bucket owner preferred (only if you require ACLs).

#### Setting up OAC with a SSE-KMS encrypted S3 origin

If the objects in the S3 bucket origin are encrypted using server-side encryption with
AWS Key Management Service (SSE-KMS), the OAC must have permission to use the KMS key.

Setting up a standard S3 origin using `S3BucketOrigin.withOriginAccessControl()` will automatically add the statement to the KMS key policy
to give the OAC permission to use the KMS key.

```go
import kms "github.com/aws/aws-cdk-go/awscdk"


myKmsKey := kms.NewKey(this, jsii.String("myKMSKey"))
myBucket := s3.NewBucket(this, jsii.String("mySSEKMSEncryptedBucket"), &BucketProps{
	Encryption: s3.BucketEncryption_KMS,
	EncryptionKey: myKmsKey,
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket),
	},
})
```

##### Scoping down the key policy

I saw this warning message during synth time. What do I do?

```text
To avoid a circular dependency between the KMS key, Bucket, and Distribution during the initial deployment, a wildcard is used in the Key policy condition to match all Distribution IDs.
After deploying once, it is strongly recommended to further scope down the policy for best security practices by following the guidance in the "Using OAC for a SSE-KMS encrypted S3 origin" section in the module README.
```

If the S3 bucket has an `encryptionKey` defined, `S3BucketOrigin.withOriginAccessControl()`
will automatically add the following policy statement to the KMS key policy to allow CloudFront read-only access (unless otherwise specified in the `originAccessLevels` property).

```json
{
    "Statement": {
        "Effect": "Allow",
        "Principal": {
            "Service": "cloudfront.amazonaws.com"
        },
        "Action": "kms:Decrypt",
        "Resource": "*",
        "Condition": {
            "ArnLike": {
                "AWS:SourceArn": "arn:aws:cloudfront::<account ID>:distribution/*"
            }
        }
    }
}
```

This policy uses a wildcard to match all distribution IDs in the account instead of referencing the specific distribution ID to resolve the circular dependency. The policy statement is not as scoped down as the example in the AWS CloudFront docs (see [SSE-KMS section](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html#create-oac-overview-s3)).

After you have deployed the Distribution, you should follow these steps to only grant permissions to the specific distribution according to AWS best practices:

**Step 1.** Copy the key policy

**Step 2.** Use an escape hatch to update the policy statement condition so that

```json
  "Condition": {
      "ArnLike": {
          "AWS:SourceArn": "arn:aws:cloudfront::<account ID>:distribution/*"
      }
  }
```

...becomes...

```json
  "Condition": {
      "StringEquals": {
          "AWS:SourceArn": "arn:aws:cloudfront::111122223333:distribution/<CloudFront distribution ID>"
      }
  }
```

> Note the change of condition operator from `ArnLike` to `StringEquals` in addition to replacing the wildcard (`*`) with the distribution ID.

To set the key policy using an escape hatch:

```go
import kms "github.com/aws/aws-cdk-go/awscdk"


kmsKey := kms.NewKey(this, jsii.String("myKMSKey"))
myBucket := s3.NewBucket(this, jsii.String("mySSEKMSEncryptedBucket"), &BucketProps{
	Encryption: s3.BucketEncryption_KMS,
	EncryptionKey: kmsKey,
	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket),
	},
})

// Add the following to scope down the key policy
scopedDownKeyPolicy := map[string]interface{}{
	"Version": jsii.String("2012-10-17"),
	"Statement": []interface{}{
		map[string]interface{}{
			"Effect": jsii.String("Allow"),
			"Principal": map[string]*string{
				"AWS": jsii.String("arn:aws:iam::111122223333:root"),
			},
			"Action": jsii.String("kms:*"),
			"Resource": jsii.String("*"),
		},
		map[string]interface{}{
			"Effect": jsii.String("Allow"),
			"Principal": map[string]*string{
				"Service": jsii.String("cloudfront.amazonaws.com"),
			},
			"Action": []*string{
				jsii.String("kms:Decrypt"),
				jsii.String("kms:Encrypt"),
				jsii.String("kms:GenerateDataKey*"),
			},
			"Resource": jsii.String("*"),
			"Condition": map[string]map[string]*string{
				"StringEquals": map[string]*string{
					"AWS:SourceArn": jsii.String("arn:aws:cloudfront::111122223333:distribution/<CloudFront distribution ID>"),
				},
			},
		},
	},
}
cfnKey := (kmsKey.Node.defaultChild.(cfnKey))
cfnKey.KeyPolicy = scopedDownKeyPolicy
```

**Step 3.** Deploy the stack

> Tip: Run `cdk diff` before deploying to verify the
> changes to your stack.

**Step 4.** Verify your final key policy includes the following statement after deploying:

```json
{
    "Effect": "Allow",
    "Principal": {
        "Service": [
            "cloudfront.amazonaws.com"
        ]
     },
    "Action": [
        "kms:Decrypt",
        "kms:Encrypt",
        "kms:GenerateDataKey*"
    ],
    "Resource": "*",
    "Condition": {
            "StringEquals": {
                "AWS:SourceArn": "arn:aws:cloudfront::111122223333:distribution/<CloudFront distribution ID>"
            }
        }
}
```

##### Updating imported key policies

If you are using an imported KMS key to encrypt your S3 bucket and want to use OAC, you will need to update the
key policy manually to allow CloudFront to use the key. Like most imported resources, CDK apps cannot modify the configuration of imported keys.

After deploying the distribution, add the following policy statement to your key policy to allow CloudFront OAC to access your KMS key for SSE-KMS:

```json
{
    "Sid": "AllowCloudFrontServicePrincipalSSE-KMS",
    "Effect": "Allow",
    "Principal": {
        "Service": [
            "cloudfront.amazonaws.com"
        ]
     },
    "Action": [
        "kms:Decrypt",
        "kms:Encrypt",
        "kms:GenerateDataKey*"
    ],
    "Resource": "*",
    "Condition": {
            "StringEquals": {
                "AWS:SourceArn": "arn:aws:cloudfront::111122223333:distribution/<CloudFront distribution ID>"
            }
        }
}
```

See CloudFront docs on [SSE-KMS](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html#create-oac-overview-s3) for more details.

#### Setting up OAC with imported S3 buckets

If you are using an imported bucket for your S3 Origin and want to use OAC,
you will need to update
the S3 bucket policy manually to allow the OAC to access the S3 origin. Like most imported resources, CDK apps cannot modify the configuration of imported buckets.

After deploying the distribution, add the following
policy statement to your
S3 bucket to allow CloudFront read-only access
(or additional S3 permissions as required):

```json
{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Principal": {
            "Service": "cloudfront.amazonaws.com"
        },
        "Action": "s3:GetObject",
        "Resource": "arn:aws:s3:::<S3 bucket name>/*",
        "Condition": {
            "StringEquals": {
                "AWS:SourceArn": "arn:aws:cloudfront::111122223333:distribution/<CloudFront distribution ID>"
            }
        }
    }
}
```

See CloudFront docs on [Giving the origin access control permission to access the S3 bucket](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html#create-oac-overview-s3) for more details.

> Note: If your bucket previously used OAI, you will need to manually remove the policy statement
> that gives the OAI access to your bucket after setting up OAC.

#### Setting up OAC with LIST permission

By default, S3 origin returns 403 Forbidden HTTP response when the requested object does not exist.
When you want to receive 404 Not Found, specify `AccessLevel.LIST` in `originAccessLevels` to add `s3:ListBucket` permission in the bucket policy.

This is useful to distinguish between responses blocked by WAF (403) and responses where the file does not exist (404).

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
	OriginAccessLevels: []accessLevel{
		cloudfront.*accessLevel_READ,
		cloudfront.*accessLevel_LIST,
	},
})
cloudfront.NewDistribution(this, jsii.String("distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: s3Origin,
	},
	DefaultRootObject: jsii.String("index.html"),
})
```

When the origin is associated to the default behavior, it is highly recommended to specify `defaultRootObject` distribution property.
Without it, the root path `https://xxxx.cloudfront.net/` will return the list of the S3 object keys.

#### Setting up an OAI (legacy)

Setup an S3 origin with origin access identity (legacy) as follows:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessIdentity(myBucket),
	},
})
```

You can also pass in a custom S3 origin access identity:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
myOai := cloudfront.NewOriginAccessIdentity(this, jsii.String("myOAI"), &OriginAccessIdentityProps{
	Comment: jsii.String("My custom OAI"),
})
s3Origin := origins.S3BucketOrigin_WithOriginAccessIdentity(myBucket, &S3BucketOriginWithOAIProps{
	OriginAccessIdentity: myOai,
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: s3Origin,
	},
})
```

#### Setting up OAI with imported S3 buckets (legacy)

If you are using an imported bucket for your S3 Origin and want to use OAI,
you will need to update
the S3 bucket policy manually to allow the OAI to access the S3 origin. Like most imported resources, CDK apps cannot modify the configuration of imported buckets.

Add the following
policy statement to your
S3 bucket to allow the OAI read access:

```json
{
    "Version": "2012-10-17",
    "Id": "PolicyForCloudFrontPrivateContent",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity <origin access identity ID>"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::<S3 bucket name>/*"
        }
    ]
}
```

See AWS docs on [Giving an origin access identity permission to read files in the Amazon S3 bucket](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html#private-content-restricting-access-to-s3-oai) for more details.

### Setting up a S3 origin with no origin access control

To setup a standard S3 origin with no access control (no OAI nor OAC), use `origins.S3BucketOrigin.withBucketDefaults()`:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithBucketDefaults(myBucket),
	},
})
```

### Migrating from OAI to OAC

If you are currently using OAI for your S3 origin and wish to migrate to OAC,
replace the `S3Origin` construct (deprecated) with `S3BucketOrigin.withOriginAccessControl()` which automatically
creates and sets up an OAC for you.

Existing setup using OAI and `S3Origin`:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
s3Origin := origins.NewS3Origin(myBucket)
distribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: s3Origin,
	},
})
```

**Step 1:**

To ensure CloudFront doesn't lose access to the bucket during the transition, add a statement to bucket policy to grant OAC access to the S3 origin. Deploy the stack. If you are okay with downtime during the transition, you can skip this step.

> Tip: Run `cdk diff` before deploying to verify the
> changes to your stack.

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


stack := awscdk.Newstack()
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
s3Origin := origins.NewS3Origin(myBucket)
distribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: s3Origin,
	},
})

// Construct the bucket policy statement
distributionArn := stack.FormatArn(&ArnComponents{
	Service: jsii.String("cloudfront"),
	Region: jsii.String(""),
	Resource: jsii.String("distribution"),
	ResourceName: distribution.DistributionId,
	ArnFormat: cdk.ArnFormat_SLASH_RESOURCE_NAME,
})

cloudfrontSP := iam.NewServicePrincipal(jsii.String("cloudfront.amazonaws.com"))

oacBucketPolicyStatement := iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Principals: []iPrincipal{
		cloudfrontSP,
	},
	Actions: []*string{
		jsii.String("s3:GetObject"),
	},
	Resources: []*string{
		myBucket.ArnForObjects(jsii.String("*")),
	},
	Conditions: map[string]interface{}{
		"StringEquals": map[string]*string{
			"AWS:SourceArn": distributionArn,
		},
	},
})

// Add statement to bucket policy
myBucket.AddToResourcePolicy(oacBucketPolicyStatement)
```

The following changes will take place:

1. The bucket policy will be modified to grant the CloudFront distribution access. At this point the bucket policy allows both an OAI and an OAC to access the S3 origin.

**Step 2:**

Replace `S3Origin` with `S3BucketOrigin.withOriginAccessControl()`, which creates an OAC and attaches it to the distribution. You can remove the code from Step 1 which updated the bucket policy, as `S3BucketOrigin.withOriginAccessControl()` updates the bucket policy automatically with the same statement when defined in the `Distribution` (no net difference).

Run `cdk diff` before deploying to verify the changes to your stack.

```go
bucket := s3.NewBucket(this, jsii.String("Bucket"))
s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(bucket)
distribution := cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: s3Origin,
	},
})
```

The following changes will take place:

1. A `AWS::CloudFront::OriginAccessControl` resource will be created.
2. The `Origin` property of the `AWS::CloudFront::Distribution` will set [`OriginAccessControlId`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-originaccesscontrolid) to the OAC ID after it is created. It will also set [`S3OriginConfig`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-s3originconfig.html#aws-properties-cloudfront-distribution-s3originconfig-properties) to `{"OriginAccessIdentity": ""}`, which deletes the origin access identity from the existing distribution.
3. The `AWS::CloudFront::CloudFrontOriginAccessIdentity` resource will be deleted.

**Will migrating from OAI to OAC cause any resource replacement?**

No, following the migration steps does not cause any replacement of the existing `AWS::CloudFront::Distribution`, `AWS::S3::Bucket` nor `AWS::S3::BucketPolicy` resources. It will modify the bucket policy, create a `AWS::CloudFront::OriginAccessControl` resource, and delete the existing `AWS::CloudFront::CloudFrontOriginAccessIdentity`.

**Will migrating from OAI to OAC have any availability implications for my application?**

Updates to bucket policies are eventually consistent. Therefore, removing OAI permissions and setting up OAC in the same CloudFormation stack deployment is not recommended as it may cause downtime where CloudFront loses access to the bucket. Following the steps outlined above lowers the risk of downtime as the bucket policy is updated to have both OAI and OAC permissions, then in a subsequent deployment, the OAI permissions are removed.

For more information, see [Migrating from origin access identity (OAI) to origin access control (OAC)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html#migrate-from-oai-to-oac).

### Adding Custom Headers

You can configure CloudFront to add custom headers to the requests that it sends to your origin. These custom headers enable you to send and gather information from your origin that you don’t get with typical viewer requests. These headers can even be customized for each origin. CloudFront supports custom headers for both for custom and Amazon S3 origins.

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
			CustomHeaders: map[string]*string{
				"Foo": jsii.String("bar"),
			},
		}),
	},
})
```

## ELBv2 Load Balancer

An Elastic Load Balancing (ELB) v2 load balancer may be used as an origin. In order for a load balancer to serve as an origin, it must be publicly
accessible (`internetFacing` is true). Both Application and Network load balancers are supported.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

// Create an application load balancer in a VPC. 'internetFacing' must be 'true'
// for CloudFront to access the load balancer and use it as an origin.
lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewLoadBalancerV2Origin(lb),
	},
})
```

The origin can also be customized to respond on different ports, have different connection properties, etc.

```go
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

var loadBalancer applicationLoadBalancer

origin := origins.NewLoadBalancerV2Origin(loadBalancer, &LoadBalancerV2OriginProps{
	ConnectionAttempts: jsii.Number(3),
	ConnectionTimeout: awscdk.Duration_Seconds(jsii.Number(5)),
	ReadTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
	KeepaliveTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
	ProtocolPolicy: cloudfront.OriginProtocolPolicy_MATCH_VIEWER,
})
```

Note that the `readTimeout` and `keepaliveTimeout` properties can extend their values over 60 seconds only if a limit increase request for CloudFront origin response timeout
quota has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time. Consider that this value is
still limited to a maximum value of 180 seconds, which is a hard limit for that quota.

## From an HTTP endpoint

Origins can also be created from any other HTTP endpoint, given the domain name, and optionally, other origin properties.

```go
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
})
```

See the documentation of `aws-cdk-lib/aws-cloudfront` for more information.

## Failover Origins (Origin Groups)

You can set up CloudFront with origin failover for scenarios that require high availability.
To get started, you create an origin group with two origins: a primary and a secondary.
If the primary origin is unavailable, or returns specific HTTP response status codes that indicate a failure,
CloudFront automatically switches to the secondary origin.
You achieve that behavior in the CDK using the `OriginGroup` class:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewOriginGroup(&OriginGroupProps{
			PrimaryOrigin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket),
			FallbackOrigin: origins.NewHttpOrigin(jsii.String("www.example.com")),
			// optional, defaults to: 500, 502, 503 and 504
			FallbackStatusCodes: []*f64{
				jsii.Number(404),
			},
		}),
	},
})
```

## From an API Gateway REST API

Origins can be created from an API Gateway REST API. It is recommended to use a
[regional API](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-endpoint-types.html) in this case. The origin path will automatically be set as the stage name.

```go
var api restApi

cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewRestApiOrigin(api),
	},
})
```

If you want to use a different origin path, you can specify it in the `originPath` property.

```go
var api restApi

cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewRestApiOrigin(api, &RestApiOriginProps{
			OriginPath: jsii.String("/custom-origin-path"),
		}),
	},
})
```

## From a Lambda Function URL

Lambda Function URLs enable direct invocation of Lambda functions via HTTP(S), without intermediaries. They can be set as CloudFront origins for streamlined function execution behind a CDN, leveraging caching and custom domains.

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var fn function

fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
	AuthType: lambda.FunctionUrlAuthType_NONE,
})

cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewFunctionUrlOrigin(fnUrl),
	},
})
```

### Lambda Function URL with Origin Access Control (OAC)

You can configure the Lambda Function URL with Origin Access Control (OAC) for enhanced security. When using OAC with Signing SIGV4_ALWAYS, it is recommended to set the Lambda Function URL authType to AWS_IAM to ensure proper authorization.

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var fn function


fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
	AuthType: lambda.FunctionUrlAuthType_AWS_IAM,
})

cloudfront.NewDistribution(this, jsii.String("MyDistribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.FunctionUrlOrigin_WithOriginAccessControl(fnUrl),
	},
})
```

If you want to explicitly add OAC for more customized access control, you can use the originAccessControl option as shown below.

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var fn function


fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
	AuthType: lambda.FunctionUrlAuthType_AWS_IAM,
})

// Define a custom OAC
oac := cloudfront.NewFunctionUrlOriginAccessControl(this, jsii.String("MyOAC"), &FunctionUrlOriginAccessControlProps{
	OriginAccessControlName: jsii.String("CustomLambdaOAC"),
	Signing: cloudfront.Signing_SIGV4_ALWAYS(),
})

// Set up Lambda Function URL with OAC in CloudFront Distribution
// Set up Lambda Function URL with OAC in CloudFront Distribution
cloudfront.NewDistribution(this, jsii.String("MyDistribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.FunctionUrlOrigin_WithOriginAccessControl(fnUrl, &FunctionUrlOriginWithOACProps{
			OriginAccessControl: oac,
		}),
	},
})
```

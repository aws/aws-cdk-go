# AWS S3 Deployment Construct Library

This library allows populating an S3 bucket with the contents of .zip files
from other S3 buckets or from local disk.

The following example defines a publicly accessible S3 bucket with web hosting
enabled and populates it from a local directory on disk.

```go
websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &BucketProps{
	WebsiteIndexDocument: jsii.String("index.html"),
	PublicReadAccess: jsii.Boolean(true),
})

s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(jsii.String("./website-dist")),
	},
	DestinationBucket: websiteBucket,
	DestinationKeyPrefix: jsii.String("web/static"),
})
```

This is what happens under the hood:

1. When this stack is deployed (either via `cdk deploy` or via CI/CD), the
   contents of the local `website-dist` directory will be archived and uploaded
   to an intermediary assets bucket (the `StagingBucket` of the CDK bootstrap stack).
   If there is more than one source, they will be individually uploaded.
2. The `BucketDeployment` construct synthesizes a Lambda-backed custom CloudFormation resource
   of type `Custom::CDKBucketDeployment` into the template. The source bucket/key
   is set to point to the assets bucket.
3. The custom resource invokes its associated Lambda function, which downloads the .zip archive,
   extracts it and issues `aws s3 sync --delete` against the destination bucket (in this case
   `websiteBucket`). If there is more than one source, the sources will be
   downloaded and merged pre-deployment at this step.

If you are referencing the filled bucket in another construct that depends on
the files already be there, be sure to use `deployment.deployedBucket`. This
will ensure the bucket deployment has finished before the resource that uses
the bucket is created:

```go
var websiteBucket bucket


deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
	},
	DestinationBucket: websiteBucket,
})

NewConstructThatReadsFromTheBucket(this, jsii.String("Consumer"), map[string]iBucket{
	// Use 'deployment.deployedBucket' instead of 'websiteBucket' here
	"bucket": deployment.deployedBucket,
})
```

It is also possible to add additional sources using the `addSource` method.

```go
var websiteBucket iBucket


deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(jsii.String("./website-dist")),
	},
	DestinationBucket: websiteBucket,
	DestinationKeyPrefix: jsii.String("web/static"),
})

deployment.AddSource(s3deploy.Source_Asset(jsii.String("./another-asset")))
```

For the Lambda function to download object(s) from the source bucket, besides the obvious
`s3:GetObject*` permissions, the Lambda's execution role needs the `kms:Decrypt` and `kms:DescribeKey`
permissions on the KMS key that is used to encrypt the bucket. By default, when the source bucket is
encrypted with the S3 managed key of the account, these permissions are granted by the key's
resource-based policy, so they do not need to be on the Lambda's execution role policy explicitly.
However, if the encryption key is not the s3 managed one, its resource-based policy is quite likely
to NOT grant such KMS permissions. In this situation, the Lambda execution will fail with an error
message like below:

```txt
download failed: ...
An error occurred (AccessDenied) when calling the GetObject operation:
User: *** is not authorized to perform: kms:Decrypt on the resource associated with this ciphertext
because no identity-based policy allows the kms:Decrypt action
```

When this happens, users can use the public `handlerRole` property of `BucketDeployment` to manually
add the KMS permissions:

```go
var destinationBucket bucket


deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployFiles"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(path.join(__dirname, jsii.String("source-files"))),
	},
	DestinationBucket: DestinationBucket,
})

deployment.HandlerRole.AddToPolicy(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("kms:Decrypt"),
		jsii.String("kms:DescribeKey"),
	},
	Effect: iam.Effect_ALLOW,
	Resources: []*string{
		jsii.String("<encryption key ARN>"),
	},
}))
```

The situation above could arise from the following scenarios:

* User created a customer managed KMS key and passed its ID to the `cdk bootstrap` command via
  the `--bootstrap-kms-key-id` CLI option.
  The [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#key-policy-default-allow-root-enable-iam)
  alone is not sufficient to grant the Lambda KMS permissions.
* A corporation uses its own custom CDK bootstrap process, which encrypts the CDK `StagingBucket`
  by a KMS key from a management account of the corporation's AWS Organization. In this cross-account
  access scenario, the KMS permissions must be explicitly present in the Lambda's execution role policy.
* One of the sources for the `BucketDeployment` comes from the `Source.bucket` static method, which
  points to a bucket whose encryption key is not the S3 managed one, and the resource-based policy
  of the encryption key is not sufficient to grant the Lambda `kms:Decrypt` and `kms:DescribeKey`
  permissions.

## Supported sources

The following source types are supported for bucket deployments:

* Local .zip file: `s3deploy.Source.asset('/path/to/local/file.zip')`
* Local directory: `s3deploy.Source.asset('/path/to/local/directory')`
* Another bucket: `s3deploy.Source.bucket(bucket, zipObjectKey)`
* String data: `s3deploy.Source.data('object-key.txt', 'hello, world!')`
  (supports [deploy-time values](#data-with-deploy-time-values))
* JSON data: `s3deploy.Source.jsonData('object-key.json', { json: 'object' })`
  (supports [deploy-time values](#data-with-deploy-time-values))
* YAML data: `s3deploy.Source.yamlData('object-key.yaml', { yaml: 'object' })`
  (supports [deploy-time values](#data-with-deploy-time-values))

To create a source from a single file, you can pass `AssetOptions` to exclude
all but a single file:

* Single file: `s3deploy.Source.asset('/path/to/local/directory', { exclude: ['**', '!onlyThisFile.txt'] })`

**IMPORTANT** The `aws-s3-deployment` module is only intended to be used with
zip files from trusted sources. Directories bundled by the CDK CLI (by using
`Source.asset()` on a directory) are safe. If you are using `Source.asset()` or
`Source.bucket()` to reference an existing zip file, make sure you trust the
file you are referencing. Zips from untrusted sources might be able to execute
arbitrary code in the Lambda Function used by this module, and use its permissions
to read or write unexpected files in the S3 bucket.

## Retain on Delete

By default, the contents of the destination bucket will **not** be deleted when the
`BucketDeployment` resource is removed from the stack or when the destination is
changed. You can use the option `retainOnDelete: false` to disable this behavior,
in which case the contents will be deleted.

Configuring this has a few implications you should be aware of:

* **Logical ID Changes**

  Changing the logical ID of the `BucketDeployment` construct, without changing the destination
  (for example due to refactoring, or intentional ID change) **will result in the deletion of the objects**.
  This is because CloudFormation will first create the new resource, which will have no affect,
  followed by a deletion of the old resource, which will cause a deletion of the objects,
  since the destination hasn't changed, and `retainOnDelete` is `false`.
* **Destination Changes**

  When the destination bucket or prefix is changed, all files in the previous destination will **first** be
  deleted and then uploaded to the new destination location. This could have availability implications
  on your users.

### General Recommendations

#### Shared Bucket

If the destination bucket **is not** dedicated to the specific `BucketDeployment` construct (i.e shared by other entities),
we recommend to always configure the `destinationKeyPrefix` property. This will prevent the deployment from
accidentally deleting data that wasn't uploaded by it.

#### Dedicated Bucket

If the destination bucket **is** dedicated, it might be reasonable to skip the prefix configuration,
in which case, we recommend to remove `retainOnDelete: false`, and instead, configure the
[`autoDeleteObjects`](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-s3-readme.html#bucket-deletion)
property on the destination bucket. This will avoid the logical ID problem mentioned above.

## Prune

By default, files in the destination bucket that don't exist in the source will be deleted
when the `BucketDeployment` resource is created or updated. You can use the option `prune: false` to disable
this behavior, in which case the files will not be deleted.

```go
var destinationBucket bucket

s3deploy.NewBucketDeployment(this, jsii.String("DeployMeWithoutDeletingFilesOnDestination"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
	},
	DestinationBucket: DestinationBucket,
	Prune: jsii.Boolean(false),
})
```

This option also enables you to
multiple bucket deployments for the same destination bucket & prefix,
each with its own characteristics. For example, you can set different cache-control headers
based on file extensions:

```go
var destinationBucket bucket

s3deploy.NewBucketDeployment(this, jsii.String("BucketDeployment"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(jsii.String("./website"), &AssetOptions{
			Exclude: []*string{
				jsii.String("index.html"),
			},
		}),
	},
	DestinationBucket: DestinationBucket,
	CacheControl: []cacheControl{
		s3deploy.*cacheControl_MaxAge(awscdk.Duration_Days(jsii.Number(365))),
		s3deploy.*cacheControl_Immutable(),
	},
	Prune: jsii.Boolean(false),
})

s3deploy.NewBucketDeployment(this, jsii.String("HTMLBucketDeployment"), &BucketDeploymentProps{
	Sources: []*iSource{
		s3deploy.Source_*Asset(jsii.String("./website"), &AssetOptions{
			Exclude: []*string{
				jsii.String("*"),
				jsii.String("!index.html"),
			},
		}),
	},
	DestinationBucket: DestinationBucket,
	CacheControl: []*cacheControl{
		s3deploy.*cacheControl_*MaxAge(awscdk.Duration_Seconds(jsii.Number(0))),
	},
	Prune: jsii.Boolean(false),
})
```

## Exclude and Include Filters

There are two points at which filters are evaluated in a deployment: asset bundling and the actual deployment. If you simply want to exclude files in the asset bundling process, you should leverage the `exclude` property of `AssetOptions` when defining your source:

```go
var destinationBucket bucket

s3deploy.NewBucketDeployment(this, jsii.String("HTMLBucketDeployment"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(jsii.String("./website"), &AssetOptions{
			Exclude: []*string{
				jsii.String("*"),
				jsii.String("!index.html"),
			},
		}),
	},
	DestinationBucket: DestinationBucket,
})
```

If you want to specify filters to be used in the deployment process, you can use the `exclude` and `include` filters on `BucketDeployment`.  If excluded, these files will not be deployed to the destination bucket. In addition, if the file already exists in the destination bucket, it will not be deleted if you are using the `prune` option:

```go
var destinationBucket bucket

s3deploy.NewBucketDeployment(this, jsii.String("DeployButExcludeSpecificFiles"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
	},
	DestinationBucket: DestinationBucket,
	Exclude: []*string{
		jsii.String("*.txt"),
	},
})
```

These filters follow the same format that is used for the AWS CLI.  See the CLI documentation for information on [Using Include and Exclude Filters](https://docs.aws.amazon.com/cli/latest/reference/s3/index.html#use-of-exclude-and-include-filters).

## Objects metadata

You can specify metadata to be set on all the objects in your deployment.
There are 2 types of metadata in S3: system-defined metadata and user-defined metadata.
System-defined metadata have a special purpose, for example cache-control defines how long to keep an object cached.
User-defined metadata are not used by S3 and keys always begin with `x-amz-meta-` (this prefix is added automatically).

System defined metadata keys include the following:

* cache-control (`--cache-control` in `aws s3 sync`)
* content-disposition (`--content-disposition` in `aws s3 sync`)
* content-encoding (`--content-encoding` in `aws s3 sync`)
* content-language (`--content-language` in `aws s3 sync`)
* content-type (`--content-type` in `aws s3 sync`)
* expires (`--expires` in `aws s3 sync`)
* x-amz-storage-class (`--storage-class` in `aws s3 sync`)
* x-amz-website-redirect-location (`--website-redirect` in `aws s3 sync`)
* x-amz-server-side-encryption (`--sse` in `aws s3 sync`)
* x-amz-server-side-encryption-aws-kms-key-id (`--sse-kms-key-id` in `aws s3 sync`)
* x-amz-server-side-encryption-customer-algorithm (`--sse-c-copy-source` in `aws s3 sync`)
* x-amz-acl (`--acl` in `aws s3 sync`)

You can find more information about system defined metadata keys in
[S3 PutObject documentation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
and [`aws s3 sync` documentation](https://docs.aws.amazon.com/cli/latest/reference/s3/sync.html).

```go
websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &BucketProps{
	WebsiteIndexDocument: jsii.String("index.html"),
	PublicReadAccess: jsii.Boolean(true),
})

s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(jsii.String("./website-dist")),
	},
	DestinationBucket: websiteBucket,
	DestinationKeyPrefix: jsii.String("web/static"),
	 // optional prefix in destination bucket
	Metadata: map[string]*string{
		"A": jsii.String("1"),
		"b": jsii.String("2"),
	},
	 // user-defined metadata

	// system-defined metadata
	ContentType: jsii.String("text/html"),
	ContentLanguage: jsii.String("en"),
	StorageClass: s3deploy.StorageClass_INTELLIGENT_TIERING,
	ServerSideEncryption: s3deploy.ServerSideEncryption_AES_256,
	CacheControl: []cacheControl{
		s3deploy.*cacheControl_SetPublic(),
		s3deploy.*cacheControl_MaxAge(awscdk.Duration_Hours(jsii.Number(1))),
	},
	AccessControl: s3.BucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
})
```

## CloudFront Invalidation

You can provide a CloudFront distribution and optional paths to invalidate after the bucket deployment finishes.

```go
import cloudfront "github.com/aws/aws-cdk-go/awscdk"
import origins "github.com/aws/aws-cdk-go/awscdk"


bucket := s3.NewBucket(this, jsii.String("Destination"))

// Handles buckets whether or not they are configured for website hosting.
distribution := cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(bucket),
	},
})

s3deploy.NewBucketDeployment(this, jsii.String("DeployWithInvalidation"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(jsii.String("./website-dist")),
	},
	DestinationBucket: bucket,
	Distribution: Distribution,
	DistributionPaths: []*string{
		jsii.String("/images/*.png"),
	},
})
```

## Signed Content Payloads

By default, deployment uses streaming uploads which set the `x-amz-content-sha256`
request header to `UNSIGNED-PAYLOAD` (matching default behavior of the AWS CLI tool).
In cases where bucket policy restrictions require signed content payloads, you can enable
generation of a signed `x-amz-content-sha256` request header with `signContent: true`.

```go
var bucket iBucket


s3deploy.NewBucketDeployment(this, jsii.String("DeployWithSignedPayloads"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(jsii.String("./website-dist")),
	},
	DestinationBucket: bucket,
	SignContent: jsii.Boolean(true),
})
```

## Size Limits

The default memory limit for the deployment resource is 128MiB. If you need to
copy larger files, you can use the `memoryLimit` configuration to increase the
size of the AWS Lambda resource handler.

The default ephemeral storage size for the deployment resource is 512MiB. If you
need to upload larger files, you may hit this limit. You can use the
`ephemeralStorageSize` configuration to increase the storage size of the AWS Lambda
resource handler.

> NOTE: a new AWS Lambda handler will be created in your stack for each combination
> of memory and storage size.

## JSON-Aware Source Processing

When using `Source.jsonData` with CDK Tokens (references to construct properties), you may need to enable the escaping option. This is particularly important when the referenced properties might contain special characters that require proper JSON escaping (like double quotes, line breaks, etc.).

```go
var bucket bucket
var param stringParameter


// Example with a secret value that contains double quotes
deployment := s3deploy.NewBucketDeployment(this, jsii.String("JsonDeployment"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_JsonData(jsii.String("config.json"), map[string]interface{}{
			"api_endpoint": jsii.String("https://api.example.com"),
			"secretValue": param.stringValue,
			 // value with double quotes
			"config": map[string]interface{}{
				"enabled": jsii.Boolean(true),
				"features": []*string{
					jsii.String("feature1"),
					jsii.String("feature2"),
				},
			},
		}, &JsonProcessingOptions{
			Escape: jsii.Boolean(true),
		}),
	},
	DestinationBucket: bucket,
})
```

## EFS Support

If your workflow needs more disk space than default (512 MB) disk space, you may attach an EFS storage to underlying
lambda function. To Enable EFS support set `efs` and `vpc` props for BucketDeployment.

Check sample usage below.
Please note that creating VPC inline may cause stack deletion failures. It is shown as below for simplicity.
To avoid such condition, keep your network infra (VPC) in a separate stack and pass as props.

```go
var destinationBucket bucket
var vpc vpc


s3deploy.NewBucketDeployment(this, jsii.String("DeployMeWithEfsStorage"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
	},
	DestinationBucket: DestinationBucket,
	DestinationKeyPrefix: jsii.String("efs/"),
	UseEfs: jsii.Boolean(true),
	Vpc: Vpc,
	RetainOnDelete: jsii.Boolean(false),
})
```

## Data with deploy-time values

The content passed to `Source.data()`, `Source.jsonData()`, or `Source.yamlData()` can include
references that will get resolved only during deployment. Only a subset of CloudFormation functions
are supported however, namely: Ref, Fn::GetAtt, Fn::Join, and Fn::Select (Fn::Split may be nested under Fn::Select).

For example:

```go
import sns "github.com/aws/aws-cdk-go/awscdk"
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

var destinationBucket bucket
var topic topic
var tg applicationTargetGroup


appConfig := map[string]interface{}{
	"topic_arn": topic.topicArn,
	"base_url": jsii.String("https://my-endpoint"),
	"lb_name": tg.firstLoadBalancerFullName,
}

s3deploy.NewBucketDeployment(this, jsii.String("BucketDeployment"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_JsonData(jsii.String("config.json"), appConfig),
	},
	DestinationBucket: DestinationBucket,
})
```

The value in `topic.topicArn` is a deploy-time value. It only gets resolved
during deployment by placing a marker in the generated source file and
substituting it when its deployed to the destination with the actual value.

### Substitutions from Templated Files

The `DeployTimeSubstitutedFile` construct allows you to specify substitutions
to make from placeholders in a local file which will be resolved during deployment. This
is especially useful in situations like creating an API from a spec file, where users might
want to reference other CDK resources they have created.

The syntax for template variables is `{{ variableName }}` in your local file. Then, you would
specify the substitutions in CDK like this:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var myLambdaFunction function
var destinationBucket bucket
//(Optional) if provided, the resulting processed file would be uploaded to the destinationBucket under the destinationKey name.
var destinationKey string
var role role


s3deploy.NewDeployTimeSubstitutedFile(this, jsii.String("MyFile"), &DeployTimeSubstitutedFileProps{
	Source: jsii.String("my-file.yaml"),
	DestinationKey: destinationKey,
	DestinationBucket: destinationBucket,
	Substitutions: map[string]*string{
		"variableName": myLambdaFunction.functionName,
	},
	Role: role,
})
```

Nested variables, like `{{ {{ foo }} }}` or `{{ foo {{ bar }} }}`, are not supported by this
construct. In the first case of a single variable being is double nested `{{ {{ foo }} }}`, only
the `{{ foo }}` would be replaced by the substitution, and the extra brackets would remain in the file.
In the second case of two nexted variables `{{ foo {{ bar }} }}`, only the `{{ bar }}` would be replaced
in the file.

## Keep Files Zipped

By default, files are zipped, then extracted into the destination bucket.

You can use the option `extract: false` to disable this behavior, in which case, files will remain in a zip file when deployed to S3. To reference the object keys, or filenames, which will be deployed to the bucket, you can use the `objectKeys` getter on the bucket deployment.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var destinationBucket bucket


myBucketDeployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployMeWithoutExtractingFilesOnDestination"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
	},
	DestinationBucket: DestinationBucket,
	Extract: jsii.Boolean(false),
})

cdk.NewCfnOutput(this, jsii.String("ObjectKey"), &CfnOutputProps{
	Value: cdk.Fn_Select(jsii.Number(0), myBucketDeployment.objectKeys),
})
```

## Controlling the Output of Source Object Keys

By default, the keys of the source objects copied to the destination bucket are returned in the Data property of the custom resource. However, you can disable this behavior by setting the outputObjectKeys property to false. This is particularly useful when the number of objects is too large and might exceed the size limit of the responseData property.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var destinationBucket bucket


myBucketDeployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployMeWithoutExtractingFilesOnDestination"), &BucketDeploymentProps{
	Sources: []iSource{
		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
	},
	DestinationBucket: DestinationBucket,
	OutputObjectKeys: jsii.Boolean(false),
})

cdk.NewCfnOutput(this, jsii.String("ObjectKey"), &CfnOutputProps{
	Value: cdk.Fn_Select(jsii.Number(0), myBucketDeployment.objectKeys),
})
```

## Notes

* This library uses an AWS CloudFormation custom resource which is about 10MiB in
  size. The code of this resource is bundled with this library.
* AWS Lambda execution time is limited to 15min. This limits the amount of data
  which can be deployed into the bucket by this timeout.
* When the `BucketDeployment` is removed from the stack, the contents are retained
  in the destination bucket ([#952](https://github.com/aws/aws-cdk/issues/952)).
* If you are using `s3deploy.Source.bucket()` to take the file source from
  another bucket: the deployed files will only be updated if the key (file name)
  of the file in the source  bucket changes. Mutating the file in place will not
  be good enough: the custom resource will simply not run if the properties don't
  change.

  * If you use assets (`s3deploy.Source.asset()`) you don't need to worry
    about this: the asset system will make sure that if the files have changed,
    the file name is unique and the deployment will run.

## Development

The custom resource is implemented in Python 3.9 in order to be able to leverage
the AWS CLI for "aws s3 sync".
The code is now in the `@aws-cdk/custom-resource-handlers` package under [`lib/aws-s3-deployment/bucket-deployment-handler`](https://github.com/aws/aws-cdk/tree/main/packages/@aws-cdk/custom-resource-handlers/lib/aws-s3-deployment/bucket-deployment-handler/) and
unit tests are under [`test/aws-s3-deployment/bucket-deployment-handler`](https://github.com/aws/aws-cdk/tree/main/packages/@aws-cdk/custom-resource-handlers/test/aws-s3-deployment/bucket-deployment-handler/).

This package requires Python 3.9 during build time in order to create the custom
resource Lambda bundle and test it. It also relies on a few bash scripts, so
might be tricky to build on Windows.

## Roadmap

* [ ] Support "blue/green" deployments ([#954](https://github.com/aws/aws-cdk/issues/954))

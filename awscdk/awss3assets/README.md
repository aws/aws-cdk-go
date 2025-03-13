# AWS CDK Assets

Assets are local files or directories which are needed by a CDK app. A common
example is a directory which contains the handler code for a Lambda function,
but assets can represent any artifact that is needed for the app's operation.

When deploying a CDK app that includes constructs with assets, the CDK toolkit
will first upload all the assets to S3, and only then deploy the stacks. The S3
locations of the uploaded assets will be passed in as CloudFormation Parameters
to the relevant stacks.

The following JavaScript example defines a directory asset which is archived as
a .zip file and uploaded to S3 during deployment.

```go
asset := assets.NewAsset(this, jsii.String("SampleAsset"), &AssetProps{
	Path: path.join(__dirname, jsii.String("sample-asset-directory")),
})
```

The following JavaScript example defines a file asset, which is uploaded as-is
to an S3 bucket during deployment.

```go
asset := assets.NewAsset(this, jsii.String("SampleAsset"), &AssetProps{
	Path: path.join(__dirname, jsii.String("file-asset.txt")),
})
```

## Attributes

`Asset` constructs expose the following deploy-time attributes:

* `s3BucketName` - the name of the assets S3 bucket.
* `s3ObjectKey` - the S3 object key of the asset file (whether it's a file or a zip archive)
* `s3ObjectUrl` - the S3 object URL of the asset (i.e. s3://amzn-s3-demo-bucket/mykey.zip)
* `httpUrl` - the S3 HTTP URL of the asset (i.e. https://s3.us-east-1.amazonaws.com/amzn-s3-demo-bucket/mykey.zip)

In the following example, the various asset attributes are exported as stack outputs:

```go
asset := assets.NewAsset(this, jsii.String("SampleAsset"), &AssetProps{
	Path: path.join(__dirname, jsii.String("sample-asset-directory")),
})

cdk.NewCfnOutput(this, jsii.String("S3BucketName"), &CfnOutputProps{
	Value: asset.S3BucketName,
})
cdk.NewCfnOutput(this, jsii.String("S3ObjectKey"), &CfnOutputProps{
	Value: asset.S3ObjectKey,
})
cdk.NewCfnOutput(this, jsii.String("S3HttpURL"), &CfnOutputProps{
	Value: asset.HttpUrl,
})
cdk.NewCfnOutput(this, jsii.String("S3ObjectURL"), &CfnOutputProps{
	Value: asset.S3ObjectUrl,
})
```

## Permissions

IAM roles, users or groups which need to be able to read assets in runtime will should be
granted IAM permissions. To do that use the `asset.grantRead(principal)` method:

The following example grants an IAM group read permissions on an asset:

```go
group := iam.NewGroup(this, jsii.String("MyUserGroup"))
asset.GrantRead(group)
```

## How does it work

When an asset is defined in a construct, a construct metadata entry
`aws:cdk:asset` is emitted with instructions on where to find the asset and what
type of packaging to perform (`zip` or `file`). Furthermore, the synthesized
CloudFormation template will also include two CloudFormation parameters: one for
the asset's bucket and one for the asset S3 key. Those parameters are used to
reference the deploy-time values of the asset (using `{ Ref: "Param" }`).

Then, when the stack is deployed, the toolkit will package the asset (i.e. zip
the directory), calculate an MD5 hash of the contents and will render an S3 key
for this asset within the toolkit's asset store. If the file doesn't exist in
the asset store, it is uploaded during deployment.

> The toolkit's asset store is an S3 bucket created by the toolkit for each
> environment the toolkit operates in (environment = account + region).

Now, when the toolkit deploys the stack, it will set the relevant CloudFormation
Parameters to point to the actual bucket and key for each asset.

## Asset Bundling

When defining an asset, you can use the `bundling` option to specify a command
to run inside a docker container. The command can read the contents of the asset
source from `/asset-input` and is expected to write files under `/asset-output`
(directories mapped inside the container). The files under `/asset-output` will
be zipped and uploaded to S3 as the asset.

The following example uses custom asset bundling to convert a markdown file to html:

```go
asset := assets.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
	Path: path.join(__dirname, jsii.String("markdown-asset")),
	 // /asset-input and working directory in the container
	Bundling: &BundlingOptions{
		Image: awscdk.DockerImage_FromBuild(path.join(__dirname, jsii.String("alpine-markdown"))),
		 // Build an image
		Command: []*string{
			jsii.String("sh"),
			jsii.String("-c"),
			jsii.String(`
			            markdown index.md > /asset-output/index.html
			          `),
		},
	},
})
```

The bundling docker image (`image`) can either come from a registry (`DockerImage.fromRegistry`)
or it can be built from a `Dockerfile` located inside your project (`DockerImage.fromBuild`).

You can set the `CDK_DOCKER` environment variable in order to provide a custom
docker program to execute. This may sometime be needed when building in
environments where the standard docker cannot be executed (see
https://github.com/aws/aws-cdk/issues/8460 for details).

Use `local` to specify a local bundling provider. The provider implements a
method `tryBundle()` which should return `true` if local bundling was performed.
If `false` is returned, docker bundling will be done:

```go
import "github.com/aws/aws-cdk-go/awscdk"


type myBundle struct {
}

func (this *myBundle) tryBundle(outputDir *string, options bundlingOptions) *bool {
	canRunLocally := true // replace with actual logic
	if canRunLocally {
		// perform local bundling here
		return jsii.Boolean(true)
	}
	return jsii.Boolean(false)
}

awscdk.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
	Path: jsii.String("/path/to/asset"),
	Bundling: &bundlingOptions{
		Local: NewMyBundle(),
		// Docker bundling fallback
		Image: cdk.DockerImage_FromRegistry(jsii.String("alpine")),
		Entrypoint: []*string{
			jsii.String("/bin/sh"),
			jsii.String("-c"),
		},
		Command: []*string{
			jsii.String("bundle"),
		},
	},
})
```

Although optional, it's recommended to provide a local bundling method which can
greatly improve performance.

If the bundling output contains a single archive file (zip or jar) it will be
uploaded to S3 as-is and will not be zipped. Otherwise the contents of the
output directory will be zipped and the zip file will be uploaded to S3. This
is the default behavior for `bundling.outputType` (`BundlingOutput.AUTO_DISCOVER`).

Use `BundlingOutput.NOT_ARCHIVED` if the bundling output must always be zipped:

```go
import "github.com/aws/aws-cdk-go/awscdk"


asset := awscdk.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
	Path: jsii.String("/path/to/asset"),
	Bundling: &BundlingOptions{
		Image: cdk.DockerImage_FromRegistry(jsii.String("alpine")),
		Command: []*string{
			jsii.String("command-that-produces-an-archive.sh"),
		},
		OutputType: cdk.BundlingOutput_NOT_ARCHIVED,
	},
})
```

Use `BundlingOutput.ARCHIVED` if the bundling output contains a single archive file and
you don't want it to be zipped.

### Docker options

Depending on your build environment, you may need to pass certain docker options to the `docker run` command that bundles assets.
This can be done using [BundlingOptions](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.BundlingOptions.html) properties.

Some optional properties to pass to the docker bundling

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"


asset := awscdk.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
	Path: jsii.String("/path/to/asset"),
	Bundling: &BundlingOptions{
		Image: lambda.Runtime_PYTHON_3_9().BundlingImage,
		Command: []*string{
			jsii.String("bash"),
			jsii.String("-c"),
			jsii.String("pip install -r requirements.txt -t /asset-output && cp -au . /asset-output"),
		},
		SecurityOpt: jsii.String("no-new-privileges:true"),
		 // https://docs.docker.com/engine/reference/commandline/run/#optional-security-options---security-opt
		Network: jsii.String("host"),
	},
})
```

## CloudFormation Resource Metadata

> NOTE: This section is relevant for authors of AWS Resource Constructs.

In certain situations, it is desirable for tools to be able to know that a certain CloudFormation
resource is using a local asset. For example, SAM CLI can be used to invoke AWS Lambda functions
locally for debugging purposes.

To enable such use cases, external tools will consult a set of metadata entries on AWS CloudFormation
resources:

* `aws:asset:path` points to the local path of the asset.
* `aws:asset:property` is the name of the resource property where the asset is used

Using these two metadata entries, tools will be able to identify that assets are used
by a certain resource, and enable advanced local experiences.

To add these metadata entries to a resource, use the
`asset.addResourceMetadata(resource, property)` method.

See https://github.com/aws/aws-cdk/issues/1432 for more details

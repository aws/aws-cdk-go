# Amazon ECR Construct Library

This package contains constructs for working with Amazon Elastic Container Registry.

## Repositories

Define a repository by creating a new instance of `Repository`. A repository
holds multiple versions of a single container image.

```go
repository := ecr.NewRepository(this, jsii.String("Repository"))
```

## Image scanning

Amazon ECR image scanning helps in identifying software vulnerabilities in your container images.
You can manually scan container images stored in Amazon ECR, or you can configure your repositories
to scan images when you push them to a repository. To create a new repository to scan on push, simply
enable `imageScanOnPush` in the properties.

```go
repository := ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
	ImageScanOnPush: jsii.Boolean(true),
})
```

To create an `onImageScanCompleted` event rule and trigger the event target

```go
var repository repository
var target someTarget


repository.OnImageScanCompleted(jsii.String("ImageScanComplete")).AddTarget(target)
```

### Authorization Token

Besides the Amazon ECR APIs, ECR also allows the Docker CLI or a language-specific Docker library to push and pull
images from an ECR repository. However, the Docker CLI does not support native IAM authentication methods and
additional steps must be taken so that Amazon ECR can authenticate and authorize Docker push and pull requests.
More information can be found at at [Registry Authentication](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth).

A Docker authorization token can be obtained using the `GetAuthorizationToken` ECR API. The following code snippets
grants an IAM user access to call this API.

```go
user := iam.NewUser(this, jsii.String("User"))
ecr.AuthorizationToken_GrantRead(user)
```

If you access images in the [Public ECR Gallery](https://gallery.ecr.aws/) as well, it is recommended you authenticate to the registry to benefit from
higher rate and bandwidth limits.

> See `Pricing` in https://aws.amazon.com/blogs/aws/amazon-ecr-public-a-new-public-container-registry/ and [Service quotas](https://docs.aws.amazon.com/AmazonECR/latest/public/public-service-quotas.html).

The following code snippet grants an IAM user access to retrieve an authorization token for the public gallery.

```go
user := iam.NewUser(this, jsii.String("User"))
ecr.PublicGalleryAuthorizationToken_GrantRead(user)
```

This user can then proceed to login to the registry using one of the [authentication methods](https://docs.aws.amazon.com/AmazonECR/latest/public/public-registries.html#public-registry-auth).

### Other Grantee

#### grantPush

The grantPush method grants the specified IAM entity (the grantee) permission to push images to the ECR repository. Specifically, it grants permissions for the following actions:

* 'ecr:CompleteLayerUpload'
* 'ecr:UploadLayerPart'
* 'ecr:InitiateLayerUpload'
* 'ecr:BatchCheckLayerAvailability'
* 'ecr:PutImage'
* 'ecr:GetAuthorizationToken'

Here is an example of granting a user push permissions:

```go
var repository repository


role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com")),
})
repository.GrantPush(role)
```

#### grantPull

The grantPull method grants the specified IAM entity (the grantee) permission to pull images from the ECR repository. Specifically, it grants permissions for the following actions:

* 'ecr:BatchCheckLayerAvailability'
* 'ecr:GetDownloadUrlForLayer'
* 'ecr:BatchGetImage'
* 'ecr:GetAuthorizationToken'

```go
var repository repository


role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com")),
})
repository.GrantPull(role)
```

#### grantPullPush

The grantPullPush method grants the specified IAM entity (the grantee) permission to both pull and push images from/to the ECR repository. Specifically, it grants permissions for all the actions required for pull and push permissions.

Here is an example of granting a user both pull and push permissions:

```go
var repository repository


role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com")),
})
repository.GrantPullPush(role)
```

By using these methods, you can grant specific operational permissions on the ECR repository to IAM entities. This allows for proper management of access to the repository and ensures security.

### Image tag immutability

You can set tag immutability on images in our repository using the `imageTagMutability` construct prop.

```go
ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
	ImageTagMutability: ecr.TagMutability_IMMUTABLE,
})
```

### Encryption

By default, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. For more control over the encryption for your Amazon ECR repositories, you can use server-side encryption with KMS keys stored in AWS Key Management Service (AWS KMS). Read more about this feature in the [ECR Developer Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html).

When you use AWS KMS to encrypt your data, you can either use the default AWS managed key, which is managed by Amazon ECR, by specifying `RepositoryEncryption.KMS` in the `encryption` property. Or specify your own customer managed KMS key, by specifying the `encryptionKey` property.

When `encryptionKey` is set, the `encryption` property must be `KMS` or empty.

In the case `encryption` is set to `KMS` but no `encryptionKey` is set, an AWS managed KMS key is used.

```go
ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
	Encryption: ecr.RepositoryEncryption_KMS(),
})
```

Otherwise, a customer-managed KMS key is used if `encryptionKey` was set and `encryption` was optionally set to `KMS`.

```go
import kms "github.com/aws/aws-cdk-go/awscdk"


ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
	EncryptionKey: kms.NewKey(this, jsii.String("Key")),
})
```

## Automatically clean up repositories

You can set life cycle rules to automatically clean up old images from your
repository. The first life cycle rule that matches an image will be applied
against that image. For example, the following deletes images older than
30 days, while keeping all images tagged with prod (note that the order
is important here):

```go
var repository repository

repository.AddLifecycleRule(&LifecycleRule{
	TagPrefixList: []*string{
		jsii.String("prod"),
	},
	MaxImageCount: jsii.Number(9999),
})
repository.AddLifecycleRule(&LifecycleRule{
	MaxImageAge: awscdk.Duration_Days(jsii.Number(30)),
})
```

When using `tagPatternList`, an image is successfully matched if it matches
the wildcard filter.

```go
var repository repository

repository.AddLifecycleRule(&LifecycleRule{
	TagPatternList: []*string{
		jsii.String("prod*"),
	},
	MaxImageCount: jsii.Number(9999),
})
```

### Repository deletion

When a repository is removed from a stack (or the stack is deleted), the ECR
repository will be removed according to its removal policy (which by default will
simply orphan the repository and leave it in your AWS account). If the removal
policy is set to `RemovalPolicy.DESTROY`, the repository will be deleted as long
as it does not contain any images.

To override this and force all images to get deleted during repository deletion,
enable the `emptyOnDelete` option as well as setting the removal policy to
`RemovalPolicy.DESTROY`.

```go
repository := ecr.NewRepository(this, jsii.String("MyTempRepo"), &RepositoryProps{
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	EmptyOnDelete: jsii.Boolean(true),
})
```

## Managing the Resource Policy

You can add statements to the resource policy of the repository using the
`addToResourcePolicy` method. However, be advised that you must not include
a `resources` section in the `PolicyStatement`.

```go
var repository repository

repository.AddToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("ecr:GetDownloadUrlForLayer"),
	},
	// resources: ['*'], // not currently allowed!
	Principals: []iPrincipal{
		iam.NewAnyPrincipal(),
	},
}))
```

## import existing repository

You can import an existing repository into your CDK app using the `Repository.fromRepositoryArn`, `Repository.fromRepositoryName` or `Repository.fromLookup` method.
These methods take the ARN or the name of the repository and returns an `IRepository` object.

```go
// import using repository name
repositoryFromName := ecr.Repository_FromRepositoryName(this, jsii.String("ImportedRepoByName"), jsii.String("my-repo-name"))

// import using repository ARN
repositoryFromArn := ecr.Repository_FromRepositoryArn(this, jsii.String("ImportedRepoByArn"), jsii.String("arn:aws:ecr:us-east-1:123456789012:repository/my-repo-name"))

// import using repository lookup
// You have to provide either `repositoryArn` or `repositoryName` to lookup the repository
repositoryFromLookup := ecr.Repository_FromLookup(this, jsii.String("ImportedRepoByLookup"), &RepositoryLookupOptions{
	RepositoryArn: jsii.String("arn:aws:ecr:us-east-1:123456789012:repository/my-repo-name"),
	RepositoryName: jsii.String("my-repo-name"),
})
```

## CloudWatch event rules

You can publish repository events to a CloudWatch event rule with `onEvent`:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


repo := ecr.NewRepository(this, jsii.String("Repo"))
lambdaHandler := lambda.NewFunction(this, jsii.String("LambdaFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Code: lambda.Code_FromInline(jsii.String("# dummy func")),
	Handler: jsii.String("index.handler"),
})

repo.OnEvent(jsii.String("OnEventTargetLambda"), &OnEventOptions{
	Target: awscdk.NewLambdaFunction(lambdaHandler),
})
```

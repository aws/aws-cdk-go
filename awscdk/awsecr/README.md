# Amazon ECR Construct Library

This package contains constructs for working with Amazon Elastic Container Registry.

## Repositories

Define a repository by creating a new instance of `Repository`. A repository
holds multiple versions of a single container image.

```go
repository := ecr.NewRepository(this, jsii.String("Repository"))
```

## Image scanning

Amazon ECR image scanning helps in identifying software vulnerabilities in your container images. You can manually scan container images stored in Amazon ECR, or you can configure your repositories to scan images when you push them to a repository. To create a new repository to scan on push, simply enable `imageScanOnPush` in the properties

```go
repository := ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
	imageScanOnPush: jsii.Boolean(true),
})
```

To create an `onImageScanCompleted` event rule and trigger the event target

```go
var repository repository
var target someTarget


repository.onImageScanCompleted(jsii.String("ImageScanComplete")).addTarget(target)
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
ecr.authorizationToken.grantRead(user)
```

If you access images in the [Public ECR Gallery](https://gallery.ecr.aws/) as well, it is recommended you authenticate to the registry to benefit from
higher rate and bandwidth limits.

> See `Pricing` in https://aws.amazon.com/blogs/aws/amazon-ecr-public-a-new-public-container-registry/ and [Service quotas](https://docs.aws.amazon.com/AmazonECR/latest/public/public-service-quotas.html).

The following code snippet grants an IAM user access to retrieve an authorization token for the public gallery.

```go
user := iam.NewUser(this, jsii.String("User"))
ecr.publicGalleryAuthorizationToken.grantRead(user)
```

This user can then proceed to login to the registry using one of the [authentication methods](https://docs.aws.amazon.com/AmazonECR/latest/public/public-registries.html#public-registry-auth).

### Image tag immutability

You can set tag immutability on images in our repository using the `imageTagMutability` construct prop.

```go
ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
	imageTagMutability: ecr.tagMutability_IMMUTABLE,
})
```

### Encryption

By default, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. For more control over the encryption for your Amazon ECR repositories, you can use server-side encryption with KMS keys stored in AWS Key Management Service (AWS KMS). Read more about this feature in the [ECR Developer Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html).

When you use AWS KMS to encrypt your data, you can either use the default AWS managed key, which is managed by Amazon ECR, by specifying `RepositoryEncryption.KMS` in the `encryption` property. Or specify your own customer managed KMS key, by specifying the `encryptionKey` property.

When `encryptionKey` is set, the `encryption` property must be `KMS` or empty.

In the case `encryption` is set to `KMS` but no `encryptionKey` is set, an AWS managed KMS key is used.

```go
ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
	encryption: ecr.repositoryEncryption_KMS(),
})
```

Otherwise, a customer-managed KMS key is used if `encryptionKey` was set and `encryption` was optionally set to `KMS`.

```go
import kms "github.com/aws/aws-cdk-go/awscdk"


ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
	encryptionKey: kms.NewKey(this, jsii.String("Key")),
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

repository.addLifecycleRule(&lifecycleRule{
	tagPrefixList: []*string{
		jsii.String("prod"),
	},
	maxImageCount: jsii.Number(9999),
})
repository.addLifecycleRule(&lifecycleRule{
	maxImageAge: awscdk.Duration.days(jsii.Number(30)),
})
```

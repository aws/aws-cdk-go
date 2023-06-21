# AWS Cloud9 Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

AWS Cloud9 is a cloud-based integrated development environment (IDE) that lets you write, run, and debug your code with just a
browser. It includes a code editor, debugger, and terminal. Cloud9 comes prepackaged with essential tools for popular
programming languages, including JavaScript, Python, PHP, and more, so you don’t need to install files or configure your
development machine to start new projects. Since your Cloud9 IDE is cloud-based, you can work on your projects from your
office, home, or anywhere using an internet-connected machine. Cloud9 also provides a seamless experience for developing
serverless applications enabling you to easily define resources, debug, and switch between local and remote execution of
serverless applications. With Cloud9, you can quickly share your development environment with your team, enabling you to pair
program and track each other's inputs in real time.

## Creating EC2 Environment

EC2 Environments are defined with `Ec2Environment`. To create an EC2 environment in the private subnet, specify
`subnetSelection` with private `subnetType`.

```go
// create a cloud9 ec2 environment in a new VPC
vpc := ec2.NewVpc(this, jsii.String("VPC"), &VpcProps{
	MaxAzs: jsii.Number(3),
})
cloud9.NewEc2Environment(this, jsii.String("Cloud9Env"), &Ec2EnvironmentProps{
	Vpc: Vpc,
	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
})

// or create the cloud9 environment in the default VPC with specific instanceType
defaultVpc := ec2.Vpc_FromLookup(this, jsii.String("DefaultVPC"), &VpcLookupOptions{
	IsDefault: jsii.Boolean(true),
})
cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &Ec2EnvironmentProps{
	Vpc: defaultVpc,
	InstanceType: ec2.NewInstanceType(jsii.String("t3.large")),
	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
})

// or specify in a different subnetSelection
c9env := cloud9.NewEc2Environment(this, jsii.String("Cloud9Env3"), &Ec2EnvironmentProps{
	Vpc: Vpc,
	SubnetSelection: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
})

// print the Cloud9 IDE URL in the output
// print the Cloud9 IDE URL in the output
awscdk.NewCfnOutput(this, jsii.String("URL"), &CfnOutputProps{
	Value: c9env.IdeUrl,
})
```

## Specifying EC2 AMI

Use `imageId` to specify the EC2 AMI image to be used:

```go
defaultVpc := ec2.Vpc_FromLookup(this, jsii.String("DefaultVPC"), &VpcLookupOptions{
	IsDefault: jsii.Boolean(true),
})
cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &Ec2EnvironmentProps{
	Vpc: defaultVpc,
	InstanceType: ec2.NewInstanceType(jsii.String("t3.large")),
	ImageId: cloud9.ImageId_UBUNTU_18_04,
})
```

## Cloning Repositories

Use `clonedRepositories` to clone one or multiple AWS Codecommit repositories into the environment:

```go
import "github.com/aws/aws-cdk-go/awscdk"

// create a new Cloud9 environment and clone the two repositories
var vpc vpc


// create a codecommit repository to clone into the cloud9 environment
repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &RepositoryProps{
	RepositoryName: jsii.String("new-repo"),
})

// import an existing codecommit repository to clone into the cloud9 environment
repoExisting := codecommit.Repository_FromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
cloud9.NewEc2Environment(this, jsii.String("C9Env"), &Ec2EnvironmentProps{
	Vpc: Vpc,
	ClonedRepositories: []cloneRepository{
		cloud9.*cloneRepository_FromCodeCommit(repoNew, jsii.String("/src/new-repo")),
		cloud9.*cloneRepository_*FromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
	},
	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
})
```

## Specifying Owners

Every Cloud9 Environment has an **owner**. An owner has full control over the environment, and can invite additional members to the environment for collaboration purposes. For more information, see [Working with shared environments in AWS Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/share-environment.html)).

By default, the owner will be the identity that creates the Environment, which is most likely your CloudFormation Execution Role when the Environment is created using CloudFormation. Provider a value for the `owner` property to assign a different owner, either a specific IAM User or the AWS Account Root User.

`Owner` is a user that owns a Cloud9 environment . `Owner` has their own access permissions, resources. And we can specify an `Owner`in an Ec2 environment which could be of two types, 1. AccountRoot and 2. Iam User. It allows AWS to determine who has permissions to manage the environment, either an IAM user or the account root user (but using the account root user is not recommended, see [environment sharing best practices](https://docs.aws.amazon.com/cloud9/latest/user-guide/share-environment.html#share-environment-best-practices)).

To specify the AWS Account Root User as the environment owner, use `Owner.accountRoot()`

```go
var vpc vpc

cloud9.NewEc2Environment(this, jsii.String("C9Env"), &Ec2EnvironmentProps{
	Vpc: Vpc,
	ImageId: cloud9.ImageId_AMAZON_LINUX_2,

	Owner: cloud9.Owner_AccountRoot(jsii.String("111111111")),
})
```

To specify a specific IAM User as the environment owner, use `Owner.user()`. The user should have the `AWSCloud9Administrator` managed policy

```go
import "github.com/aws/aws-cdk-go/awscdk"
var vpc vpc


user := iam.NewUser(this, jsii.String("user"))
user.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AWSCloud9Administrator")))
cloud9.NewEc2Environment(this, jsii.String("C9Env"), &Ec2EnvironmentProps{
	Vpc: Vpc,
	ImageId: cloud9.ImageId_AMAZON_LINUX_2,

	Owner: cloud9.Owner_User(user),
})
```

## Auto-Hibernation

A Cloud9 environemnt can automatically start and stop the associated EC2 instance to reduce costs.

Use `automaticStop` to specify the number of minutes until the running instance is shut down after the environment was last used.

```go
defaultVpc := ec2.Vpc_FromLookup(this, jsii.String("DefaultVPC"), &VpcLookupOptions{
	IsDefault: jsii.Boolean(true),
})
cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &Ec2EnvironmentProps{
	Vpc: defaultVpc,
	ImageId: cloud9.ImageId_AMAZON_LINUX_2,
	AutomaticStop: awscdk.Duration_Minutes(jsii.Number(30)),
})
```

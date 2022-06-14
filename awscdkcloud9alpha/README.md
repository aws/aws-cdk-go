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
vpc := ec2.NewVpc(this, jsii.String("VPC"), &vpcProps{
	maxAzs: jsii.Number(3),
})
cloud9.NewEc2Environment(this, jsii.String("Cloud9Env"), &ec2EnvironmentProps{
	vpc: vpc,
})

// or create the cloud9 environment in the default VPC with specific instanceType
defaultVpc := ec2.vpc.fromLookup(this, jsii.String("DefaultVPC"), &vpcLookupOptions{
	isDefault: jsii.Boolean(true),
})
cloud9.NewEc2Environment(this, jsii.String("Cloud9Env2"), &ec2EnvironmentProps{
	vpc: defaultVpc,
	instanceType: ec2.NewInstanceType(jsii.String("t3.large")),
})

// or specify in a different subnetSelection
c9env := cloud9.NewEc2Environment(this, jsii.String("Cloud9Env3"), &ec2EnvironmentProps{
	vpc: vpc,
	subnetSelection: &subnetSelection{
		subnetType: ec2.subnetType_PRIVATE_WITH_NAT,
	},
})

// print the Cloud9 IDE URL in the output
// print the Cloud9 IDE URL in the output
awscdk.NewCfnOutput(this, jsii.String("URL"), &cfnOutputProps{
	value: c9env.ideUrl,
})
```

## Cloning Repositories

Use `clonedRepositories` to clone one or multiple AWS Codecommit repositories into the environment:

```go
import codecommit "github.com/aws/aws-cdk-go/awscdk"

// create a new Cloud9 environment and clone the two repositories
var vpc vpc


// create a codecommit repository to clone into the cloud9 environment
repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &repositoryProps{
	repositoryName: jsii.String("new-repo"),
})

// import an existing codecommit repository to clone into the cloud9 environment
repoExisting := codecommit.repository.fromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
	vpc: vpc,
	clonedRepositories: []cloneRepository{
		cloud9.*cloneRepository.fromCodeCommit(repoNew, jsii.String("/src/new-repo")),
		cloud9.*cloneRepository.fromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
	},
})
```

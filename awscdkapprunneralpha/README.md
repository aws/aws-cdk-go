# AWS::AppRunner Construct Library

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

```go
import apprunner "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
```

## Introduction

AWS App Runner is a fully managed service that makes it easy for developers to quickly deploy containerized web applications and APIs, at scale and with no prior infrastructure experience required. Start with your source code or a container image. App Runner automatically builds and deploys the web application and load balances traffic with encryption. App Runner also scales up or down automatically to meet your traffic needs. With App Runner, rather than thinking about servers or scaling, you have more time to focus on your applications.

## Service

The `Service` construct allows you to create AWS App Runner services with `ECR Public`, `ECR` or `Github` with the `source` property in the following scenarios:

* `Source.fromEcr()` - To define the source repository from `ECR`.
* `Source.fromEcrPublic()` - To define the source repository from `ECR Public`.
* `Source.fromGitHub()` - To define the source repository from the `Github repository`.
* `Source.fromAsset()` - To define the source from local asset directory.

## ECR Public

To create a `Service` with ECR Public:

```go
apprunner.NewService(this, jsii.String("Service"), &serviceProps{
	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
		imageConfiguration: &imageConfiguration{
			port: jsii.Number(8000),
		},
		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
	}),
})
```

## ECR

To create a `Service` from an existing ECR repository:

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"


apprunner.NewService(this, jsii.String("Service"), &serviceProps{
	source: apprunner.source.fromEcr(&ecrProps{
		imageConfiguration: &imageConfiguration{
			port: jsii.Number(80),
		},
		repository: ecr.repository.fromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
		tagOrDigest: jsii.String("latest"),
	}),
})
```

To create a `Service` from local docker image asset directory  built and pushed to Amazon ECR:

```go
import assets "github.com/aws/aws-cdk-go/awscdk"


imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &dockerImageAssetProps{
	directory: path.join(__dirname, jsii.String("./docker.assets")),
})
apprunner.NewService(this, jsii.String("Service"), &serviceProps{
	source: apprunner.source.fromAsset(&assetProps{
		imageConfiguration: &imageConfiguration{
			port: jsii.Number(8000),
		},
		asset: imageAsset,
	}),
})
```

## GitHub

To create a `Service` from the GitHub repository, you need to specify an existing App Runner `Connection`.

See [Managing App Runner connections](https://docs.aws.amazon.com/apprunner/latest/dg/manage-connections.html) for more details.

```go
apprunner.NewService(this, jsii.String("Service"), &serviceProps{
	source: apprunner.source.fromGitHub(&githubRepositoryProps{
		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
		branch: jsii.String("main"),
		configurationSource: apprunner.configurationSourceType_REPOSITORY,
		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
	}),
})
```

Use `codeConfigurationValues` to override configuration values with the `API` configuration source type.

```go
apprunner.NewService(this, jsii.String("Service"), &serviceProps{
	source: apprunner.source.fromGitHub(&githubRepositoryProps{
		repositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
		branch: jsii.String("main"),
		configurationSource: apprunner.configurationSourceType_API,
		codeConfigurationValues: &codeConfigurationValues{
			runtime: apprunner.runtime_PYTHON_3(),
			port: jsii.String("8000"),
			startCommand: jsii.String("python app.py"),
			buildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
		},
		connection: apprunner.gitHubConnection.fromConnectionArn(jsii.String("CONNECTION_ARN")),
	}),
})
```

## IAM Roles

You are allowed to define `instanceRole` and `accessRole` for the `Service`.

`instanceRole` - The IAM role that provides permissions to your App Runner service. These are permissions that
your code needs when it calls any AWS APIs.

`accessRole` - The IAM role that grants the App Runner service access to a source repository. It's required for
ECR image repositories (but not for ECR Public repositories). If not defined, a new access role will be generated
when required.

See [App Runner IAM Roles](https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles) for more details.

## VPC Connector

To associate an App Runner service with a custom VPC, define `vpcConnector` for the service.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
	ipAddresses: ec2.ipAddresses.cidr(jsii.String("10.0.0.0/16")),
})

vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &vpcConnectorProps{
	vpc: vpc,
	vpcSubnets: vpc.selectSubnets(&subnetSelection{
		subnetType: ec2.subnetType_PUBLIC,
	}),
	vpcConnectorName: jsii.String("MyVpcConnector"),
})

apprunner.NewService(this, jsii.String("Service"), &serviceProps{
	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
		imageConfiguration: &imageConfiguration{
			port: jsii.Number(8000),
		},
		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
	}),
	vpcConnector: vpcConnector,
})
```

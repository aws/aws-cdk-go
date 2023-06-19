# AWS::AppRunner Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

```go
import apprunner "github.com/aws/aws-cdk-go/awscdk"
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
apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(8000),
		},
		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
	}),
})
```

## ECR

To create a `Service` from an existing ECR repository:

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"


apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromEcr(&EcrProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(80),
		},
		Repository: ecr.Repository_FromRepositoryName(this, jsii.String("NginxRepository"), jsii.String("nginx")),
		TagOrDigest: jsii.String("latest"),
	}),
})
```

To create a `Service` from local docker image asset directory  built and pushed to Amazon ECR:

```go
import assets "github.com/aws/aws-cdk-go/awscdk"


imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &DockerImageAssetProps{
	Directory: path.join(__dirname, jsii.String("./docker.assets")),
})
apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromAsset(&AssetProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(8000),
		},
		Asset: imageAsset,
	}),
})
```

## GitHub

To create a `Service` from the GitHub repository, you need to specify an existing App Runner `Connection`.

See [Managing App Runner connections](https://docs.aws.amazon.com/apprunner/latest/dg/manage-connections.html) for more details.

```go
apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromGitHub(&GithubRepositoryProps{
		RepositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
		Branch: jsii.String("main"),
		ConfigurationSource: apprunner.ConfigurationSourceType_REPOSITORY,
		Connection: apprunner.GitHubConnection_FromConnectionArn(jsii.String("CONNECTION_ARN")),
	}),
})
```

Use `codeConfigurationValues` to override configuration values with the `API` configuration source type.

```go
apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromGitHub(&GithubRepositoryProps{
		RepositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
		Branch: jsii.String("main"),
		ConfigurationSource: apprunner.ConfigurationSourceType_API,
		CodeConfigurationValues: &CodeConfigurationValues{
			Runtime: apprunner.Runtime_PYTHON_3(),
			Port: jsii.String("8000"),
			StartCommand: jsii.String("python app.py"),
			BuildCommand: jsii.String("yum install -y pycairo && pip install -r requirements.txt"),
		},
		Connection: apprunner.GitHubConnection_FromConnectionArn(jsii.String("CONNECTION_ARN")),
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
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	Cidr: jsii.String("10.0.0.0/16"),
})

vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &VpcConnectorProps{
	Vpc: Vpc,
	VpcSubnets: vpc.selectSubnets(&SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	}),
	VpcConnectorName: jsii.String("MyVpcConnector"),
})

apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(8000),
		},
		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
	}),
	VpcConnector: VpcConnector,
})
```

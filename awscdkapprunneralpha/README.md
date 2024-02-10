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

The `Service` construct implements `IGrantable`.

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

To create a `Service` from local docker image asset directory built and pushed to Amazon ECR:

You can specify whether to enable continuous integration from the source repository with the `autoDeploymentsEnabled` flag.

```go
import assets "github.com/aws/aws-cdk-go/awscdk"


imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &DockerImageAssetProps{
	Directory: path.join(__dirname, jsii.String("docker.assets")),
})
apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromAsset(&AssetProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(8000),
		},
		Asset: imageAsset,
	}),
	AutoDeploymentsEnabled: jsii.Boolean(true),
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
your code needs when it calls any AWS APIs. If not defined, a new instance role will be generated
when required.

To add IAM policy statements to this role, use `addToRolePolicy()`:

```go
import "github.com/aws/aws-cdk-go/awscdk"


service := apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(8000),
		},
		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
	}),
})

service.AddToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Actions: []*string{
		jsii.String("s3:GetObject"),
	},
	Resources: []*string{
		jsii.String("*"),
	},
}))
```

`accessRole` - The IAM role that grants the App Runner service access to a source repository. It's required for
ECR image repositories (but not for ECR Public repositories). If not defined, a new access role will be generated
when required.

See [App Runner IAM Roles](https://docs.aws.amazon.com/apprunner/latest/dg/security_iam_service-with-iam.html#security_iam_service-with-iam-roles) for more details.

## VPC Connector

To associate an App Runner service with a custom VPC, define `vpcConnector` for the service.

```go
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	IpAddresses: ec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
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

## Secrets Manager

To include environment variables integrated with AWS Secrets Manager, use the `environmentSecrets` attribute.
You can use the `addSecret` method from the App Runner `Service` class to include secrets from outside the
service definition.

```go
import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
import ssm "github.com/aws/aws-cdk-go/awscdk"

var stack stack


secret := secretsmanager.NewSecret(stack, jsii.String("Secret"))
parameter := ssm.StringParameter_FromSecureStringParameterAttributes(stack, jsii.String("Parameter"), &SecureStringParameterAttributes{
	ParameterName: jsii.String("/name"),
	Version: jsii.Number(1),
})

service := apprunner.NewService(stack, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(8000),
			EnvironmentSecrets: map[string]secret{
				"SECRET": apprunner.*secret_fromSecretsManager(secret),
				"PARAMETER": apprunner.*secret_fromSsmParameter(parameter),
				"SECRET_ID": apprunner.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
					"versionId": jsii.String("version-id"),
				}),
				"SECRET_STAGE": apprunner.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
					"versionStage": jsii.String("version-stage"),
				}),
			},
		},
		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
	}),
})

service.AddSecret(jsii.String("LATER_SECRET"), apprunner.secret_FromSecretsManager(secret, jsii.String("field")))
```

## HealthCheck

To configure the health check for the service, use the `healthCheck` attribute.

You can specify it by static methods `HealthCheck.http` or `HealthCheck.tcp`.

```go
apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
		ImageConfiguration: &ImageConfiguration{
			Port: jsii.Number(8000),
		},
		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
	}),
	HealthCheck: apprunner.HealthCheck_Http(&HttpHealthCheckOptions{
		HealthyThreshold: jsii.Number(5),
		Interval: awscdk.Duration_Seconds(jsii.Number(10)),
		Path: jsii.String("/"),
		Timeout: awscdk.Duration_*Seconds(jsii.Number(10)),
		UnhealthyThreshold: jsii.Number(10),
	}),
})
```

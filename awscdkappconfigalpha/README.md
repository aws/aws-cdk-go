# AWS AppConfig Construct Library

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

Use AWS AppConfig, a capability of AWS Systems Manager, to create, manage, and quickly deploy application configurations. A configuration is a collection of settings that influence the behavior of your application. You can use AWS AppConfig with applications hosted on Amazon Elastic Compute Cloud (Amazon EC2) instances, AWS Lambda, containers, mobile applications, or IoT devices. To view examples of the types of configurations you can manage by using AWS AppConfig, see [Example configurations](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html#appconfig-creating-configuration-and-profile-examples).

## Application

In AWS AppConfig, an application is simply an organizational construct like a folder. This organizational construct has a
relationship with some unit of executable code. For example, you could create an application called MyMobileApp to organize and
manage configuration data for a mobile application installed by your users. Configurations and environments are associated with
the application.

The name and description of an application are optional.

Create a simple application:

```go
appconfig.NewApplication(this, jsii.String("MyApplication"))
```

Create an application with a name and description:

```go
appconfig.NewApplication(this, jsii.String("MyApplication"), &ApplicationProps{
	Name: jsii.String("App1"),
	Description: jsii.String("This is my application created through CDK."),
})
```

## Deployment Strategy

A deployment strategy defines how a configuration will roll out. The roll out is defined by four parameters: deployment type,
step percentage, deployment time, and bake time.
See: https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html

Deployment strategy with predefined values:

```go
appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
	RolloutStrategy: appconfig.RolloutStrategy_CANARY_10_PERCENT_20_MINUTES(),
})
```

Deployment strategy with custom values:

```go
appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
	RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
		GrowthFactor: jsii.Number(20),
		DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
		FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(30)),
	}),
})
```

## Configuration

A configuration is a higher-level construct that can either be a `HostedConfiguration` (stored internally through AWS
AppConfig) or a `SourcedConfiguration` (stored in an Amazon S3 bucket, AWS Secrets Manager secrets, Systems Manager (SSM)
Parameter Store parameters, SSM documents, or AWS CodePipeline). This construct manages deployments on creation.

### HostedConfiguration

A hosted configuration represents configuration stored in the AWS AppConfig hosted configuration store. A hosted configuration
takes in the configuration content and associated AWS AppConfig application. On construction of a hosted configuration, the
configuration is deployed.

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
})
```

You can define hosted configuration content using any of the following ConfigurationContent methods:

* `fromFile` - Defines the hosted configuration content from a file.
* `fromInlineText` - Defines the hosted configuration from inline text.
* `fromInlineJson` - Defines the hosted configuration from inline JSON.
* `fromInlineYaml` - Defines the hosted configuration from inline YAML.
* `fromInline` - Defines the hosted configuration from user-specified content types.

AWS AppConfig supports the following types of configuration profiles.

* **Feature flag**: Use a feature flag configuration to turn on new features that require a timely deployment, such as a product launch or announcement.
* **Freeform**: Use a freeform configuration to carefully introduce changes to your application.

A hosted configuration with type:

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
})
```

When you create a configuration and configuration profile, you can specify up to two validators. A validator ensures that your
configuration data is syntactically and semantically correct. You can create validators in either JSON Schema or as an AWS
Lambda function.
See [About validators](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html#appconfig-creating-configuration-and-profile-validators) for more information.

A hosted configuration with validators:

```go
var application application
var fn function


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
	Validators: []iValidator{
		appconfig.JsonSchemaValidator_FromFile(jsii.String("schema.json")),
		appconfig.LambdaValidator_FromFunction(fn),
	},
})
```

You can attach a deployment strategy (as described in the previous section) to your configuration to specify how you want your
configuration to roll out.

A hosted configuration with a deployment strategy:

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
	DeploymentStrategy: appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
		RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
			GrowthFactor: jsii.Number(15),
			DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
			FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(15)),
		}),
	}),
})
```

The `deployTo` parameter is used to specify which environments to deploy the configuration to. If this parameter is not
specified, there will not be a deployment.

A hosted configuration with `deployTo`:

```go
var application application
var env environment


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
	DeployTo: []iEnvironment{
		env,
	},
})
```

### SourcedConfiguration

A sourced configuration represents configuration stored in an Amazon S3 bucket, AWS Secrets Manager secret, Systems Manager
(SSM) Parameter Store parameter, SSM document, or AWS CodePipeline. A sourced configuration takes in the location source
construct and optionally a version number to deploy. On construction of a sourced configuration, the configuration is deployed
only if a version number is specified.

### S3

Use an Amazon S3 bucket to store a configuration.

```go
var application application


bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	Versioned: jsii.Boolean(true),
})

appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
})
```

Use an encrypted bucket:

```go
var application application


bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
	Versioned: jsii.Boolean(true),
	Encryption: s3.BucketEncryption_KMS,
})

appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
})
```

### AWS Secrets Manager secret

Use a Secrets Manager secret to store a configuration.

```go
var application application
var secret secret


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromSecret(secret),
})
```

### SSM Parameter Store parameter

Use an SSM parameter to store a configuration.

```go
var application application
var parameter stringParameter


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromParameter(parameter),
	VersionNumber: jsii.String("1"),
})
```

### SSM document

Use an SSM document to store a configuration.

```go
var application application
var document cfnDocument


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromCfnDocument(document),
})
```

### AWS CodePipeline

Use an AWS CodePipeline pipeline to store a configuration.

```go
var application application
var pipeline pipeline


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromPipeline(pipeline),
})
```

Similar to a hosted configuration, a sourced configuration can optionally take in a type, validators, a `deployTo` parameter, and a deployment strategy.

A sourced configuration with type:

```go
var application application
var bucket bucket


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
	Name: jsii.String("MyConfig"),
	Description: jsii.String("This is my sourced configuration from CDK."),
})
```

A sourced configuration with validators:

```go
var application application
var bucket bucket
var fn function


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
	Validators: []iValidator{
		appconfig.JsonSchemaValidator_FromFile(jsii.String("schema.json")),
		appconfig.LambdaValidator_FromFunction(fn),
	},
})
```

A sourced configuration with a deployment strategy:

```go
var application application
var bucket bucket


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
	DeploymentStrategy: appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
		RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
			GrowthFactor: jsii.Number(15),
			DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
			FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(15)),
		}),
	}),
})
```

The `deployTo` parameter is used to specify which environments to deploy the configuration to. If this parameter is not
specified, there will not be a deployment.

A sourced configuration with `deployTo`:

```go
var application application
var bucket bucket
var env environment


appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
	DeployTo: []iEnvironment{
		env,
	},
})
```

## Environment

For each AWS AppConfig application, you define one or more environments. An environment is a logical deployment group of AWS
AppConfig targets, such as applications in a Beta or Production environment. You can also define environments for application
subcomponents such as the Web, Mobile, and Back-end components for your application. You can configure Amazon CloudWatch alarms
for each environment. The system monitors alarms during a configuration deployment. If an alarm is triggered, the system rolls
back the configuration.

Basic environment with monitors:

```go
var application application
var alarm alarm


appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
	Application: Application,
	Monitors: []monitor{
		appconfig.*monitor_FromCloudWatchAlarm(alarm),
	},
})
```

Environment monitors also support L1 CfnEnvironment.MonitorsProperty constructs. However, this is not the recommended approach
for CloudWatch alarms because a role will not be auto-generated if not provided.

## Extension

An extension augments your ability to inject logic or behavior at different points during the AWS AppConfig workflow of
creating or deploying a configuration.
See: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html

### AWS Lambda destination

Use an AWS Lambda as the event destination for an extension.

```go
var fn function


appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
	Actions: []action{
		appconfig.NewAction(&ActionProps{
			ActionPoints: []actionPoint{
				appconfig.*actionPoint_ON_DEPLOYMENT_START,
			},
			EventDestination: appconfig.NewLambdaDestination(fn),
		}),
	},
})
```

Lambda extension with parameters:

```go
var fn function


appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
	Actions: []action{
		appconfig.NewAction(&ActionProps{
			ActionPoints: []actionPoint{
				appconfig.*actionPoint_ON_DEPLOYMENT_START,
			},
			EventDestination: appconfig.NewLambdaDestination(fn),
		}),
	},
	Parameters: []parameter{
		appconfig.*parameter_Required(jsii.String("testParam"), jsii.String("true")),
		appconfig.*parameter_NotRequired(jsii.String("testNotRequiredParam")),
	},
})
```

### Amazon Simple Queue Service (SQS) destination

Use a queue as the event destination for an extension.

```go
var queue queue


appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
	Actions: []action{
		appconfig.NewAction(&ActionProps{
			ActionPoints: []actionPoint{
				appconfig.*actionPoint_ON_DEPLOYMENT_START,
			},
			EventDestination: appconfig.NewSqsDestination(queue),
		}),
	},
})
```

### Amazon Simple Notification Service (SNS) destination

Use an SNS topic as the event destination for an extension.

```go
var topic topic


appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
	Actions: []action{
		appconfig.NewAction(&ActionProps{
			ActionPoints: []actionPoint{
				appconfig.*actionPoint_ON_DEPLOYMENT_START,
			},
			EventDestination: appconfig.NewSnsDestination(topic),
		}),
	},
})
```

### Amazon EventBridge destination

Use the default event bus as the event destination for an extension.

```go
bus := events.EventBus_FromEventBusName(this, jsii.String("MyEventBus"), jsii.String("default"))

appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
	Actions: []action{
		appconfig.NewAction(&ActionProps{
			ActionPoints: []actionPoint{
				appconfig.*actionPoint_ON_DEPLOYMENT_START,
			},
			EventDestination: appconfig.NewEventBridgeDestination(bus),
		}),
	},
})
```

You can also add extensions and their associations directly by calling `onDeploymentComplete()` or any other action point
method on the AWS AppConfig application, configuration, or environment resource. To add an association to an existing
extension, you can call `addExtension()` on the resource.

Adding an association to an AWS AppConfig application:

```go
var application application
var extension extension
var lambdaDestination lambdaDestination


application.addExtension(extension)
application.onDeploymentComplete(lambdaDestination)
```

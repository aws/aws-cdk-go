# AWS AppConfig Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

For a high level overview of what AWS AppConfig is and how it works, please take a look here:
[What is AWS AppConfig?](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html)

## Basic Hosted Configuration Use Case

> The main way most AWS AppConfig users utilize the service is through hosted configuration, which involves storing
> configuration data directly within AWS AppConfig.

An example use case:

```go
app := appconfig.NewApplication(this, jsii.String("MyApp"))
env := appconfig.NewEnvironment(this, jsii.String("MyEnv"), &EnvironmentProps{
	Application: app,
})

appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfig"), &HostedConfigurationProps{
	Application: app,
	DeployTo: []iEnvironment{
		env,
	},
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
})
```

This will create the application and environment for your configuration and then deploy your configuration to the
specified environment.

For more information about what these resources are: [Creating feature flags and free form configuration data in AWS AppConfig](https://docs.aws.amazon.com/appconfig/latest/userguide/creating-feature-flags-and-configuration-data.html).

For more information about deploying configuration: [Deploying feature flags and configuration data in AWS AppConfig](https://docs.aws.amazon.com/appconfig/latest/userguide/deploying-feature-flags.html)

---


For an in-depth walkthrough of specific resources and how to use them, please take a look at the following sections.

## Application

[AWS AppConfig Application Documentation](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-namespace.html)

In AWS AppConfig, an application is simply an organizational
construct like a folder. Configurations and environments are
associated with the application.

When creating an application through CDK, the name and
description of an application are optional.

Create a simple application:

```go
appconfig.NewApplication(this, jsii.String("MyApplication"))
```

## Environment

[AWS AppConfig Environment Documentation](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-environment.html)

Basic environment with monitors:

```go
var application application
var alarm alarm
var compositeAlarm compositeAlarm


appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
	Application: Application,
	Monitors: []monitor{
		appconfig.*monitor_FromCloudWatchAlarm(alarm),
		appconfig.*monitor_*FromCloudWatchAlarm(compositeAlarm),
	},
})
```

Environment monitors also support L1 `CfnEnvironment.MonitorsProperty` constructs through the `fromCfnMonitorsProperty` method.
However, this is not the recommended approach for CloudWatch alarms because a role will not be auto-generated if not provided.

See [About the AWS AppConfig data plane service](https://docs.aws.amazon.com/appconfig/latest/userguide/about-data-plane.html) for more information.

### Permissions

You can grant read permission on the environment's configurations with the grantReadConfig method as follows:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"


app := appconfig.NewApplication(this, jsii.String("MyAppConfig"))
env := appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
	Application: app,
})

user := iam.NewUser(this, jsii.String("MyUser"))
env.grantReadConfig(user)
```

## Deployment Strategy

[AWS AppConfig Deployment Strategy Documentation](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html)

A deployment strategy defines how a configuration will roll out. The roll out is defined by four parameters: deployment type,
growth factor, deployment duration, and final bake time.

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

Referencing a deployment strategy by ID:

```go
appconfig.DeploymentStrategy_FromDeploymentStrategyId(this, jsii.String("MyImportedDeploymentStrategy"), appconfig.DeploymentStrategyId_FromString(jsii.String("abc123")))
```

Referencing an AWS AppConfig predefined deployment strategy by ID:

```go
appconfig.DeploymentStrategy_FromDeploymentStrategyId(this, jsii.String("MyImportedPredefinedDeploymentStrategy"), appconfig.DeploymentStrategyId_CANARY_10_PERCENT_20_MINUTES())
```

## Configuration

A configuration is a higher-level construct that can either be a `HostedConfiguration` (stored internally through AWS
AppConfig) or a `SourcedConfiguration` (stored in an Amazon S3 bucket, AWS Secrets Manager secrets, Systems Manager (SSM)
Parameter Store parameters, SSM documents, or AWS CodePipeline). This construct manages deployments on creation.

### HostedConfiguration

A hosted configuration represents configuration stored in the AWS AppConfig hosted configuration store. A hosted configuration
takes in the configuration content and associated AWS AppConfig application. On construction of a hosted configuration, the
configuration is deployed.

You can define hosted configuration content using any of the following ConfigurationContent methods:

* `fromFile` - Defines the hosted configuration content from a file (you can specify a relative path). The content type will
  be determined by the file extension unless specified.

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromFile(jsii.String("config.json")),
})
```

* `fromInlineText` - Defines the hosted configuration from inline text. The content type will be set as `text/plain`.

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
})
```

* `fromInlineJson` - Defines the hosted configuration from inline JSON. The content type will be set as `application/json` unless specified.

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineJson(jsii.String("{}")),
})
```

* `fromInlineYaml` - Defines the hosted configuration from inline YAML. The content type will be set as `application/x-yaml`.

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineYaml(jsii.String("MyConfig: This is my content.")),
})
```

* `fromInline` - Defines the hosted configuration from user-specified content types. The content type will be set as `application/octet-stream` unless specified.

```go
var application application


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInline(jsii.String("This is my configuration content.")),
})
```

AWS AppConfig supports the following types of configuration profiles.

* **[Feature flag](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile-feature-flags.html)**: Use a feature flag configuration to turn on new features that require a timely deployment, such as a product launch or announcement.
* **[Freeform](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-free-form-configurations-creating.html)**: Use a freeform configuration to carefully introduce changes to your application.

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

When you import a JSON Schema validator from a file, you can pass in a relative path.

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

The `deployTo` parameter is used to specify which environments to deploy the configuration to.

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

When more than one configuration is set to deploy to the same environment, the
deployments will occur one at a time. This is done to satisfy
[AppConfig's constraint](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-deploying.html):

> [!NOTE]
> You can only deploy one configuration at a time to an environment.
> However, you can deploy one configuration each to different environments at the same time.

The deployment order matches the order in which the configurations are declared.

```go
app := appconfig.NewApplication(this, jsii.String("MyApp"))
env := appconfig.NewEnvironment(this, jsii.String("MyEnv"), &EnvironmentProps{
	Application: app,
})

appconfig.NewHostedConfiguration(this, jsii.String("MyFirstHostedConfig"), &HostedConfigurationProps{
	Application: app,
	DeployTo: []iEnvironment{
		env,
	},
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my first configuration content.")),
})

appconfig.NewHostedConfiguration(this, jsii.String("MySecondHostedConfig"), &HostedConfigurationProps{
	Application: app,
	DeployTo: []*iEnvironment{
		env,
	},
	Content: appconfig.ConfigurationContent_*FromInlineText(jsii.String("This is my second configuration content.")),
})
```

If an application would benefit from a deployment order that differs from the
declared order, you can defer the decision by using `IEnvironment.addDeployment`
rather than the `deployTo` property.
In this example, `firstConfig` will be deployed before `secondConfig`.

```go
app := appconfig.NewApplication(this, jsii.String("MyApp"))
env := appconfig.NewEnvironment(this, jsii.String("MyEnv"), &EnvironmentProps{
	Application: app,
})

secondConfig := appconfig.NewHostedConfiguration(this, jsii.String("MySecondHostedConfig"), &HostedConfigurationProps{
	Application: app,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my second configuration content.")),
})

firstConfig := appconfig.NewHostedConfiguration(this, jsii.String("MyFirstHostedConfig"), &HostedConfigurationProps{
	Application: app,
	DeployTo: []iEnvironment{
		env,
	},
	Content: appconfig.ConfigurationContent_*FromInlineText(jsii.String("This is my first configuration content.")),
})

env.addDeployment(secondConfig)
```

Alternatively, you can defer multiple deployments in favor of
`IEnvironment.addDeployments`, which allows you to declare multiple
configurations in the order they will be deployed.
In this example the deployment order will be
`firstConfig`, then `secondConfig`, and finally `thirdConfig`.

```go
app := appconfig.NewApplication(this, jsii.String("MyApp"))
env := appconfig.NewEnvironment(this, jsii.String("MyEnv"), &EnvironmentProps{
	Application: app,
})

secondConfig := appconfig.NewHostedConfiguration(this, jsii.String("MySecondHostedConfig"), &HostedConfigurationProps{
	Application: app,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my second configuration content.")),
})

thirdConfig := appconfig.NewHostedConfiguration(this, jsii.String("MyThirdHostedConfig"), &HostedConfigurationProps{
	Application: app,
	Content: appconfig.ConfigurationContent_*FromInlineText(jsii.String("This is my third configuration content.")),
})

firstConfig := appconfig.NewHostedConfiguration(this, jsii.String("MyFirstHostedConfig"), &HostedConfigurationProps{
	Application: app,
	Content: appconfig.ConfigurationContent_*FromInlineText(jsii.String("This is my first configuration content.")),
})

env.addDeployments(firstConfig, secondConfig, thirdConfig)
```

Any mix of `deployTo`, `addDeployment`, and `addDeployments` is permitted.
The declaration order will be respected regardless of the approach used.

> [!IMPORTANT]
> If none of these options are utilized, there will not be any deployments.

You can use customer managed key to encrypt a hosted configuration. For mora information, see [Data encryption at rest for AWS AppConfig](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-security.html#appconfig-security-data-encryption).

```go
var application application
var kmsKey key


appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
	KmsKey: KmsKey,
})
```

### SourcedConfiguration

A sourced configuration represents configuration stored in any of the following:

* Amazon S3 bucket
* AWS Secrets Manager secret
* Systems Manager
* (SSM) Parameter Store parameter
* SSM document
* AWS CodePipeline.

A sourced configuration takes in the location source
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

## Deletion Protection Check

You can enable [deletion protection](https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html) on the environment and configuration profile by setting the `deletionProtectionCheck` property.

* ACCOUNT_DEFAULT: The default setting, which uses account-level deletion protection. To configure account-level deletion protection, use the UpdateAccountSettings API.
* APPLY: Instructs the deletion protection check to run, even if deletion protection is disabled at the account level. APPLY also forces the deletion protection check to run against resources created in the past hour, which are normally excluded from deletion protection checks.
* BYPASS: Instructs AWS AppConfig to bypass the deletion protection check and delete an environment even if deletion protection would have otherwise prevented it.

```go
var application application
var alarm alarm
var compositeAlarm compositeAlarm
var bucket bucket


// Environment deletion protection check
// Environment deletion protection check
appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
	Application: Application,
	DeletionProtectionCheck: appconfig.DeletionProtectionCheck_APPLY,
})

// configuration profile with deletion protection check
// configuration profile with deletion protection check
appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfigFromFile"), &HostedConfigurationProps{
	Application: Application,
	Content: appconfig.ConfigurationContent_FromFile(jsii.String("config.json")),
	DeletionProtectionCheck: appconfig.DeletionProtectionCheck_BYPASS,
})

appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
	Application: Application,
	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
	DeletionProtectionCheck: appconfig.DeletionProtectionCheck_ACCOUNT_DEFAULT,
})
```

## Extension

An extension augments your ability to inject logic or behavior at different points during the AWS AppConfig workflow of
creating or deploying a configuration. You can associate these types of tasks with AWS AppConfig applications, environments, and configuration profiles.
See: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html

An extension defines one or more actions, that it performs during an AWS AppConfig workflow. Each action is invoked either when you interact with AWS AppConfig or when AWS AppConfig is performing a process on your behalf. These invocation points are called action points. AWS AppConfig extensions support the following action points:

* PRE_START_DEPLOYMENT
* PRE_CREATE_HOSTED_CONFIGURATION_VERSION
* ON_DEPLOYMENT_START
* ON_DEPLOYMENT_STEP
* ON_DEPLOYMENT_BAKING
* ON_DEPLOYMENT_COMPLETE
* ON_DEPLOYMENT_ROLLED_BACK
* AT_DEPLOYMENT_TICK

See: https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions-about.html

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

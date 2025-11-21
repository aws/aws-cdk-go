# EC2 Image Builder Construct Library

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

## README

[Amazon EC2 Image Builder](https://docs.aws.amazon.com/imagebuilder/latest/userguide/what-is-image-builder.html) is a
fully managed AWS service that helps you automate the creation, management, and deployment of customized, secure, and
up-to-date server images. You can use Image Builder to create Amazon Machine Images (AMIs) and container images for use
across AWS Regions.

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project. It allows you to define
Image Builder pipelines, images, recipes, components, workflows, and lifecycle policies.
A component defines the sequence of steps required to customize an instance during image creation (build component) or
test an instance launched from the created image (test component). Components are created from declarative YAML or JSON
documents that describe runtime configuration for building, validating, or testing instances. Components are included
when added to the image recipe or container recipe for an image build.

EC2 Image Builder supports AWS-managed components for common tasks, AWS Marketplace components, and custom components
that you create. Components run during specific workflow phases: build and validate phases during the build stage, and
test phase during the test stage.

### Component

A component defines the sequence of steps required to customize an instance during image creation (build component) or
test an instance launched from the created image (test component). Components are created from declarative YAML or JSON
documents that describe runtime configuration for building, validating, or testing instances. Components are included
when added to the image recipe or container recipe for an image build.

EC2 Image Builder supports AWS-managed components for common tasks, AWS Marketplace components, and custom components
that you create. Components run during specific workflow phases: build and validate phases during the build stage, and
test phase during the test stage.

#### Basic Usage

Create a component with the required properties: platform and component data.

```go
component := imagebuilder.NewComponent(this, jsii.String("MyComponent"), &ComponentProps{
	Platform: imagebuilder.Platform_LINUX,
	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
		"phases": []interface{}{
			map[string]interface{}{
				"name": imagebuilder.ComponentPhaseName_BUILD,
				"steps": []map[string]interface{}{
					map[string]interface{}{
						"name": jsii.String("install-app"),
						"action": imagebuilder.ComponentAction_EXECUTE_BASH,
						"inputs": map[string][]*string{
							"commands": []*string{
								jsii.String("echo \"Installing my application...\""),
								jsii.String("yum update -y"),
							},
						},
					},
				},
			},
		},
	}),
})
```

#### Component Data Sources

##### Inline Component Data

Use `ComponentData.fromInline()` for existing YAML/JSON definitions:

```go
component := imagebuilder.NewComponent(this, jsii.String("InlineComponent"), &ComponentProps{
	Platform: imagebuilder.Platform_LINUX,
	Data: imagebuilder.ComponentData_FromInline(jsii.String(`
	name: my-component
	schemaVersion: 1.0
	phases:
	  - name: build
	    steps:
	      - name: update-os
	        action: ExecuteBash
	        inputs:
	          commands: ['yum update -y']
	`)),
})
```

##### JSON Object Component Data

Most developer-friendly approach using objects:

```go
component := imagebuilder.NewComponent(this, jsii.String("JsonComponent"), &ComponentProps{
	Platform: imagebuilder.Platform_LINUX,
	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
		"phases": []interface{}{
			map[string]interface{}{
				"name": imagebuilder.ComponentPhaseName_BUILD,
				"steps": []map[string]interface{}{
					map[string]interface{}{
						"name": jsii.String("configure-app"),
						"action": imagebuilder.ComponentAction_CREATE_FILE,
						"inputs": map[string]*string{
							"path": jsii.String("/etc/myapp/config.json"),
							"content": jsii.String("{\"env\": \"production\"}"),
						},
					},
				},
			},
		},
	}),
})
```

##### Structured Component Document

For type-safe, CDK-native definitions with enhanced properties like `timeout` and `onFailure`.

###### Defining a component step

You can define steps in the component which will be executed in order when the component is applied:

```go
step := &ComponentDocumentStep{
	Name: jsii.String("configure-app"),
	Action: imagebuilder.ComponentAction_CREATE_FILE,
	Inputs: imagebuilder.ComponentStepInputs_FromObject(map[string]interface{}{
		"path": jsii.String("/etc/myapp/config.json"),
		"content": jsii.String("{\"env\": \"production\"}"),
	}),
}
```

###### Defining a component phase

Phases group steps together, which run in sequence when building, validating or testing in the component:

```go
phase := &ComponentDocumentPhase{
	Name: imagebuilder.ComponentPhaseName_BUILD,
	Steps: []ComponentDocumentStep{
		&ComponentDocumentStep{
			Name: jsii.String("configure-app"),
			Action: imagebuilder.ComponentAction_CREATE_FILE,
			Inputs: imagebuilder.ComponentStepInputs_FromObject(map[string]interface{}{
				"path": jsii.String("/etc/myapp/config.json"),
				"content": jsii.String("{\"env\": \"production\"}"),
			}),
		},
	},
}
```

###### Defining a component

The component data defines all steps across the provided phases to execute during the build:

```go
component := imagebuilder.NewComponent(this, jsii.String("StructuredComponent"), &ComponentProps{
	Platform: imagebuilder.Platform_LINUX,
	Data: imagebuilder.ComponentData_FromComponentDocumentJsonObject(&ComponentDocument{
		SchemaVersion: imagebuilder.ComponentSchemaVersion_V1_0,
		Phases: []ComponentDocumentPhase{
			&ComponentDocumentPhase{
				Name: imagebuilder.ComponentPhaseName_BUILD,
				Steps: []ComponentDocumentStep{
					&ComponentDocumentStep{
						Name: jsii.String("install-with-timeout"),
						Action: imagebuilder.ComponentAction_EXECUTE_BASH,
						Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
						OnFailure: imagebuilder.ComponentOnFailure_CONTINUE,
						Inputs: imagebuilder.ComponentStepInputs_FromObject(map[string]interface{}{
							"commands": []interface{}{
								jsii.String("./install-script.sh"),
							},
						}),
					},
				},
			},
		},
	}),
})
```

##### S3 Component Data

For those components you want to upload or have uploaded to S3:

```go
// Upload a local file
componentFromAsset := imagebuilder.NewComponent(this, jsii.String("AssetComponent"), &ComponentProps{
	Platform: imagebuilder.Platform_LINUX,
	Data: imagebuilder.ComponentData_FromAsset(this, jsii.String("ComponentAsset"), jsii.String("./my-component.yml")),
})

// Reference an existing S3 object
bucket := s3.Bucket_FromBucketName(this, jsii.String("ComponentBucket"), jsii.String("my-components-bucket"))
componentFromS3 := imagebuilder.NewComponent(this, jsii.String("S3Component"), &ComponentProps{
	Platform: imagebuilder.Platform_LINUX,
	Data: imagebuilder.ComponentData_FromS3(bucket, jsii.String("components/my-component.yml")),
})
```

#### Encrypt component data with a KMS key

You can encrypt component data with a KMS key, so that only principals with access to decrypt with the key are able to
access the component data.

```go
component := imagebuilder.NewComponent(this, jsii.String("EncryptedComponent"), &ComponentProps{
	Platform: imagebuilder.Platform_LINUX,
	KmsKey: kms.NewKey(this, jsii.String("ComponentKey")),
	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
		"phases": []interface{}{
			map[string]interface{}{
				"name": imagebuilder.ComponentPhaseName_BUILD,
				"steps": []map[string]interface{}{
					map[string]interface{}{
						"name": jsii.String("secure-setup"),
						"action": imagebuilder.ComponentAction_EXECUTE_BASH,
						"inputs": map[string][]*string{
							"commands": []*string{
								jsii.String("echo \"This component data is encrypted with KMS\""),
							},
						},
					},
				},
			},
		},
	}),
})
```

#### AWS-Managed Components

AWS provides a collection of managed components for common tasks:

```go
// Install AWS CLI v2
awsCliComponent := imagebuilder.AwsManagedComponent_AwsCliV2(this, jsii.String("AwsCli"), &AwsManagedComponentAttributes{
	Platform: imagebuilder.Platform_LINUX,
})

// Update the operating system
updateComponent := imagebuilder.AwsManagedComponent_UpdateOS(this, jsii.String("UpdateOS"), &AwsManagedComponentAttributes{
	Platform: imagebuilder.Platform_LINUX,
})

// Reference any AWS-managed component by name
customAwsComponent := imagebuilder.AwsManagedComponent_FromAwsManagedComponentName(this, jsii.String("CloudWatchAgent"), jsii.String("amazon-cloudwatch-agent-linux"))
```

#### AWS Marketplace Components

You can reference AWS Marketplace components using the marketplace component name and its product ID:

```go
marketplaceComponent := imagebuilder.AwsMarketplaceComponent_FromAwsMarketplaceComponentAttributes(this, jsii.String("MarketplaceComponent"), &AwsMarketplaceComponentAttributes{
	ComponentName: jsii.String("my-marketplace-component"),
	MarketplaceProductId: jsii.String("prod-1234567890abcdef0"),
})
```

### Infrastructure Configuration

Infrastructure configuration defines the compute resources and environment settings used during the image building
process. This includes instance types, IAM instance profile, VPC settings, subnets, security groups, SNS topics for
notifications, logging configuration, and troubleshooting settings like whether to terminate instances on failure or
keep them running for debugging. These settings are applied to builds when included in an image or an image pipeline.

```go
infrastructureConfiguration := imagebuilder.NewInfrastructureConfiguration(this, jsii.String("InfrastructureConfiguration"), &InfrastructureConfigurationProps{
	InfrastructureConfigurationName: jsii.String("test-infrastructure-configuration"),
	Description: jsii.String("An Infrastructure Configuration"),
	// Optional - instance types to use for build/test
	InstanceTypes: []InstanceType{
		ec2.InstanceType_*Of(ec2.InstanceClass_STANDARD7_INTEL, ec2.InstanceSize_LARGE),
		ec2.InstanceType_*Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_LARGE),
	},
	// Optional - create an instance profile with necessary permissions
	InstanceProfile: iam.NewInstanceProfile(this, jsii.String("InstanceProfile"), &InstanceProfileProps{
		InstanceProfileName: jsii.String("test-instance-profile"),
		Role: iam.NewRole(this, jsii.String("InstanceProfileRole"), &RoleProps{
			AssumedBy: iam.ServicePrincipal_FromStaticServicePrincipleName(jsii.String("ec2.amazonaws.com")),
			ManagedPolicies: []IManagedPolicy{
				iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonSSMManagedInstanceCore")),
				iam.ManagedPolicy_*FromAwsManagedPolicyName(jsii.String("EC2InstanceProfileForImageBuilder")),
			},
		}),
	}),
	// Use VPC network configuration
	Vpc: Vpc,
	SubnetSelection: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
	SecurityGroups: []ISecurityGroup{
		ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("SecurityGroup"), vpc.VpcDefaultSecurityGroup),
	},
	KeyPair: ec2.KeyPair_FromKeyPairName(this, jsii.String("KeyPair"), jsii.String("imagebuilder-instance-key-pair")),
	TerminateInstanceOnFailure: jsii.Boolean(true),
	// Optional - IMDSv2 settings
	HttpTokens: imagebuilder.HttpTokens_REQUIRED,
	HttpPutResponseHopLimit: jsii.Number(1),
	// Optional - publish image completion messages to an SNS topic
	NotificationTopic: sns.Topic_FromTopicArn(this, jsii.String("Topic"), this.FormatArn(&ArnComponents{
		Service: jsii.String("sns"),
		Resource: jsii.String("image-builder-topic"),
	})),
	// Optional - log settings. Logging is enabled by default
	Logging: &InfrastructureConfigurationLogging{
		S3Bucket: s3.Bucket_FromBucketName(this, jsii.String("LogBucket"), fmt.Sprintf("imagebuilder-logging-%v", awscdk.Aws_ACCOUNT_ID())),
		S3KeyPrefix: jsii.String("imagebuilder-logs"),
	},
	// Optional - host placement settings
	Ec2InstanceAvailabilityZone: awscdk.*stack_*Of(this).availabilityZones[jsii.Number(0)],
	Ec2InstanceHostId: dedicatedHost.AttrHostId,
	Ec2InstanceTenancy: imagebuilder.Tenancy_HOST,
	ResourceTags: map[string]*string{
		"Environment": jsii.String("production"),
	},
})
```

### Distribution Configuration

Distribution configuration defines how and where your built images are distributed after successful creation. For AMIs,
this includes target AWS Regions, KMS encryption keys, account sharing permissions, License Manager associations, and
launch template configurations. For container images, it specifies the target Amazon ECR repositories across regions.
A distribution configuration can be associated with an image or an image pipeline to define these distribution settings
for image builds.

#### AMI Distributions

AMI distributions can be defined to copy and modify AMIs in different accounts and regions, and apply them to launch
templates, SSM parameters, etc.:

```go
distributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("DistributionConfiguration"), &DistributionConfigurationProps{
	DistributionConfigurationName: jsii.String("test-distribution-configuration"),
	Description: jsii.String("A Distribution Configuration"),
	AmiDistributions: []AmiDistribution{
		&AmiDistribution{
			// Distribute AMI to us-east-2 and publish the AMI ID to an SSM parameter
			Region: jsii.String("us-east-2"),
			SsmParameters: []SSMParameterConfigurations{
				&SSMParameterConfigurations{
					Parameter: ssm.StringParameter_FromStringParameterAttributes(this, jsii.String("CrossRegionParameter"), &StringParameterAttributes{
						ParameterName: jsii.String("/imagebuilder/ami"),
						ForceDynamicReference: jsii.Boolean(true),
					}),
				},
			},
		},
	},
})

// For AMI-based image builds - add an AMI distribution in the current region
distributionConfiguration.AddAmiDistributions(&AmiDistribution{
	AmiName: jsii.String("imagebuilder-{{ imagebuilder:buildDate }}"),
	AmiDescription: jsii.String("Build AMI"),
	AmiKmsKey: kms.Key_FromLookup(this, jsii.String("ComponentKey"), &KeyLookupOptions{
		AliasName: jsii.String("alias/distribution-encryption-key"),
	}),
	// Copy the AMI to different accounts
	AmiTargetAccountIds: []*string{
		jsii.String("123456789012"),
		jsii.String("098765432109"),
	},
	// Add launch permissions on the AMI
	AmiLaunchPermission: &AmiLaunchPermission{
		OrganizationArns: []*string{
			this.FormatArn(&ArnComponents{
				Region: jsii.String(""),
				Service: jsii.String("organizations"),
				Resource: jsii.String("organization"),
				ResourceName: jsii.String("o-1234567abc"),
			}),
		},
		OrganizationalUnitArns: []*string{
			this.*FormatArn(&ArnComponents{
				Region: jsii.String(""),
				Service: jsii.String("organizations"),
				Resource: jsii.String("ou"),
				ResourceName: jsii.String("o-1234567abc/ou-a123-b4567890"),
			}),
		},
		IsPublicUserGroup: jsii.Boolean(true),
		AccountIds: []*string{
			jsii.String("234567890123"),
		},
	},
	// Attach tags to the AMI
	AmiTags: map[string]*string{
		"Environment": jsii.String("production"),
		"Version": jsii.String("{{ imagebuilder:buildVersion }}"),
	},
	// Optional - publish the distributed AMI ID to an SSM parameter
	SsmParameters: []SSMParameterConfigurations{
		&SSMParameterConfigurations{
			Parameter: ssm.StringParameter_*FromStringParameterAttributes(this, jsii.String("Parameter"), &StringParameterAttributes{
				ParameterName: jsii.String("/imagebuilder/ami"),
				ForceDynamicReference: jsii.Boolean(true),
			}),
		},
		&SSMParameterConfigurations{
			AmiAccount: jsii.String("098765432109"),
			DataType: ssm.ParameterDataType_TEXT,
			Parameter: ssm.StringParameter_*FromStringParameterAttributes(this, jsii.String("CrossAccountParameter"), &StringParameterAttributes{
				ParameterName: jsii.String("imagebuilder-prod-ami"),
				ForceDynamicReference: jsii.Boolean(true),
			}),
		},
	},
	// Optional - create a new launch template version with the distributed AMI ID
	LaunchTemplates: []LaunchTemplateConfiguration{
		&LaunchTemplateConfiguration{
			LaunchTemplate: ec2.LaunchTemplate_FromLaunchTemplateAttributes(this, jsii.String("LaunchTemplate"), &LaunchTemplateAttributes{
				LaunchTemplateId: jsii.String("lt-1234"),
			}),
			SetDefaultVersion: jsii.Boolean(true),
		},
		&LaunchTemplateConfiguration{
			AccountId: jsii.String("123456789012"),
			LaunchTemplate: ec2.LaunchTemplate_*FromLaunchTemplateAttributes(this, jsii.String("CrossAccountLaunchTemplate"), &LaunchTemplateAttributes{
				LaunchTemplateId: jsii.String("lt-5678"),
			}),
			SetDefaultVersion: jsii.Boolean(true),
		},
	},
	// Optional - enable Fast Launch on an imported launch template
	FastLaunchConfigurations: []FastLaunchConfiguration{
		&FastLaunchConfiguration{
			Enabled: jsii.Boolean(true),
			LaunchTemplate: ec2.LaunchTemplate_*FromLaunchTemplateAttributes(this, jsii.String("FastLaunchLT"), &LaunchTemplateAttributes{
				LaunchTemplateName: jsii.String("fast-launch-lt"),
			}),
			MaxParallelLaunches: jsii.Number(10),
			TargetSnapshotCount: jsii.Number(2),
		},
	},
	// Optional - license configurations to apply to the AMI
	LicenseConfigurationArns: []*string{
		jsii.String("arn:aws:license-manager:us-west-2:123456789012:license-configuration:lic-abcdefghijklmnopqrstuvwxyz"),
	},
})
```

#### Container Distributions

##### Container repositories

Container distributions can be configured to distribute to ECR repositories:

```go
ecrRepository := ecr.Repository_FromRepositoryName(this, jsii.String("ECRRepository"), jsii.String("my-repo"))
imageBuilderRepository := imagebuilder.Repository_FromEcr(ecrRepository)
```

##### Defining a container distribution

You can configure the container repositories as well as the description and tags applied to the distributed container
images:

```go
ecrRepository := ecr.Repository_FromRepositoryName(this, jsii.String("ECRRepository"), jsii.String("my-repo"))
containerRepository := imagebuilder.Repository_FromEcr(ecrRepository)
containerDistributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("ContainerDistributionConfiguration"))

containerDistributionConfiguration.AddContainerDistributions(&ContainerDistribution{
	ContainerRepository: ContainerRepository,
	ContainerDescription: jsii.String("Test container image"),
	ContainerTags: []*string{
		jsii.String("latest"),
		jsii.String("latest-1.0"),
	},
})
```

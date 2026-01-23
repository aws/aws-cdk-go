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

### Image Pipeline

An image pipeline provides the automation framework for building secure AMIs and container images. The pipeline
orchestrates the entire image creation process by combining an image recipe or container recipe with infrastructure
configuration and distribution configuration. Pipelines can run on a schedule or be triggered manually, and they manage
the build, test, and distribution phases automatically.

#### Image Pipeline Basic Usage

Create a simple AMI pipeline with just an image recipe:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("MyImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
})

imagePipeline := imagebuilder.NewImagePipeline(this, jsii.String("MyImagePipeline"), &ImagePipelineProps{
	Recipe: exampleImageRecipe,
})
```

Create a simple container pipeline with just a container recipe:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("MyContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
})

containerPipeline := imagebuilder.NewImagePipeline(this, jsii.String("MyContainerPipeline"), &ImagePipelineProps{
	Recipe: exampleContainerRecipe,
})
```

#### Image Pipeline Scheduling

##### Manual Pipeline Execution

Create a pipeline that runs only when manually triggered:

```go
manualPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ManualPipeline"), &ImagePipelineProps{
	ImagePipelineName: jsii.String("my-manual-pipeline"),
	Description: jsii.String("Pipeline triggered manually for production builds"),
	Recipe: exampleImageRecipe,
})

// Grant Lambda function permission to trigger the pipeline
manualPipeline.grantStartExecution(lambdaRole)
```

##### Automated Pipeline Scheduling

Schedule a pipeline to run automatically using cron expressions:

```go
weeklyPipeline := imagebuilder.NewImagePipeline(this, jsii.String("WeeklyPipeline"), &ImagePipelineProps{
	ImagePipelineName: jsii.String("weekly-build-pipeline"),
	Recipe: exampleImageRecipe,
	Schedule: &ImagePipelineSchedule{
		Expression: events.Schedule_Cron(&CronOptions{
			Minute: jsii.String("0"),
			Hour: jsii.String("6"),
			WeekDay: jsii.String("MON"),
		}),
	},
})
```

Use rate expressions for regular intervals:

```go
dailyPipeline := imagebuilder.NewImagePipeline(this, jsii.String("DailyPipeline"), &ImagePipelineProps{
	Recipe: exampleContainerRecipe,
	Schedule: &ImagePipelineSchedule{
		Expression: events.Schedule_Rate(awscdk.Duration_Days(jsii.Number(1))),
	},
})
```

##### Pipeline Schedule Configuration

Configure advanced scheduling options:

```go
advancedSchedulePipeline := imagebuilder.NewImagePipeline(this, jsii.String("AdvancedSchedulePipeline"), &ImagePipelineProps{
	Recipe: exampleImageRecipe,
	Schedule: &ImagePipelineSchedule{
		Expression: events.Schedule_Rate(awscdk.Duration_Days(jsii.Number(7))),
		// Only trigger when dependencies are updated (new base images, components, etc.)
		StartCondition: imagebuilder.ScheduleStartCondition_EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE,
		// Automatically disable after 3 consecutive failures
		AutoDisableFailureCount: jsii.Number(3),
	},
	// Start enabled
	Status: imagebuilder.ImagePipelineStatus_ENABLED,
})
```

#### Image Pipeline Configuration

##### Infrastructure and Distribution in Image Pipelines

Configure custom infrastructure and distribution settings:

```go
infrastructureConfiguration := imagebuilder.NewInfrastructureConfiguration(this, jsii.String("Infrastructure"), &InfrastructureConfigurationProps{
	InfrastructureConfigurationName: jsii.String("production-infrastructure"),
	InstanceTypes: []InstanceType{
		ec2.InstanceType_*Of(ec2.InstanceClass_COMPUTE7_INTEL, ec2.InstanceSize_LARGE),
	},
	Vpc: vpc,
	SubnetSelection: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
})

distributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("Distribution"))
distributionConfiguration.AddAmiDistributions(&AmiDistribution{
	AmiName: jsii.String("production-ami-{{ imagebuilder:buildDate }}"),
	AmiTargetAccountIds: []*string{
		jsii.String("123456789012"),
		jsii.String("098765432109"),
	},
})

productionPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ProductionPipeline"), &ImagePipelineProps{
	Recipe: exampleImageRecipe,
	InfrastructureConfiguration: infrastructureConfiguration,
	DistributionConfiguration: distributionConfiguration,
})
```

##### Pipeline Logging Configuration

Configure custom CloudWatch log groups for pipeline and image logs:

```go
pipelineLogGroup := logs.NewLogGroup(this, jsii.String("PipelineLogGroup"), &LogGroupProps{
	LogGroupName: jsii.String("/custom/imagebuilder/pipeline/logs"),
	Retention: logs.RetentionDays_ONE_MONTH,
})

imageLogGroup := logs.NewLogGroup(this, jsii.String("ImageLogGroup"), &LogGroupProps{
	LogGroupName: jsii.String("/custom/imagebuilder/image/logs"),
	Retention: logs.RetentionDays_ONE_WEEK,
})

loggedPipeline := imagebuilder.NewImagePipeline(this, jsii.String("LoggedPipeline"), &ImagePipelineProps{
	Recipe: exampleImageRecipe,
	ImagePipelineLogGroup: pipelineLogGroup,
	ImageLogGroup: imageLogGroup,
})
```

##### Workflow Integration in Image Pipelines

Use AWS-managed workflows for common pipeline phases:

```go
workflowPipeline := imagebuilder.NewImagePipeline(this, jsii.String("WorkflowPipeline"), &ImagePipelineProps{
	Recipe: exampleImageRecipe,
	Workflows: []WorkflowConfiguration{
		&WorkflowConfiguration{
			Workflow: imagebuilder.AmazonManagedWorkflow_BuildImage(this, jsii.String("BuildWorkflow")),
		},
		&WorkflowConfiguration{
			Workflow: imagebuilder.AmazonManagedWorkflow_TestImage(this, jsii.String("TestWorkflow")),
		},
	},
})
```

For container pipelines, use container-specific workflows:

```go
containerWorkflowPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ContainerWorkflowPipeline"), &ImagePipelineProps{
	Recipe: exampleContainerRecipe,
	Workflows: []WorkflowConfiguration{
		&WorkflowConfiguration{
			Workflow: imagebuilder.AmazonManagedWorkflow_BuildContainer(this, jsii.String("BuildContainer")),
		},
		&WorkflowConfiguration{
			Workflow: imagebuilder.AmazonManagedWorkflow_TestContainer(this, jsii.String("TestContainer")),
		},
		&WorkflowConfiguration{
			Workflow: imagebuilder.AmazonManagedWorkflow_DistributeContainer(this, jsii.String("DistributeContainer")),
		},
	},
})
```

##### Advanced Features in Image Pipelines

Configure image scanning for container pipelines:

```go
scanningRepository := ecr.NewRepository(this, jsii.String("ScanningRepo"))

scannedContainerPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ScannedContainerPipeline"), &ImagePipelineProps{
	Recipe: exampleContainerRecipe,
	ImageScanningEnabled: jsii.Boolean(true),
	ImageScanningEcrRepository: scanningRepository,
	ImageScanningEcrTags: []*string{
		jsii.String("security-scan"),
		jsii.String("latest"),
	},
})
```

Control metadata collection and testing:

```go
controlledPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ControlledPipeline"), &ImagePipelineProps{
	Recipe: exampleImageRecipe,
	EnhancedImageMetadataEnabled: jsii.Boolean(true),
	 // Collect detailed OS and package info
	ImageTestsEnabled: jsii.Boolean(false),
})
```

#### Image Pipeline Events

##### Pipeline Event Handling

Handle specific pipeline events:

```go
// Monitor CVE detection
examplePipeline.onCVEDetected(jsii.String("CVEAlert"), &OnEventOptions{
	Target: targets.NewSnsTopic(topic),
})

// Handle pipeline auto-disable events
examplePipeline.onImagePipelineAutoDisabled(jsii.String("PipelineDisabledAlert"), &OnEventOptions{
	Target: targets.NewLambdaFunction(lambdaFunction),
})
```

#### Importing Image Pipelines

Reference existing pipelines created outside CDK:

```go
// Import by name
existingPipelineByName := imagebuilder.ImagePipeline_FromImagePipelineName(this, jsii.String("ExistingPipelineByName"), jsii.String("my-existing-pipeline"))

// Import by ARN
existingPipelineByArn := imagebuilder.ImagePipeline_FromImagePipelineArn(this, jsii.String("ExistingPipelineByArn"), jsii.String("arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/imported-pipeline"))

// Grant permissions to imported pipelines
automationRole := iam.NewRole(this, jsii.String("AutomationRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

existingPipelineByName.GrantStartExecution(automationRole)
existingPipelineByArn.GrantRead(lambdaRole)
```

### Image

An image is the output resource created by Image Builder, consisting of an AMI or container image plus metadata such as
version, platform, and creation details. Images are used as base images for future builds and can be shared across AWS
accounts. While images are the output from image pipeline executions, they can also be created in an ad-hoc manner
outside a pipeline, defined as a standalone resource.

#### Image Basic Usage

Create a simple AMI-based image from an image recipe:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("MyImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
})

amiImage := imagebuilder.NewImage(this, jsii.String("MyAmiImage"), &ImageProps{
	Recipe: imageRecipe,
})
```

Create a simple container image from a container recipe:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("MyContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
})

containerImage := imagebuilder.NewImage(this, jsii.String("MyContainerImage"), &ImageProps{
	Recipe: containerRecipe,
})
```

#### AWS-Managed Images

##### Pre-defined OS Images

Use AWS-managed images for common operating systems:

```go
// Amazon Linux 2023 AMI for x86_64
amazonLinux2023Ami := imagebuilder.AmazonManagedImage_AmazonLinux2023(this, jsii.String("AmazonLinux2023"), &AmazonManagedImageOptions{
	ImageType: imagebuilder.ImageType_AMI,
	ImageArchitecture: imagebuilder.ImageArchitecture_X86_64,
})

// Ubuntu 22.04 AMI for ARM64
ubuntu2204Ami := imagebuilder.AmazonManagedImage_UbuntuServer2204(this, jsii.String("Ubuntu2204"), &AmazonManagedImageOptions{
	ImageType: imagebuilder.ImageType_AMI,
	ImageArchitecture: imagebuilder.ImageArchitecture_ARM64,
})

// Windows Server 2022 Full AMI
windows2022Ami := imagebuilder.AmazonManagedImage_WindowsServer2022Full(this, jsii.String("Windows2022"), &AmazonManagedImageOptions{
	ImageType: imagebuilder.ImageType_AMI,
	ImageArchitecture: imagebuilder.ImageArchitecture_X86_64,
})

// Use as base image in recipe
managedImageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("ManagedImageRecipe"), &ImageRecipeProps{
	BaseImage: amazonLinux2023Ami.ToBaseImage(),
})
```

##### Custom AWS-Managed Images

Import AWS-managed images by name or attributes:

```go
// Import by name
managedImageByName := imagebuilder.AmazonManagedImage_FromAmazonManagedImageName(this, jsii.String("ManagedImageByName"), jsii.String("amazon-linux-2023-x86"))

// Import by attributes with specific version
managedImageByAttributes := imagebuilder.AmazonManagedImage_FromAmazonManagedImageAttributes(this, jsii.String("ManagedImageByAttributes"), &AmazonManagedImageAttributes{
	ImageName: jsii.String("ubuntu-server-22-lts-x86"),
	ImageVersion: jsii.String("2024.11.25"),
})
```

#### Image Configuration

##### Infrastructure and Distribution in Images

Configure custom infrastructure and distribution settings:

```go
infrastructureConfiguration := imagebuilder.NewInfrastructureConfiguration(this, jsii.String("Infrastructure"), &InfrastructureConfigurationProps{
	InfrastructureConfigurationName: jsii.String("production-infrastructure"),
	InstanceTypes: []InstanceType{
		ec2.InstanceType_*Of(ec2.InstanceClass_COMPUTE7_INTEL, ec2.InstanceSize_LARGE),
	},
	Vpc: vpc,
	SubnetSelection: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
	},
})

distributionConfiguration := imagebuilder.NewDistributionConfiguration(this, jsii.String("Distribution"))
distributionConfiguration.AddAmiDistributions(&AmiDistribution{
	AmiName: jsii.String("production-ami-{{ imagebuilder:buildDate }}"),
	AmiTargetAccountIds: []*string{
		jsii.String("123456789012"),
		jsii.String("098765432109"),
	},
})

productionImage := imagebuilder.NewImage(this, jsii.String("ProductionImage"), &ImageProps{
	Recipe: exampleImageRecipe,
	InfrastructureConfiguration: infrastructureConfiguration,
	DistributionConfiguration: distributionConfiguration,
})
```

##### Logging Configuration

Configure custom CloudWatch log groups for image builds:

```go
logGroup := logs.NewLogGroup(this, jsii.String("ImageLogGroup"), &LogGroupProps{
	LogGroupName: jsii.String("/custom/imagebuilder/image/logs"),
	Retention: logs.RetentionDays_ONE_MONTH,
})

loggedImage := imagebuilder.NewImage(this, jsii.String("LoggedImage"), &ImageProps{
	Recipe: exampleImageRecipe,
	LogGroup: logGroup,
})
```

##### Workflow Integration in Images

Use workflows for custom build, test, and distribution processes:

```go
imageWithWorkflows := imagebuilder.NewImage(this, jsii.String("ImageWithWorkflows"), &ImageProps{
	Recipe: exampleImageRecipe,
	Workflows: []WorkflowConfiguration{
		&WorkflowConfiguration{
			Workflow: imagebuilder.AmazonManagedWorkflow_BuildImage(this, jsii.String("BuildWorkflow")),
		},
		&WorkflowConfiguration{
			Workflow: imagebuilder.AmazonManagedWorkflow_TestImage(this, jsii.String("TestWorkflow")),
		},
	},
})
```

##### Advanced Features in Images

Configure image scanning, metadata collection, and testing:

```go
scanningRepository := ecr.NewRepository(this, jsii.String("ScanningRepository"))

advancedContainerImage := imagebuilder.NewImage(this, jsii.String("AdvancedContainerImage"), &ImageProps{
	Recipe: exampleContainerRecipe,
	ImageScanningEnabled: jsii.Boolean(true),
	ImageScanningEcrRepository: scanningRepository,
	ImageScanningEcrTags: []*string{
		jsii.String("security-scan"),
		jsii.String("latest"),
	},
	EnhancedImageMetadataEnabled: jsii.Boolean(true),
	ImageTestsEnabled: jsii.Boolean(false),
})
```

#### Importing Images

Reference existing images created outside CDK:

```go
// Import by name
existingImageByName := imagebuilder.Image_FromImageName(this, jsii.String("ExistingImageByName"), jsii.String("my-existing-image"))

// Import by ARN
existingImageByArn := imagebuilder.Image_FromImageArn(this, jsii.String("ExistingImageByArn"), jsii.String("arn:aws:imagebuilder:us-east-1:123456789012:image/imported-image/1.0.0"))

// Import by attributes
existingImageByAttributes := imagebuilder.Image_FromImageAttributes(this, jsii.String("ExistingImageByAttributes"), &ImageAttributes{
	ImageName: jsii.String("shared-base-image"),
	ImageVersion: jsii.String("2024.11.25"),
})

// Grant permissions to imported images
role := iam.NewRole(this, jsii.String("ImageAccessRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

existingImageByName.GrantRead(role)
existingImageByArn.Grant(role, jsii.String("imagebuilder:GetImage"), jsii.String("imagebuilder:ListImagePackages"))
```

### Image Recipe

#### Image Recipe Basic Usage

Create an image recipe with the required base image:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("MyImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
})
```

#### Image Recipe Base Images

To create a recipe, you have to select a base image to build and customize from. This base image can be referenced from
various sources, such as from SSM parameters, AWS Marketplace products, and AMI IDs directly.

##### SSM Parameters

Using SSM parameter references:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("SsmImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
})

// Using an SSM parameter construct
parameter := ssm.StringParameter_FromStringParameterName(this, jsii.String("BaseImageParameter"), jsii.String("/aws/service/ami-windows-latest/Windows_Server-2022-English-Full-Base"))
windowsRecipe := imagebuilder.NewImageRecipe(this, jsii.String("WindowsImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_FromSsmParameter(parameter),
})
```

##### AMI IDs

When you have a specific AMI to use:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("AmiImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_FromAmiId(jsii.String("ami-12345678")),
})
```

##### Marketplace Images

For marketplace base images:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("MarketplaceImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_FromMarketplaceProductId(jsii.String("prod-1234567890abcdef0")),
})
```

#### Image Recipe Components

Components from various sources, such as custom-owned, AWS-owned, or AWS Marketplace-owned, can optionally be included
in recipes. For parameterized components, you are able to provide the parameters to use in the recipe, which will be
applied during the image build when executing components.

##### Custom Components in Image Recipes

Add your own components to the recipe:

```go
customComponent := imagebuilder.NewComponent(this, jsii.String("MyComponent"), &ComponentProps{
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
								jsii.String("yum install -y my-application"),
							},
						},
					},
				},
			},
		},
	}),
})

imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("ComponentImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
	Components: []ComponentConfiguration{
		&ComponentConfiguration{
			Component: customComponent,
		},
	},
})
```

##### AWS-Managed Components in Image Recipes

Use pre-built AWS components:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("AmazonManagedImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
	Components: []ComponentConfiguration{
		&ComponentConfiguration{
			Component: imagebuilder.AmazonManagedComponent_UpdateOs(this, jsii.String("UpdateOS"), &AmazonManagedComponentOptions{
				Platform: imagebuilder.Platform_LINUX,
			}),
		},
		&ComponentConfiguration{
			Component: imagebuilder.AmazonManagedComponent_AwsCliV2(this, jsii.String("AwsCli"), &AmazonManagedComponentOptions{
				Platform: imagebuilder.Platform_LINUX,
			}),
		},
	},
})
```

##### Component Parameters in Image Recipes

Pass parameters to components that accept them:

```go
parameterizedComponent := imagebuilder.Component_FromComponentName(this, jsii.String("ParameterizedComponent"), jsii.String("my-parameterized-component"))

imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("ParameterizedImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
	Components: []ComponentConfiguration{
		&ComponentConfiguration{
			Component: parameterizedComponent,
			Parameters: map[string]ComponentParameterValue{
				"environment": imagebuilder.ComponentParameterValue_fromString(jsii.String("production")),
				"version": imagebuilder.ComponentParameterValue_fromString(jsii.String("1.0.0")),
			},
		},
	},
})
```

#### Image Recipe Configuration

##### Block Device Configuration

Configure storage for the build instance:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("BlockDeviceImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
	BlockDevices: []BlockDevice{
		&BlockDevice{
			DeviceName: jsii.String("/dev/sda1"),
			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(100), &EbsDeviceOptions{
				Encrypted: jsii.Boolean(true),
				VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD_GP3,
			}),
		},
	},
})
```

##### AMI Tagging

Tag the output AMI:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("TaggedImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
	AmiTags: map[string]*string{
		"Environment": jsii.String("Production"),
		"Application": jsii.String("WebServer"),
		"Owner": jsii.String("DevOps Team"),
	},
})
```

### Container Recipe

A container recipe is similar to an image recipe but specifically for container images. It defines the base container
image and components applied to produce the desired configuration for the output container image. Container recipes work
with Docker images from DockerHub, Amazon ECR, or Amazon-managed container images as starting points.

#### Container Recipe Basic Usage

Create a container recipe with the required base image and target repository:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("MyContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
})
```

#### Container Recipe Base Images

##### DockerHub Images

Using public Docker Hub images:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("DockerHubContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
})
```

##### ECR Images

Using images from your own ECR repositories:

```go
sourceRepo := ecr.Repository_FromRepositoryName(this, jsii.String("SourceRepo"), jsii.String("my-base-image"))
targetRepo := ecr.Repository_FromRepositoryName(this, jsii.String("TargetRepo"), jsii.String("my-container-repo"))

containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("EcrContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_FromEcr(sourceRepo, jsii.String("1.0.0")),
	TargetRepository: imagebuilder.Repository_*FromEcr(targetRepo),
})
```

##### ECR Public Images

Using images from Amazon ECR Public:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("EcrPublicContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_FromEcrPublic(jsii.String("amazonlinux"), jsii.String("amazonlinux"), jsii.String("2023")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
})
```

#### Container Recipe Components

##### Custom Components in Container Recipes

Add your own components to the container recipe:

```go
customComponent := imagebuilder.NewComponent(this, jsii.String("MyComponent"), &ComponentProps{
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
								jsii.String("yum install -y my-container-application"),
							},
						},
					},
				},
			},
		},
	}),
})

containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("ComponentContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
	Components: []ComponentConfiguration{
		&ComponentConfiguration{
			Component: customComponent,
		},
	},
})
```

##### AWS-Managed Components in Container Recipes

Use pre-built AWS components:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("AmazonManagedContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
	Components: []ComponentConfiguration{
		&ComponentConfiguration{
			Component: imagebuilder.AmazonManagedComponent_UpdateOs(this, jsii.String("UpdateOS"), &AmazonManagedComponentOptions{
				Platform: imagebuilder.Platform_LINUX,
			}),
		},
		&ComponentConfiguration{
			Component: imagebuilder.AmazonManagedComponent_AwsCliV2(this, jsii.String("AwsCli"), &AmazonManagedComponentOptions{
				Platform: imagebuilder.Platform_LINUX,
			}),
		},
	},
})
```

#### Container Recipe Configuration

##### Custom Dockerfile

Provide your own Dockerfile template:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("CustomDockerfileContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
	Dockerfile: imagebuilder.DockerfileData_FromInline(jsii.String(`
	FROM {{{ imagebuilder:parentImage }}}
	CMD ["echo", "Hello, world!"]
	{{{ imagebuilder:environments }}}
	{{{ imagebuilder:components }}}
	`)),
})
```

##### Instance Configuration

Configure the build instance:

```go
containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("InstanceConfigContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
	// Custom ECS-optimized AMI for building
	InstanceImage: imagebuilder.ContainerInstanceImage_FromSsmParameterName(jsii.String("/aws/service/ecs/optimized-ami/amazon-linux-2023/recommended/image_id")),
	// Additional storage for build process
	InstanceBlockDevices: []BlockDevice{
		&BlockDevice{
			DeviceName: jsii.String("/dev/xvda"),
			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(50), &EbsDeviceOptions{
				Encrypted: jsii.Boolean(true),
				VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD_GP3,
			}),
		},
	},
})
```

### Component

A component defines the sequence of steps required to customize an instance during image creation (build component) or
test an instance launched from the created image (test component). Components are created from declarative YAML or JSON
documents that describe runtime configuration for building, validating, or testing instances. Components are included
when added to the image recipe or container recipe for an image build.

EC2 Image Builder supports AWS-managed components for common tasks, AWS Marketplace components, and custom components
that you create. Components run during specific workflow phases: build and validate phases during the build stage, and
test phase during the test stage.

#### Basic Component Usage

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
awsCliComponent := imagebuilder.AmazonManagedComponent_AwsCliV2(this, jsii.String("AwsCli"), &AmazonManagedComponentOptions{
	Platform: imagebuilder.Platform_LINUX,
})

// Update the operating system
updateComponent := imagebuilder.AmazonManagedComponent_UpdateOs(this, jsii.String("UpdateOS"), &AmazonManagedComponentOptions{
	Platform: imagebuilder.Platform_LINUX,
})

// Reference any AWS-managed component by name
customAwsComponent := imagebuilder.AmazonManagedComponent_FromAmazonManagedComponentName(this, jsii.String("CloudWatchAgent"), jsii.String("amazon-cloudwatch-agent-linux"))
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
	Ec2InstanceHostId: dedicatedHost.attrHostId,
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

### Workflow

Workflows define the sequence of steps that Image Builder performs during image creation. There are three workflow
types: BUILD (image building), TEST (testing images), and DISTRIBUTION (distributing container images).

#### Basic Workflow Usage

Create a workflow with the required properties: workflow type and workflow data.

```go
workflow := imagebuilder.NewWorkflow(this, jsii.String("MyWorkflow"), &WorkflowProps{
	WorkflowType: imagebuilder.WorkflowType_BUILD,
	Data: imagebuilder.WorkflowData_FromJsonObject(map[string]interface{}{
		"schemaVersion": imagebuilder.WorkflowSchemaVersion_V1_0,
		"steps": []interface{}{
			map[string]interface{}{
				"name": jsii.String("LaunchBuildInstance"),
				"action": imagebuilder.WorkflowAction_LAUNCH_INSTANCE,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"waitFor": jsii.String("ssmAgent"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("ExecuteComponents"),
				"action": imagebuilder.WorkflowAction_EXECUTE_COMPONENTS,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("CreateImage"),
				"action": imagebuilder.WorkflowAction_CREATE_IMAGE,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("TerminateInstance"),
				"action": imagebuilder.WorkflowAction_TERMINATE_INSTANCE,
				"onFailure": imagebuilder.WorkflowOnFailure_CONTINUE,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
		},
		"outputs": []interface{}{
			map[string]*string{
				"name": jsii.String("ImageId"),
				"value": jsii.String("$.stepOutputs.CreateImage.imageId"),
			},
		},
	}),
})
```

#### Workflow Data Sources

##### Inline Workflow Data

Use `WorkflowData.fromInline()` for existing YAML/JSON definitions:

```go
workflow := imagebuilder.NewWorkflow(this, jsii.String("InlineWorkflow"), &WorkflowProps{
	WorkflowType: imagebuilder.WorkflowType_TEST,
	Data: imagebuilder.WorkflowData_FromInline(jsii.String(`
	schemaVersion: 1.0
	steps:
	  - name: LaunchTestInstance
	    action: LaunchInstance
	    onFailure: Abort
	    inputs:
	      waitFor: ssmAgent
	  - name: RunTests
	    action: RunCommand
	    onFailure: Abort
	    inputs:
	      instanceId.$: "$.stepOutputs.LaunchTestInstance.instanceId"
	      commands: ['./run-tests.sh']
	  - name: TerminateTestInstance
	    action: TerminateInstance
	    onFailure: Continue
	    inputs:
	      instanceId.$: "$.stepOutputs.LaunchTestInstance.instanceId"
	`)),
})
```

##### JSON Object Workflow Data

Most developer-friendly approach using JavaScript objects:

```go
workflow := imagebuilder.NewWorkflow(this, jsii.String("JsonWorkflow"), &WorkflowProps{
	WorkflowType: imagebuilder.WorkflowType_BUILD,
	Data: imagebuilder.WorkflowData_FromJsonObject(map[string]interface{}{
		"schemaVersion": imagebuilder.WorkflowSchemaVersion_V1_0,
		"steps": []interface{}{
			map[string]interface{}{
				"name": jsii.String("LaunchBuildInstance"),
				"action": imagebuilder.WorkflowAction_LAUNCH_INSTANCE,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"waitFor": jsii.String("ssmAgent"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("ExecuteComponents"),
				"action": imagebuilder.WorkflowAction_EXECUTE_COMPONENTS,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("CreateImage"),
				"action": imagebuilder.WorkflowAction_CREATE_IMAGE,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("TerminateInstance"),
				"action": imagebuilder.WorkflowAction_TERMINATE_INSTANCE,
				"onFailure": imagebuilder.WorkflowOnFailure_CONTINUE,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
		},
		"outputs": []interface{}{
			map[string]*string{
				"name": jsii.String("ImageId"),
				"value": jsii.String("$.stepOutputs.CreateImage.imageId"),
			},
		},
	}),
})
```

##### S3 Workflow Data

For those workflows you want to upload or have uploaded to S3:

```go
// Upload a local file
workflowFromAsset := imagebuilder.NewWorkflow(this, jsii.String("AssetWorkflow"), &WorkflowProps{
	WorkflowType: imagebuilder.WorkflowType_BUILD,
	Data: imagebuilder.WorkflowData_FromAsset(this, jsii.String("WorkflowAsset"), jsii.String("./my-workflow.yml")),
})

// Reference an existing S3 object
bucket := s3.Bucket_FromBucketName(this, jsii.String("WorkflowBucket"), jsii.String("my-workflows-bucket"))
workflowFromS3 := imagebuilder.NewWorkflow(this, jsii.String("S3Workflow"), &WorkflowProps{
	WorkflowType: imagebuilder.WorkflowType_BUILD,
	Data: imagebuilder.WorkflowData_FromS3(bucket, jsii.String("workflows/my-workflow.yml")),
})
```

#### Encrypt workflow data with a KMS key

You can encrypt workflow data with a KMS key, so that only principals with access to decrypt with the key are able to
access the workflow data.

```go
workflow := imagebuilder.NewWorkflow(this, jsii.String("EncryptedWorkflow"), &WorkflowProps{
	WorkflowType: imagebuilder.WorkflowType_BUILD,
	KmsKey: kms.NewKey(this, jsii.String("WorkflowKey")),
	Data: imagebuilder.WorkflowData_FromJsonObject(map[string]interface{}{
		"schemaVersion": imagebuilder.WorkflowSchemaVersion_V1_0,
		"steps": []interface{}{
			map[string]interface{}{
				"name": jsii.String("LaunchBuildInstance"),
				"action": imagebuilder.WorkflowAction_LAUNCH_INSTANCE,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"waitFor": jsii.String("ssmAgent"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("CreateImage"),
				"action": imagebuilder.WorkflowAction_CREATE_IMAGE,
				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
			map[string]interface{}{
				"name": jsii.String("TerminateInstance"),
				"action": imagebuilder.WorkflowAction_TERMINATE_INSTANCE,
				"onFailure": imagebuilder.WorkflowOnFailure_CONTINUE,
				"inputs": map[string]*string{
					"instanceId": jsii.String("i-123"),
				},
			},
		},
		"outputs": []interface{}{
			map[string]*string{
				"name": jsii.String("ImageId"),
				"value": jsii.String("$.stepOutputs.CreateImage.imageId"),
			},
		},
	}),
})
```

#### AWS-Managed Workflows

AWS provides a collection of workflows for common scenarios:

```go
// Build workflows
buildImageWorkflow := imagebuilder.AmazonManagedWorkflow_BuildImage(this, jsii.String("BuildImage"))
buildContainerWorkflow := imagebuilder.AmazonManagedWorkflow_BuildContainer(this, jsii.String("BuildContainer"))

// Test workflows
testImageWorkflow := imagebuilder.AmazonManagedWorkflow_TestImage(this, jsii.String("TestImage"))
testContainerWorkflow := imagebuilder.AmazonManagedWorkflow_TestContainer(this, jsii.String("TestContainer"))

// Distribution workflows
distributeContainerWorkflow := imagebuilder.AmazonManagedWorkflow_DistributeContainer(this, jsii.String("DistributeContainer"))
```

### Lifecycle Policy

Lifecycle policies help you manage the retention and cleanup of Image Builder resources automatically. These policies
define rules for deprecating or deleting old image versions, managing AMI snapshots, and controlling resource costs by
removing unused images based on age, count, or other criteria.

#### Lifecycle Policy Basic Usage

Create a lifecycle policy to automatically delete old AMI images after 30 days:

```go
lifecyclePolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("MyLifecyclePolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_Days(jsii.Number(30)),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Environment": jsii.String("development"),
		},
	},
})
```

Create a lifecycle policy to keep only the 10 most recent container images:

```go
containerLifecyclePolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("ContainerLifecyclePolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_CONTAINER_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				CountFilter: &LifecyclePolicyCountFilter{
					Count: jsii.Number(10),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Application": jsii.String("web-app"),
		},
	},
})
```

#### Lifecycle Policy Resource Selection

##### Tag-Based Resource Selection

Apply lifecycle policies to images with specific tags:

```go
tagBasedPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("TagBasedPolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_Days(jsii.Number(90)),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Environment": jsii.String("staging"),
			"Team": jsii.String("backend"),
		},
	},
})
```

##### Recipe-Based Resource Selection

Apply lifecycle policies to specific image or container recipes:

```go
imageRecipe := imagebuilder.NewImageRecipe(this, jsii.String("MyImageRecipe"), &ImageRecipeProps{
	BaseImage: imagebuilder.BaseImage_*FromSsmParameterName(jsii.String("/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64")),
})

containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("MyContainerRecipe"), &ContainerRecipeProps{
	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
})

recipeBasedPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("RecipeBasedPolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				CountFilter: &LifecyclePolicyCountFilter{
					Count: jsii.Number(5),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Recipes: []IRecipeBase{
			imageRecipe,
			containerRecipe,
		},
	},
})
```

#### Lifecycle Policy Rules

##### Age-Based Rules

Delete images older than a specific time period:

```go
ageBasedPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("AgeBasedPolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
				IncludeAmis: jsii.Boolean(true),
				IncludeSnapshots: jsii.Boolean(true),
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_Days(jsii.Number(60)),
					RetainAtLeast: jsii.Number(3),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Environment": jsii.String("testing"),
		},
	},
})
```

##### Count-Based Rules

Keep only a specific number of the most recent images:

```go
countBasedPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("CountBasedPolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_CONTAINER_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				CountFilter: &LifecyclePolicyCountFilter{
					Count: jsii.Number(15),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Application": jsii.String("microservice"),
		},
	},
})
```

##### Multiple Lifecycle Rules

Implement a graduated approach with multiple actions:

```go
graduatedPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("GraduatedPolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			// First: Deprecate images after 30 days
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DEPRECATE,
				IncludeAmis: jsii.Boolean(true),
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_Days(jsii.Number(30)),
					RetainAtLeast: jsii.Number(5),
				},
			},
		},
		&LifecyclePolicyDetail{
			// Second: Disable images after 60 days
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DISABLE,
				IncludeAmis: jsii.Boolean(true),
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_*Days(jsii.Number(60)),
					RetainAtLeast: jsii.Number(3),
				},
			},
		},
		&LifecyclePolicyDetail{
			// Finally: Delete images after 90 days
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
				IncludeAmis: jsii.Boolean(true),
				IncludeSnapshots: jsii.Boolean(true),
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_*Days(jsii.Number(90)),
					RetainAtLeast: jsii.Number(1),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Environment": jsii.String("production"),
		},
	},
})
```

#### Lifecycle Policy Exclusion Rules

##### AMI Exclusion Rules

Exclude specific AMIs from lifecycle actions based on various criteria:

```go
excludeAmisPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("ExcludeAmisPolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_Days(jsii.Number(30)),
				},
			},
			ExclusionRules: &LifecyclePolicyExclusionRules{
				AmiExclusionRules: &LifecyclePolicyAmiExclusionRules{
					IsPublic: jsii.Boolean(true),
					 // Exclude public AMIs
					LastLaunched: awscdk.Duration_*Days(jsii.Number(7)),
					 // Exclude AMIs launched in last 7 days
					Regions: []*string{
						jsii.String("us-west-2"),
						jsii.String("eu-west-1"),
					},
					 // Exclude AMIs in specific regions
					SharedAccounts: []*string{
						jsii.String("123456789012"),
					},
					 // Exclude AMIs shared with specific accounts
					Tags: map[string]*string{
						"Protected": jsii.String("true"),
						"Environment": jsii.String("production"),
					},
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Team": jsii.String("infrastructure"),
		},
	},
})
```

##### Image Exclusion Rules

Exclude Image Builder images with protective tags:

```go
excludeImagesPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("ExcludeImagesPolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_CONTAINER_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				CountFilter: &LifecyclePolicyCountFilter{
					Count: jsii.Number(20),
				},
			},
			ExclusionRules: &LifecyclePolicyExclusionRules{
				ImageExclusionRules: &LifecyclePolicyImageExclusionRules{
					Tags: map[string]*string{
						"DoNotDelete": jsii.String("true"),
						"Critical": jsii.String("baseline"),
					},
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Application": jsii.String("frontend"),
		},
	},
})
```

#### Advanced Lifecycle Configuration

##### Custom Execution Roles

Provide your own IAM execution role with specific permissions:

```go
executionRole := iam.NewRole(this, jsii.String("LifecycleExecutionRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("imagebuilder.amazonaws.com")),
	ManagedPolicies: []IManagedPolicy{
		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/EC2ImageBuilderLifecycleExecutionPolicy")),
	},
})

customRolePolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("CustomRolePolicy"), &LifecyclePolicyProps{
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	ExecutionRole: executionRole,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_Days(jsii.Number(45)),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Environment": jsii.String("development"),
		},
	},
})
```

##### Lifecycle Policy Status

Control whether the lifecycle policy is active:

```go
disabledPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("DisabledPolicy"), &LifecyclePolicyProps{
	LifecyclePolicyName: jsii.String("my-disabled-policy"),
	Description: jsii.String("A lifecycle policy that is temporarily disabled"),
	Status: imagebuilder.LifecyclePolicyStatus_DISABLED,
	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
	Details: []LifecyclePolicyDetail{
		&LifecyclePolicyDetail{
			Action: &LifecyclePolicyAction{
				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
			},
			Filter: &LifecyclePolicyFilter{
				AgeFilter: &LifecyclePolicyAgeFilter{
					Age: awscdk.Duration_Days(jsii.Number(30)),
				},
			},
		},
	},
	ResourceSelection: &LifecyclePolicyResourceSelection{
		Tags: map[string]*string{
			"Environment": jsii.String("testing"),
		},
	},
	Tags: map[string]*string{
		"Owner": jsii.String("DevOps"),
		"CostCenter": jsii.String("Engineering"),
	},
})
```

##### Importing Lifecycle Policies

Reference lifecycle policies created outside CDK:

```go
// Import by name
importedByName := imagebuilder.LifecyclePolicy_FromLifecyclePolicyName(this, jsii.String("ImportedByName"), jsii.String("existing-lifecycle-policy"))

// Import by ARN
importedByArn := imagebuilder.LifecyclePolicy_FromLifecyclePolicyArn(this, jsii.String("ImportedByArn"), jsii.String("arn:aws:imagebuilder:us-east-1:123456789012:lifecycle-policy/my-policy"))

importedByName.GrantRead(lambdaRole)
importedByArn.Grant(lambdaRole, jsii.String("imagebuilder:UpdateLifecyclePolicy"))
```

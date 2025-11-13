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

package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for creating an Infrastructure Configuration resource.
//
// Example:
//   infrastructureConfiguration := imagebuilder.NewInfrastructureConfiguration(this, jsii.String("InfrastructureConfiguration"), &InfrastructureConfigurationProps{
//   	InfrastructureConfigurationName: jsii.String("test-infrastructure-configuration"),
//   	Description: jsii.String("An Infrastructure Configuration"),
//   	// Optional - instance types to use for build/test
//   	InstanceTypes: []InstanceType{
//   		ec2.InstanceType_*Of(ec2.InstanceClass_STANDARD7_INTEL, ec2.InstanceSize_LARGE),
//   		ec2.InstanceType_*Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_LARGE),
//   	},
//   	// Optional - create an instance profile with necessary permissions
//   	InstanceProfile: iam.NewInstanceProfile(this, jsii.String("InstanceProfile"), &InstanceProfileProps{
//   		InstanceProfileName: jsii.String("test-instance-profile"),
//   		Role: iam.NewRole(this, jsii.String("InstanceProfileRole"), &RoleProps{
//   			AssumedBy: iam.ServicePrincipal_FromStaticServicePrincipleName(jsii.String("ec2.amazonaws.com")),
//   			ManagedPolicies: []IManagedPolicy{
//   				iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonSSMManagedInstanceCore")),
//   				iam.ManagedPolicy_*FromAwsManagedPolicyName(jsii.String("EC2InstanceProfileForImageBuilder")),
//   			},
//   		}),
//   	}),
//   	// Use VPC network configuration
//   	Vpc: Vpc,
//   	SubnetSelection: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	SecurityGroups: []ISecurityGroup{
//   		ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("SecurityGroup"), vpc.VpcDefaultSecurityGroup),
//   	},
//   	KeyPair: ec2.KeyPair_FromKeyPairName(this, jsii.String("KeyPair"), jsii.String("imagebuilder-instance-key-pair")),
//   	TerminateInstanceOnFailure: jsii.Boolean(true),
//   	// Optional - IMDSv2 settings
//   	HttpTokens: imagebuilder.HttpTokens_REQUIRED,
//   	HttpPutResponseHopLimit: jsii.Number(1),
//   	// Optional - publish image completion messages to an SNS topic
//   	NotificationTopic: sns.Topic_FromTopicArn(this, jsii.String("Topic"), this.FormatArn(&ArnComponents{
//   		Service: jsii.String("sns"),
//   		Resource: jsii.String("image-builder-topic"),
//   	})),
//   	// Optional - log settings. Logging is enabled by default
//   	Logging: &InfrastructureConfigurationLogging{
//   		S3Bucket: s3.Bucket_FromBucketName(this, jsii.String("LogBucket"), fmt.Sprintf("imagebuilder-logging-%v", awscdk.Aws_ACCOUNT_ID())),
//   		S3KeyPrefix: jsii.String("imagebuilder-logs"),
//   	},
//   	// Optional - host placement settings
//   	Ec2InstanceAvailabilityZone: awscdk.*stack_*Of(this).availabilityZones[jsii.Number(0)],
//   	Ec2InstanceHostId: dedicatedHost.AttrHostId,
//   	Ec2InstanceTenancy: imagebuilder.Tenancy_HOST,
//   	ResourceTags: map[string]*string{
//   		"Environment": jsii.String("production"),
//   	},
//   })
//
// Experimental.
type InfrastructureConfigurationProps struct {
	// The description of the infrastructure configuration.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The availability zone to place Image Builder build and test EC2 instances.
	// Default: EC2 will select a random zone.
	//
	// Experimental.
	Ec2InstanceAvailabilityZone *string `field:"optional" json:"ec2InstanceAvailabilityZone" yaml:"ec2InstanceAvailabilityZone"`
	// The ID of the Dedicated Host on which build and test instances run.
	//
	// This only applies if the instance tenancy is
	// `host`. This cannot be used with the `ec2InstanceHostResourceGroupArn` parameter.
	// Default: None.
	//
	// Experimental.
	Ec2InstanceHostId *string `field:"optional" json:"ec2InstanceHostId" yaml:"ec2InstanceHostId"`
	// The ARN of the host resource group on which build and test instances run.
	//
	// This only applies if the instance tenancy
	// is `host`. This cannot be used with the `ec2InstanceHostId` parameter.
	// Default: None.
	//
	// Experimental.
	Ec2InstanceHostResourceGroupArn *string `field:"optional" json:"ec2InstanceHostResourceGroupArn" yaml:"ec2InstanceHostResourceGroupArn"`
	// The tenancy of the instance.
	//
	// Dedicated tenancy runs instances on single-tenant hardware, while host tenancy runs
	// instances on a dedicated host. Shared tenancy is used by default.
	// Default: Tenancy.DEFAULT
	//
	// Experimental.
	Ec2InstanceTenancy Tenancy `field:"optional" json:"ec2InstanceTenancy" yaml:"ec2InstanceTenancy"`
	// The maximum number of hops that an instance metadata request can traverse to reach its destination.
	//
	// By default,
	// this is set to 2.
	// Default: 2.
	//
	// Experimental.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Indicates whether a signed token header is required for instance metadata retrieval requests.
	//
	// By default, this is
	// set to `required` to require IMDSv2 on build and test EC2 instances.
	// Default: HttpTokens.REQUIRED
	//
	// Experimental.
	HttpTokens HttpTokens `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// The name of the infrastructure configuration.
	//
	// This name must be normalized by transforming all alphabetical
	// characters to lowercase, and replacing all spaces and underscores with hyphens.
	// Default: A name is generated.
	//
	// Experimental.
	InfrastructureConfigurationName *string `field:"optional" json:"infrastructureConfigurationName" yaml:"infrastructureConfigurationName"`
	// The instance profile to associate with the instance used to customize the AMI.
	//
	// By default, an instance profile and role will be created with minimal permissions needed to build the image,
	// attached to the EC2 instance.
	//
	// If an S3 logging bucket and key prefix is provided, an IAM inline policy will be attached to the instance profile's
	// role, allowing s3:PutObject permissions on the bucket.
	// Default: An instance profile will be generated.
	//
	// Experimental.
	InstanceProfile awsiam.IInstanceProfile `field:"optional" json:"instanceProfile" yaml:"instanceProfile"`
	// The instance types to launch build and test EC2 instances with.
	// Default: Image Builder will choose from a default set of instance types compatible with the AMI.
	//
	// Experimental.
	InstanceTypes *[]awsec2.InstanceType `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The key pair used to connect to the build and test EC2 instances.
	//
	// The key pair can be used to log into the build
	// or test instances for troubleshooting any failures.
	// Default: None.
	//
	// Experimental.
	KeyPair awsec2.IKeyPair `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The log settings for detailed build logging.
	// Default: None.
	//
	// Experimental.
	Logging *InfrastructureConfigurationLogging `field:"optional" json:"logging" yaml:"logging"`
	// The SNS topic on which notifications are sent when an image build completes.
	// Default: No notifications are sent.
	//
	// Experimental.
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
	// The additional tags to assign to the Amazon EC2 instance that Image Builder launches during the build process.
	// Default: None.
	//
	// Experimental.
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// An IAM role to associate with the instance profile used by Image Builder.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	// Note: You can provide an instanceProfile or a role, but not both.
	//
	// Example:
	//   role := iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	//   	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	//   })
	//
	// Default: A role will automatically be created, it can be accessed via the `role` property.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The security groups to associate with the instance used to customize the AMI.
	// Default: The default security group for the VPC will be used.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Select which subnet to place the instance used to customize the AMI.
	//
	// The first subnet that is selected will be used.
	// You must specify the VPC to customize the subnet selection.
	// Default: The first subnet selected from the provided VPC will be used.
	//
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The tags to apply to the infrastructure configuration.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to terminate the EC2 instance when the build or test workflow fails.
	// Default: true.
	//
	// Experimental.
	TerminateInstanceOnFailure *bool `field:"optional" json:"terminateInstanceOnFailure" yaml:"terminateInstanceOnFailure"`
	// The VPC to place the instance used to customize the AMI.
	// Default: The default VPC will be used.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}


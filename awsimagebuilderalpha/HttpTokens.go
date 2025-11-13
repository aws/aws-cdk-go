package awsimagebuilderalpha


// Indicates whether a signed token header is required for instance metadata retrieval requests.
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
// See: https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_InstanceMetadataOptions.html#imagebuilder-Type-InstanceMetadataOptions-httpTokens
//
// Experimental.
type HttpTokens string

const (
	// Allows retrieval of instance metadata with or without a signed token header in the request.
	// Experimental.
	HttpTokens_OPTIONAL HttpTokens = "OPTIONAL"
	// Requires a signed token header in instance metadata retrieval requests.
	// Experimental.
	HttpTokens_REQUIRED HttpTokens = "REQUIRED"
)


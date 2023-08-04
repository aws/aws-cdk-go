package awscdk


// Configuration properties for DefaultStackSynthesizer.
//
// Example:
//   NewMyStack(app, jsii.String("MyStack"), &stackProps{
//   	Synthesizer: awscdk.NewDefaultStackSynthesizer(&DefaultStackSynthesizerProps{
//   		FileAssetsBucketName: jsii.String("my-orgs-asset-bucket"),
//   	}),
//   })
//
type DefaultStackSynthesizerProps struct {
	// Bootstrap stack version SSM parameter.
	//
	// The placeholder `${Qualifier}` will be replaced with the value of qualifier.
	// Default: DefaultStackSynthesizer.DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER
	//
	BootstrapStackVersionSsmParameter *string `field:"optional" json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// bucketPrefix to use while storing S3 Assets.
	// Default: - DefaultStackSynthesizer.DEFAULT_FILE_ASSET_PREFIX
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The role CloudFormation will assume when deploying the Stack.
	//
	// You must supply this if you have given a non-standard name to the execution role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_CLOUDFORMATION_ROLE_ARN
	//
	CloudFormationExecutionRole *string `field:"optional" json:"cloudFormationExecutionRole" yaml:"cloudFormationExecutionRole"`
	// The role to assume to initiate a deployment in this environment.
	//
	// You must supply this if you have given a non-standard name to the publishing role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_DEPLOY_ROLE_ARN
	//
	DeployRoleArn *string `field:"optional" json:"deployRoleArn" yaml:"deployRoleArn"`
	// External ID to use when assuming role for cloudformation deployments.
	// Default: - No external ID.
	//
	DeployRoleExternalId *string `field:"optional" json:"deployRoleExternalId" yaml:"deployRoleExternalId"`
	// A prefix to use while tagging and uploading Docker images to ECR.
	//
	// This does not add any separators - the source hash will be appended to
	// this string directly.
	// Default: - DefaultStackSynthesizer.DEFAULT_DOCKER_ASSET_PREFIX
	//
	DockerTagPrefix *string `field:"optional" json:"dockerTagPrefix" yaml:"dockerTagPrefix"`
	// External ID to use when assuming role for file asset publishing.
	// Default: - No external ID.
	//
	FileAssetPublishingExternalId *string `field:"optional" json:"fileAssetPublishingExternalId" yaml:"fileAssetPublishingExternalId"`
	// The role to use to publish file assets to the S3 bucket in this environment.
	//
	// You must supply this if you have given a non-standard name to the publishing role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN
	//
	FileAssetPublishingRoleArn *string `field:"optional" json:"fileAssetPublishingRoleArn" yaml:"fileAssetPublishingRoleArn"`
	// Name of the S3 bucket to hold file assets.
	//
	// You must supply this if you have given a non-standard name to the staging bucket.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_FILE_ASSETS_BUCKET_NAME
	//
	FileAssetsBucketName *string `field:"optional" json:"fileAssetsBucketName" yaml:"fileAssetsBucketName"`
	// Whether to add a Rule to the stack template verifying the bootstrap stack version.
	//
	// This generally should be left set to `true`, unless you explicitly
	// want to be able to deploy to an unbootstrapped environment.
	// Default: true.
	//
	GenerateBootstrapVersionRule *bool `field:"optional" json:"generateBootstrapVersionRule" yaml:"generateBootstrapVersionRule"`
	// External ID to use when assuming role for image asset publishing.
	// Default: - No external ID.
	//
	ImageAssetPublishingExternalId *string `field:"optional" json:"imageAssetPublishingExternalId" yaml:"imageAssetPublishingExternalId"`
	// The role to use to publish image assets to the ECR repository in this environment.
	//
	// You must supply this if you have given a non-standard name to the publishing role.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN
	//
	ImageAssetPublishingRoleArn *string `field:"optional" json:"imageAssetPublishingRoleArn" yaml:"imageAssetPublishingRoleArn"`
	// Name of the ECR repository to hold Docker Image assets.
	//
	// You must supply this if you have given a non-standard name to the ECR repository.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME
	//
	ImageAssetsRepositoryName *string `field:"optional" json:"imageAssetsRepositoryName" yaml:"imageAssetsRepositoryName"`
	// The role to use to look up values from the target AWS account during synthesis.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// External ID to use when assuming lookup role.
	// Default: - No external ID.
	//
	LookupRoleExternalId *string `field:"optional" json:"lookupRoleExternalId" yaml:"lookupRoleExternalId"`
	// Qualifier to disambiguate multiple environments in the same account.
	//
	// You can use this and leave the other naming properties empty if you have deployed
	// the bootstrap environment with standard names but only different qualifiers.
	// Default: - Value of context key '@aws-cdk/core:bootstrapQualifier' if set, otherwise `DefaultStackSynthesizer.DEFAULT_QUALIFIER`
	//
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// Use the bootstrapped lookup role for (read-only) stack operations.
	//
	// Use the lookup role when performing a `cdk diff`. If set to `false`, the
	// `deploy role` credentials will be used to perform a `cdk diff`.
	//
	// Requires bootstrap stack version 8.
	// Default: true.
	//
	UseLookupRoleForStackOperations *bool `field:"optional" json:"useLookupRoleForStackOperations" yaml:"useLookupRoleForStackOperations"`
}


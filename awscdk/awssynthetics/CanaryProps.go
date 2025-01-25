package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for a canary.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &CanaryProps{
//   	Schedule: synthetics.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(5))),
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Code: synthetics.Code_FromAsset(path.join(__dirname, jsii.String("canary"))),
//   		Handler: jsii.String("index.handler"),
//   	}),
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_6_2(),
//   	Memory: cdk.Size_Mebibytes(jsii.Number(1024)),
//   })
//
type CanaryProps struct {
	// Specify the runtime version to use for the canary.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html
	//
	Runtime Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// The type of test that you want your canary to run.
	//
	// Use `Test.custom()` to specify the test to run.
	Test Test `field:"required" json:"test" yaml:"test"`
	// Specifies whether this canary is to use active AWS X-Ray tracing when it runs.
	//
	// Active tracing enables this canary run to be displayed in the ServiceLens and X-Ray service maps even if the
	// canary does not hit an endpoint that has X-Ray tracing enabled. Using X-Ray tracing incurs charges.
	//
	// You can enable active tracing only for canaries that use version `syn-nodejs-2.0` or later for their canary runtime.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_tracing.html
	//
	// Default: false.
	//
	ActiveTracing *bool `field:"optional" json:"activeTracing" yaml:"activeTracing"`
	// Canary Artifacts in S3 encryption mode.
	//
	// Artifact encryption is only supported for canaries that use Synthetics runtime
	// version `syn-nodejs-puppeteer-3.3` or later.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_artifact_encryption.html
	//
	// Default: - Artifacts are encrypted at rest using an AWS managed key. `ArtifactsEncryptionMode.KMS` is set if you specify `artifactS3KmsKey`.
	//
	ArtifactS3EncryptionMode ArtifactsEncryptionMode `field:"optional" json:"artifactS3EncryptionMode" yaml:"artifactS3EncryptionMode"`
	// The KMS key used to encrypt canary artifacts.
	// Default: - no kms key if `artifactS3EncryptionMode` is set to `S3_MANAGED`. A key will be created if one is not provided and `artifactS3EncryptionMode` is set to `KMS`.
	//
	ArtifactS3KmsKey awskms.IKey `field:"optional" json:"artifactS3KmsKey" yaml:"artifactS3KmsKey"`
	// Lifecycle rules for the generated canary artifact bucket.
	//
	// Has no effect
	// if a bucket is passed to `artifactsBucketLocation`. If you pass a bucket
	// to `artifactsBucketLocation`, you can add lifecycle rules to the bucket
	// itself.
	// Default: - no rules applied to the generated bucket.
	//
	ArtifactsBucketLifecycleRules *[]*awss3.LifecycleRule `field:"optional" json:"artifactsBucketLifecycleRules" yaml:"artifactsBucketLifecycleRules"`
	// The s3 location that stores the data of the canary runs.
	// Default: - A new s3 bucket will be created without a prefix.
	//
	ArtifactsBucketLocation *ArtifactsBucketLocation `field:"optional" json:"artifactsBucketLocation" yaml:"artifactsBucketLocation"`
	// The name of the canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from
	// other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary name. The canary name
	// makes up part of the canary ARN, which is included in outbound calls over the internet.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html
	//
	// Default: - A unique name will be generated from the construct ID.
	//
	CanaryName *string `field:"optional" json:"canaryName" yaml:"canaryName"`
	// Specify the underlying resources to be cleaned up when the canary is deleted.
	//
	// Using `Cleanup.LAMBDA` will create a Custom Resource to achieve this.
	// Default: Cleanup.NOTHING
	//
	// Deprecated: use provisionedResourceCleanup.
	Cleanup Cleanup `field:"optional" json:"cleanup" yaml:"cleanup"`
	// Key-value pairs that the Synthetics caches and makes available for your canary scripts.
	//
	// Use environment variables
	// to apply configuration changes, such as test and production environment configurations, without changing your
	// Canary script source code.
	// Default: - No environment variables.
	//
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// How many days should failed runs be retained.
	// Default: Duration.days(31)
	//
	FailureRetentionPeriod awscdk.Duration `field:"optional" json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// The maximum amount of memory that the canary can use while running.
	//
	// This value must be a multiple of 64 Mib.
	// The range is 960 MiB to 3008 MiB.
	// Default: Size.mebibytes(1024)
	//
	Memory awscdk.Size `field:"optional" json:"memory" yaml:"memory"`
	// Whether to also delete the Lambda functions and layers used by this canary when the canary is deleted.
	// Default: undefined - the default behavior is to not delete the Lambda functions and layers.
	//
	ProvisionedResourceCleanup *bool `field:"optional" json:"provisionedResourceCleanup" yaml:"provisionedResourceCleanup"`
	// Canary execution role.
	//
	// This is the role that will be assumed by the canary upon execution.
	// It controls the permissions that the canary will have. The role must
	// be assumable by the AWS Lambda service principal.
	//
	// If not supplied, a role will be created with all the required permissions.
	// If you provide a Role, you must add the required permissions.
	// See: required permissions: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-executionrolearn
	//
	// Default: - A unique role will be generated for this canary.
	// You can add permissions to roles by calling 'addToRolePolicy'.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Specify the schedule for how often the canary runs.
	//
	// For example, if you set `schedule` to `rate(10 minutes)`, then the canary will run every 10 minutes.
	// You can set the schedule with `Schedule.rate(Duration)` (recommended) or you can specify an expression using `Schedule.expression()`.
	// Default: 'rate(5 minutes)'.
	//
	Schedule Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// The list of security groups to associate with the canary's network interfaces.
	//
	// You must provide `vpc` when using this prop.
	// Default: - If the canary is placed within a VPC and a security group is
	// not specified a dedicated security group will be created for this canary.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Whether or not the canary should start after creation.
	// Default: true.
	//
	StartAfterCreation *bool `field:"optional" json:"startAfterCreation" yaml:"startAfterCreation"`
	// How many days should successful runs be retained.
	// Default: Duration.days(31)
	//
	SuccessRetentionPeriod awscdk.Duration `field:"optional" json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// How long the canary is allowed to run before it must stop.
	//
	// You can't set this time to be longer than the frequency of the runs of this canary.
	//
	// The minimum allowed value is 3 seconds.
	// The maximum allowed value is 840 seconds (14 minutes).
	// Default: - the frequency of the canary is used as this value, up to a maximum of 900 seconds.
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// How long the canary will be in a 'RUNNING' state.
	//
	// For example, if you set `timeToLive` to be 1 hour and `schedule` to be `rate(10 minutes)`,
	// your canary will run at 10 minute intervals for an hour, for a total of 6 times.
	// Default: - no limit.
	//
	TimeToLive awscdk.Duration `field:"optional" json:"timeToLive" yaml:"timeToLive"`
	// The VPC where this canary is run.
	//
	// Specify this if the canary needs to access resources in a VPC.
	// Default: - Not in VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// You must provide `vpc` when using this prop.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}


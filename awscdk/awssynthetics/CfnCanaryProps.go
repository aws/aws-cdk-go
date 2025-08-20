package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCanary`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCanaryProps := &CfnCanaryProps{
//   	ArtifactS3Location: jsii.String("artifactS3Location"),
//   	Code: &CodeProperty{
//   		Handler: jsii.String("handler"),
//
//   		// the properties below are optional
//   		Dependencies: []interface{}{
//   			&DependencyProperty{
//   				Reference: jsii.String("reference"),
//
//   				// the properties below are optional
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		Script: jsii.String("script"),
//   		SourceLocationArn: jsii.String("sourceLocationArn"),
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	RuntimeVersion: jsii.String("runtimeVersion"),
//   	Schedule: &ScheduleProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		DurationInSeconds: jsii.String("durationInSeconds"),
//   		RetryConfig: &RetryConfigProperty{
//   			MaxRetries: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ArtifactConfig: &ArtifactConfigProperty{
//   		S3Encryption: &S3EncryptionProperty{
//   			EncryptionMode: jsii.String("encryptionMode"),
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	DeleteLambdaResourcesOnCanaryDeletion: jsii.Boolean(false),
//   	DryRunAndUpdate: jsii.Boolean(false),
//   	FailureRetentionPeriod: jsii.Number(123),
//   	ProvisionedResourceCleanup: jsii.String("provisionedResourceCleanup"),
//   	ResourcesToReplicateTags: []*string{
//   		jsii.String("resourcesToReplicateTags"),
//   	},
//   	RunConfig: &RunConfigProperty{
//   		ActiveTracing: jsii.Boolean(false),
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		EphemeralStorage: jsii.Number(123),
//   		MemoryInMb: jsii.Number(123),
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	StartCanaryAfterCreation: jsii.Boolean(false),
//   	SuccessRetentionPeriod: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VisualReference: &VisualReferenceProperty{
//   		BaseCanaryRunId: jsii.String("baseCanaryRunId"),
//
//   		// the properties below are optional
//   		BaseScreenshots: []interface{}{
//   			&BaseScreenshotProperty{
//   				ScreenshotName: jsii.String("screenshotName"),
//
//   				// the properties below are optional
//   				IgnoreCoordinates: []*string{
//   					jsii.String("ignoreCoordinates"),
//   				},
//   			},
//   		},
//   	},
//   	VpcConfig: &VPCConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		Ipv6AllowedForDualStack: jsii.Boolean(false),
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html
//
type CfnCanaryProps struct {
	// The location in Amazon S3 where Synthetics stores artifacts from the runs of this canary.
	//
	// Artifacts include the log file, screenshots, and HAR files. Specify the full location path, including `s3://` at the beginning of the path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-artifacts3location
	//
	ArtifactS3Location *string `field:"required" json:"artifactS3Location" yaml:"artifactS3Location"`
	// Use this structure to input your script code for the canary.
	//
	// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-code
	//
	Code interface{} `field:"required" json:"code" yaml:"code"`
	// The ARN of the IAM role to be used to run the canary.
	//
	// This role must already exist, and must include `lambda.amazonaws.com` as a principal in the trust policy. The role must also have the following permissions:
	//
	// - `s3:PutObject`
	// - `s3:GetBucketLocation`
	// - `s3:ListAllMyBuckets`
	// - `cloudwatch:PutMetricData`
	// - `logs:CreateLogGroup`
	// - `logs:CreateLogStream`
	// - `logs:PutLogEvents`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name for this canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the canary ARN, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the runtime version to use for the canary.
	//
	// For more information about runtime versions, see [Canary Runtime Versions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runtimeversion
	//
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// A structure that contains information about how often the canary is to run, and when these runs are to stop.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-schedule
	//
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-artifactconfig
	//
	ArtifactConfig interface{} `field:"optional" json:"artifactConfig" yaml:"artifactConfig"`
	// Deletes associated lambda resources created by Synthetics if set to True.
	//
	// Default is False.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-deletelambdaresourcesoncanarydeletion
	//
	// Deprecated: this property has been deprecated.
	DeleteLambdaResourcesOnCanaryDeletion interface{} `field:"optional" json:"deleteLambdaResourcesOnCanaryDeletion" yaml:"deleteLambdaResourcesOnCanaryDeletion"`
	// Specifies whether to perform a dry run before updating the canary.
	//
	// If set to `true` , CloudFormation will execute a dry run to validate the changes before applying them to the canary. If the dry run succeeds, the canary will be updated with the changes. If the dry run fails, the CloudFormation deployment will fail with the dry run’s failure reason.
	//
	// If set to `false` or omitted, the canary will be updated directly without first performing a dry run. The default value is `false` .
	//
	// For more information, see [Performing safe canary updates](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/performing-safe-canary-upgrades.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-dryrunandupdate
	//
	DryRunAndUpdate interface{} `field:"optional" json:"dryRunAndUpdate" yaml:"dryRunAndUpdate"`
	// The number of days to retain data about failed runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	//
	// This setting affects the range of information returned by [GetCanaryRuns](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_GetCanaryRuns.html) , as well as the range of information displayed in the Synthetics console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-failureretentionperiod
	//
	FailureRetentionPeriod *float64 `field:"optional" json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// Specifies whether to also delete the Lambda functions and layers used by this canary when the canary is deleted.
	//
	// If it is `AUTOMATIC` , the Lambda functions and layers will be deleted when the canary is deleted.
	//
	// If the value of this parameter is `OFF` , then the value of the `DeleteLambda` parameter of the [DeleteCanary](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DeleteCanary.html) operation determines whether the Lambda functions and layers will be deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-provisionedresourcecleanup
	//
	ProvisionedResourceCleanup *string `field:"optional" json:"provisionedResourceCleanup" yaml:"provisionedResourceCleanup"`
	// To have the tags that you apply to this canary also be applied to the Lambda function that the canary uses, specify this property with the value `lambda-function` .
	//
	// If you do this, CloudWatch Synthetics will keep the tags of the canary and the Lambda function synchronized. Any future changes you make to the canary's tags will also be applied to the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-resourcestoreplicatetags
	//
	ResourcesToReplicateTags *[]*string `field:"optional" json:"resourcesToReplicateTags" yaml:"resourcesToReplicateTags"`
	// A structure that contains input information for a canary run.
	//
	// If you omit this structure, the frequency of the canary is used as canary's timeout value, up to a maximum of 900 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runconfig
	//
	RunConfig interface{} `field:"optional" json:"runConfig" yaml:"runConfig"`
	// Specify TRUE to have the canary start making runs immediately after it is created.
	//
	// A canary that you create using CloudFormation can't be used to monitor the CloudFormation stack that creates the canary or to roll back that stack if there is a failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-startcanaryaftercreation
	//
	StartCanaryAfterCreation interface{} `field:"optional" json:"startCanaryAfterCreation" yaml:"startCanaryAfterCreation"`
	// The number of days to retain data about successful runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	//
	// This setting affects the range of information returned by [GetCanaryRuns](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_GetCanaryRuns.html) , as well as the range of information displayed in the Synthetics console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-successretentionperiod
	//
	SuccessRetentionPeriod *float64 `field:"optional" json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// The list of key-value pairs that are associated with the canary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// If this canary performs visual monitoring by comparing screenshots, this structure contains the ID of the canary run to use as the baseline for screenshots, and the coordinates of any parts of the screen to ignore during the visual monitoring comparison.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-visualreference
	//
	VisualReference interface{} `field:"optional" json:"visualReference" yaml:"visualReference"`
	// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
	//
	// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


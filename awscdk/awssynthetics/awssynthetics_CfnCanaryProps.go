package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCanary`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCanaryProps := &cfnCanaryProps{
//   	artifactS3Location: jsii.String("artifactS3Location"),
//   	code: &codeProperty{
//   		handler: jsii.String("handler"),
//
//   		// the properties below are optional
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		script: jsii.String("script"),
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	name: jsii.String("name"),
//   	runtimeVersion: jsii.String("runtimeVersion"),
//   	schedule: &scheduleProperty{
//   		expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		durationInSeconds: jsii.String("durationInSeconds"),
//   	},
//   	startCanaryAfterCreation: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	artifactConfig: &artifactConfigProperty{
//   		s3Encryption: &s3EncryptionProperty{
//   			encryptionMode: jsii.String("encryptionMode"),
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	deleteLambdaResourcesOnCanaryDeletion: jsii.Boolean(false),
//   	failureRetentionPeriod: jsii.Number(123),
//   	runConfig: &runConfigProperty{
//   		activeTracing: jsii.Boolean(false),
//   		environmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		memoryInMb: jsii.Number(123),
//   		timeoutInSeconds: jsii.Number(123),
//   	},
//   	successRetentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	visualReference: &visualReferenceProperty{
//   		baseCanaryRunId: jsii.String("baseCanaryRunId"),
//
//   		// the properties below are optional
//   		baseScreenshots: []interface{}{
//   			&baseScreenshotProperty{
//   				screenshotName: jsii.String("screenshotName"),
//
//   				// the properties below are optional
//   				ignoreCoordinates: []*string{
//   					jsii.String("ignoreCoordinates"),
//   				},
//   			},
//   		},
//   	},
//   	vpcConfig: &vPCConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		vpcId: jsii.String("vpcId"),
//   	},
//   }
//
type CfnCanaryProps struct {
	// The location in Amazon S3 where Synthetics stores artifacts from the runs of this canary.
	//
	// Artifacts include the log file, screenshots, and HAR files. Specify the full location path, including `s3://` at the beginning of the path.
	ArtifactS3Location *string `field:"required" json:"artifactS3Location" yaml:"artifactS3Location"`
	// Use this structure to input your script code for the canary.
	//
	// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
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
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name for this canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the canary ARN, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the runtime version to use for the canary.
	//
	// For more information about runtime versions, see [Canary Runtime Versions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) .
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// A structure that contains information about how often the canary is to run, and when these runs are to stop.
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Specify TRUE to have the canary start making runs immediately after it is created.
	//
	// A canary that you create using CloudFormation can't be used to monitor the CloudFormation stack that creates the canary or to roll back that stack if there is a failure.
	StartCanaryAfterCreation interface{} `field:"required" json:"startCanaryAfterCreation" yaml:"startCanaryAfterCreation"`
	// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
	ArtifactConfig interface{} `field:"optional" json:"artifactConfig" yaml:"artifactConfig"`
	// Specifies whether AWS CloudFormation is to also delete the Lambda functions and layers used by this canary, when the canary is deleted.
	//
	// The default is false.
	DeleteLambdaResourcesOnCanaryDeletion interface{} `field:"optional" json:"deleteLambdaResourcesOnCanaryDeletion" yaml:"deleteLambdaResourcesOnCanaryDeletion"`
	// The number of days to retain data about failed runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod *float64 `field:"optional" json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// A structure that contains input information for a canary run.
	//
	// If you omit this structure, the frequency of the canary is used as canary's timeout value, up to a maximum of 900 seconds.
	RunConfig interface{} `field:"optional" json:"runConfig" yaml:"runConfig"`
	// The number of days to retain data about successful runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod *float64 `field:"optional" json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// The list of key-value pairs that are associated with the canary.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// If this canary performs visual monitoring by comparing screenshots, this structure contains the ID of the canary run to use as the baseline for screenshots, and the coordinates of any parts of the screen to ignore during the visual monitoring comparison.
	VisualReference interface{} `field:"optional" json:"visualReference" yaml:"visualReference"`
	// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
	//
	// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


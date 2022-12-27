package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for an AWS CloudTrail trail.
//
// Example:
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &trailProps{
//   	// ...
//   	managementEvents: cloudtrail.readWriteType_READ_ONLY,
//   })
//
type TrailProps struct {
	// The Amazon S3 bucket.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Log Group to which CloudTrail to push logs to.
	//
	// Ignored if sendToCloudWatchLogs is set to false.
	CloudWatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudWatchLogGroup" yaml:"cloudWatchLogGroup"`
	// How long to retain logs in CloudWatchLogs.
	//
	// Ignored if sendToCloudWatchLogs is false or if cloudWatchLogGroup is set.
	CloudWatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudWatchLogsRetention" yaml:"cloudWatchLogsRetention"`
	// To determine whether a log file was modified, deleted, or unchanged after CloudTrail delivered it, you can use CloudTrail log file integrity validation.
	//
	// This feature is built using industry standard algorithms: SHA-256 for hashing and SHA-256 with RSA for digital signing.
	// This makes it computationally infeasible to modify, delete or forge CloudTrail log files without detection.
	// You can use the AWS CLI to validate the files in the location where CloudTrail delivered them.
	EnableFileValidation *bool `field:"optional" json:"enableFileValidation" yaml:"enableFileValidation"`
	// The AWS Key Management Service (AWS KMS) key ID that you want to use to encrypt CloudTrail logs.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// For most services, events are recorded in the region where the action occurred.
	//
	// For global services such as AWS Identity and Access Management (IAM), AWS STS, Amazon CloudFront, and Route 53,
	// events are delivered to any trail that includes global services, and are logged as occurring in US East (N. Virginia) Region.
	IncludeGlobalServiceEvents *bool `field:"optional" json:"includeGlobalServiceEvents" yaml:"includeGlobalServiceEvents"`
	// A JSON string that contains the insight types you want to log on a trail.
	InsightTypes *[]InsightType `field:"optional" json:"insightTypes" yaml:"insightTypes"`
	// Whether or not this trail delivers log files from multiple regions to a single S3 bucket for a single account.
	IsMultiRegionTrail *bool `field:"optional" json:"isMultiRegionTrail" yaml:"isMultiRegionTrail"`
	// Specifies whether the trail is applied to all accounts in an organization in AWS Organizations, or only for the current AWS account.
	//
	// If this is set to true then the current account _must_ be the management account. If it is not, then CloudFormation will throw an error.
	//
	// If this is set to true and the current account is a management account for an organization in AWS Organizations, the trail will be created in all AWS accounts that belong to the organization.
	// If this is set to false, the trail will remain in the current AWS account but be deleted from all member accounts in the organization.
	IsOrganizationTrail *bool `field:"optional" json:"isOrganizationTrail" yaml:"isOrganizationTrail"`
	// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
	//
	// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
	//
	// This method sets the management configuration for this trail.
	//
	// Management events provide insight into management operations that are performed on resources in your AWS account.
	// These are also known as control plane operations.
	// Management events can also include non-API events that occur in your account.
	// For example, when a user logs in to your account, CloudTrail logs the ConsoleLogin event.
	ManagementEvents ReadWriteType `field:"optional" json:"managementEvents" yaml:"managementEvents"`
	// An Amazon S3 object key prefix that precedes the name of all log files.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// If CloudTrail pushes logs to CloudWatch Logs in addition to S3.
	//
	// Disabled for cost out of the box.
	SendToCloudWatchLogs *bool `field:"optional" json:"sendToCloudWatchLogs" yaml:"sendToCloudWatchLogs"`
	// SNS topic that is notified when new log files are published.
	SnsTopic awssns.ITopic `field:"optional" json:"snsTopic" yaml:"snsTopic"`
	// The name of the trail.
	//
	// We recommend customers do not set an explicit name.
	TrailName *string `field:"optional" json:"trailName" yaml:"trailName"`
}


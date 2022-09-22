package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTrail`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrailProps := &cfnTrailProps{
//   	isLogging: jsii.Boolean(false),
//   	s3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	cloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	enableLogFileValidation: jsii.Boolean(false),
//   	eventSelectors: []interface{}{
//   		&eventSelectorProperty{
//   			dataResources: []interface{}{
//   				&dataResourceProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			excludeManagementEventSources: []*string{
//   				jsii.String("excludeManagementEventSources"),
//   			},
//   			includeManagementEvents: jsii.Boolean(false),
//   			readWriteType: jsii.String("readWriteType"),
//   		},
//   	},
//   	includeGlobalServiceEvents: jsii.Boolean(false),
//   	insightSelectors: []interface{}{
//   		&insightSelectorProperty{
//   			insightType: jsii.String("insightType"),
//   		},
//   	},
//   	isMultiRegionTrail: jsii.Boolean(false),
//   	isOrganizationTrail: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	snsTopicName: jsii.String("snsTopicName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trailName: jsii.String("trailName"),
//   }
//
type CfnTrailProps struct {
	// Whether the CloudTrail trail is currently logging AWS API calls.
	IsLogging interface{} `field:"required" json:"isLogging" yaml:"isLogging"`
	// Specifies the name of the Amazon S3 bucket designated for publishing log files.
	//
	// See [Amazon S3 Bucket Naming Requirements](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html) .
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs are delivered.
	//
	// Not required unless you specify `CloudWatchLogsRoleArn` .
	CloudWatchLogsLogGroupArn *string `field:"optional" json:"cloudWatchLogsLogGroupArn" yaml:"cloudWatchLogsLogGroupArn"`
	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
	CloudWatchLogsRoleArn *string `field:"optional" json:"cloudWatchLogsRoleArn" yaml:"cloudWatchLogsRoleArn"`
	// Specifies whether log file validation is enabled. The default is false.
	//
	// > When you disable log file integrity validation, the chain of digest files is broken after one hour. CloudTrail does not create digest files for log files that were delivered during a period in which log file integrity validation was disabled. For example, if you enable log file integrity validation at noon on January 1, disable it at noon on January 2, and re-enable it at noon on January 10, digest files will not be created for the log files delivered from noon on January 2 to noon on January 10. The same applies whenever you stop CloudTrail logging or delete a trail.
	EnableLogFileValidation interface{} `field:"optional" json:"enableLogFileValidation" yaml:"enableLogFileValidation"`
	// Use event selectors to further specify the management and data event settings for your trail.
	//
	// By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event.
	//
	// You can configure up to five event selectors for a trail.
	//
	// You cannot apply both event selectors and advanced event selectors to a trail.
	EventSelectors interface{} `field:"optional" json:"eventSelectors" yaml:"eventSelectors"`
	// Specifies whether the trail is publishing events from global services such as IAM to the log files.
	IncludeGlobalServiceEvents interface{} `field:"optional" json:"includeGlobalServiceEvents" yaml:"includeGlobalServiceEvents"`
	// `AWS::CloudTrail::Trail.InsightSelectors`.
	InsightSelectors interface{} `field:"optional" json:"insightSelectors" yaml:"insightSelectors"`
	// Specifies whether the trail applies only to the current region or to all regions.
	//
	// The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
	IsMultiRegionTrail interface{} `field:"optional" json:"isMultiRegionTrail" yaml:"isMultiRegionTrail"`
	// Specifies whether the trail is applied to all accounts in an organization in AWS Organizations , or only for the current AWS account .
	//
	// The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the management account for an organization in AWS Organizations . If the trail is not an organization trail and this is set to `true` , the trail will be created in all AWS accounts that belong to the organization. If the trail is an organization trail and this is set to `false` , the trail will remain in the current AWS account but be deleted from all member accounts in the organization.
	IsOrganizationTrail interface{} `field:"optional" json:"isOrganizationTrail" yaml:"isOrganizationTrail"`
	// Specifies the AWS KMS key ID to use to encrypt the logs delivered by CloudTrail.
	//
	// The value can be an alias name prefixed by "alias/", a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// CloudTrail also supports AWS KMS multi-Region keys. For more information about multi-Region keys, see [Using multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// Examples:
	//
	// - alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	// - 12345678-1234-1234-1234-123456789012.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.
	//
	// For more information, see [Finding Your CloudTrail Log Files](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html) . The maximum length is 200 characters.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic defined for notification of log file delivery.
	//
	// The maximum length is 256 characters.
	SnsTopicName *string `field:"optional" json:"snsTopicName" yaml:"snsTopicName"`
	// A custom set of tags (key-value pairs) for this trail.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the name of the trail. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	// - Start with a letter or number, and end with a letter or number
	// - Be between 3 and 128 characters
	// - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
	// - Not be in IP address format (for example, 192.168.5.4)
	TrailName *string `field:"optional" json:"trailName" yaml:"trailName"`
}


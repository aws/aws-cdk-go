package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for creating a data protection policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var dataIdentifier dataIdentifier
//   var logGroup logGroup
//
//   dataProtectionPolicyProps := &DataProtectionPolicyProps{
//   	Identifiers: []*dataIdentifier{
//   		dataIdentifier,
//   	},
//
//   	// the properties below are optional
//   	DeliveryStreamNameAuditDestination: jsii.String("deliveryStreamNameAuditDestination"),
//   	Description: jsii.String("description"),
//   	LogGroupAuditDestination: logGroup,
//   	Name: jsii.String("name"),
//   	S3BucketAuditDestination: bucket,
//   }
//
type DataProtectionPolicyProps struct {
	// List of data protection identifiers.
	//
	// Must be in the following list: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/protect-sensitive-log-data-types.html
	Identifiers *[]DataIdentifier `field:"required" json:"identifiers" yaml:"identifiers"`
	// Amazon Kinesis Data Firehose delivery stream to send audit findings to.
	//
	// The delivery stream must already exist.
	DeliveryStreamNameAuditDestination *string `field:"optional" json:"deliveryStreamNameAuditDestination" yaml:"deliveryStreamNameAuditDestination"`
	// Description of the data protection policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// CloudWatch Logs log group to send audit findings to.
	//
	// The log group must already exist prior to creating the data protection policy.
	LogGroupAuditDestination ILogGroup `field:"optional" json:"logGroupAuditDestination" yaml:"logGroupAuditDestination"`
	// Name of the data protection policy.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// S3 bucket to send audit findings to.
	//
	// The bucket must already exist.
	S3BucketAuditDestination awss3.IBucket `field:"optional" json:"s3BucketAuditDestination" yaml:"s3BucketAuditDestination"`
}


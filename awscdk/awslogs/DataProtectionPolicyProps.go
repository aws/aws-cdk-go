package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for creating a data protection policy.
//
// Example:
//   import kinesisfirehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//
//   logGroupDestination := logs.NewLogGroup(this, jsii.String("LogGroupLambdaAudit"), &LogGroupProps{
//   	LogGroupName: jsii.String("auditDestinationForCDK"),
//   })
//
//   bucket := s3.NewBucket(this, jsii.String("audit-bucket"))
//   s3Destination := destinations.NewS3Bucket(bucket)
//
//   deliveryStream := kinesisfirehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
//   dataProtectionPolicy := logs.NewDataProtectionPolicy(&DataProtectionPolicyProps{
//   	Name: jsii.String("data protection policy"),
//   	Description: jsii.String("policy description"),
//   	Identifiers: []dataIdentifier{
//   		logs.*dataIdentifier_DRIVERSLICENSE_US(),
//   		 // managed data identifier
//   		logs.NewDataIdentifier(jsii.String("EmailAddress")),
//   		 // forward compatibility for new managed data identifiers
//   		logs.NewCustomDataIdentifier(jsii.String("EmployeeId"), jsii.String("EmployeeId-\\d{9}")),
//   	},
//   	 // custom data identifier
//   	LogGroupAuditDestination: logGroupDestination,
//   	S3BucketAuditDestination: bucket,
//   	DeliveryStreamNameAuditDestination: deliveryStream.DeliveryStreamName,
//   })
//
//   logs.NewLogGroup(this, jsii.String("LogGroupLambda"), &LogGroupProps{
//   	LogGroupName: jsii.String("cdkIntegLogGroup"),
//   	DataProtectionPolicy: dataProtectionPolicy,
//   })
//
type DataProtectionPolicyProps struct {
	// List of data protection identifiers.
	//
	// Managed data identifiers must be in the following list: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL-managed-data-identifiers.html
	// Custom data identifiers must have a valid regex defined: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL-custom-data-identifiers.html#custom-data-identifiers-constraints
	Identifiers *[]DataIdentifier `field:"required" json:"identifiers" yaml:"identifiers"`
	// Amazon Kinesis Data Firehose delivery stream to send audit findings to.
	//
	// The delivery stream must already exist.
	// Default: - no firehose delivery stream audit destination.
	//
	DeliveryStreamNameAuditDestination *string `field:"optional" json:"deliveryStreamNameAuditDestination" yaml:"deliveryStreamNameAuditDestination"`
	// Description of the data protection policy.
	// Default: - 'cdk generated data protection policy'.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// CloudWatch Logs log group to send audit findings to.
	//
	// The log group must already exist prior to creating the data protection policy.
	// Default: - no CloudWatch Logs audit destination.
	//
	LogGroupAuditDestination ILogGroup `field:"optional" json:"logGroupAuditDestination" yaml:"logGroupAuditDestination"`
	// Name of the data protection policy.
	// Default: - 'data-protection-policy-cdk'.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// S3 bucket to send audit findings to.
	//
	// The bucket must already exist.
	// Default: - no S3 bucket audit destination.
	//
	S3BucketAuditDestination awss3.IBucket `field:"optional" json:"s3BucketAuditDestination" yaml:"s3BucketAuditDestination"`
}


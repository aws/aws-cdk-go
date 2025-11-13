package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Properties for a LogGroup.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroupDestination := logs.NewLogGroup(this, jsii.String("LogGroupLambdaAudit"), &LogGroupProps{
//   	LogGroupName: jsii.String("auditDestinationForCDK"),
//   })
//
//   bucket := s3.NewBucket(this, jsii.String("audit-bucket"))
//   s3Destination := firehose.NewS3Bucket(bucket)
//
//   deliveryStream := firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
//   dataProtectionPolicy := logs.NewDataProtectionPolicy(&DataProtectionPolicyProps{
//   	Name: jsii.String("data protection policy"),
//   	Description: jsii.String("policy description"),
//   	Identifiers: []DataIdentifier{
//   		logs.DataIdentifier_DRIVERSLICENSE_US(),
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
type LogGroupProps struct {
	// Data Protection Policy for this log group.
	// Default: - no data protection policy.
	//
	DataProtectionPolicy DataProtectionPolicy `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// The KMS customer managed key to encrypt the log group with.
	// Default: Server-side encryption managed by the CloudWatch Logs service.
	//
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Field Index Policies for this log group.
	// Default: - no field index policies for this log group.
	//
	FieldIndexPolicies *[]FieldIndexPolicy `field:"optional" json:"fieldIndexPolicies" yaml:"fieldIndexPolicies"`
	// The class of the log group. Possible values are: STANDARD and INFREQUENT_ACCESS.
	//
	// INFREQUENT_ACCESS class provides customers a cost-effective way to consolidate
	// logs which supports querying using Logs Insights. The logGroupClass property cannot
	// be changed once the log group is created.
	// Default: LogGroupClass.STANDARD
	//
	LogGroupClass LogGroupClass `field:"optional" json:"logGroupClass" yaml:"logGroupClass"`
	// Name of the log group.
	// Default: Automatically generated.
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Determine the removal policy of this log group.
	//
	// Normally you want to retain the log group so you can diagnose issues
	// from logs even after a deployment that no longer includes the log group.
	// In that case, use the normal date-based retention policy to age out your
	// logs.
	// Default: RemovalPolicy.Retain
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// How long, in days, the log contents will be retained.
	//
	// To retain all logs, set this value to RetentionDays.INFINITE.
	// Default: RetentionDays.TWO_YEARS
	//
	Retention RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}


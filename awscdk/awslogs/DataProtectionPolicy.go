package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Creates a data protection policy for CloudWatch Logs log groups.
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
//   	DeliveryStreamNameAuditDestination: deliveryStream.deliveryStreamName,
//   })
//
//   logs.NewLogGroup(this, jsii.String("LogGroupLambda"), &LogGroupProps{
//   	LogGroupName: jsii.String("cdkIntegLogGroup"),
//   	DataProtectionPolicy: dataProtectionPolicy,
//   })
//
type DataProtectionPolicy interface {
}

// The jsii proxy struct for DataProtectionPolicy
type jsiiProxy_DataProtectionPolicy struct {
	_ byte // padding
}

func NewDataProtectionPolicy(props *DataProtectionPolicyProps) DataProtectionPolicy {
	_init_.Initialize()

	if err := validateNewDataProtectionPolicyParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataProtectionPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.DataProtectionPolicy",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewDataProtectionPolicy_Override(d DataProtectionPolicy, props *DataProtectionPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.DataProtectionPolicy",
		[]interface{}{props},
		d,
	)
}


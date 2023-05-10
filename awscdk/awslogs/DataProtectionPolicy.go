package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Creates a data protection policy for CloudWatch Logs log groups.
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
//   dataProtectionPolicy := awscdk.Aws_logs.NewDataProtectionPolicy(&DataProtectionPolicyProps{
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


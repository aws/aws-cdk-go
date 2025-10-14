package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an SVCB record value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   svcbRecordValue := awscdk.Aws_route53.SvcbRecordValue_Alias(jsii.String("targetName"))
//
type SvcbRecordValue interface {
	// Returns the string representation of SVCB and HTTPS record value.
	ToString() *string
}

// The jsii proxy struct for SvcbRecordValue
type jsiiProxy_SvcbRecordValue struct {
	_ byte // padding
}

// An SVCB AliasMode record value.
func SvcbRecordValue_Alias(targetName *string) SvcbRecordValue {
	_init_.Initialize()

	if err := validateSvcbRecordValue_AliasParameters(targetName); err != nil {
		panic(err)
	}
	var returns SvcbRecordValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.SvcbRecordValue",
		"alias",
		[]interface{}{targetName},
		&returns,
	)

	return returns
}

// An SVCB ServiceMode record value.
func SvcbRecordValue_Service(props *SvcbRecordServiceModeProps) SvcbRecordValue {
	_init_.Initialize()

	if err := validateSvcbRecordValue_ServiceParameters(props); err != nil {
		panic(err)
	}
	var returns SvcbRecordValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.SvcbRecordValue",
		"service",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SvcbRecordValue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


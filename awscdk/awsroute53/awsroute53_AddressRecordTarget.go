package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Target for a DNS A Record.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var aliasRecordTarget iAliasRecordTarget
//
//   addressRecordTarget := awscdk.Aws_route53.addressRecordTarget.fromAlias(aliasRecordTarget)
//
// Deprecated: Use RecordTarget.
type AddressRecordTarget interface {
	RecordTarget
	// alias for targets such as CloudFront distribution to route traffic to.
	// Deprecated: Use RecordTarget.
	AliasTarget() IAliasRecordTarget
	// correspond with the chosen record type (e.g. for 'A' Type, specify one or more IP addresses).
	// Deprecated: Use RecordTarget.
	Values() *[]*string
}

// The jsii proxy struct for AddressRecordTarget
type jsiiProxy_AddressRecordTarget struct {
	jsiiProxy_RecordTarget
}

func (j *jsiiProxy_AddressRecordTarget) AliasTarget() IAliasRecordTarget {
	var returns IAliasRecordTarget
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AddressRecordTarget) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Deprecated: Use RecordTarget.
func NewAddressRecordTarget(values *[]*string, aliasTarget IAliasRecordTarget) AddressRecordTarget {
	_init_.Initialize()

	j := jsiiProxy_AddressRecordTarget{}

	_jsii_.Create(
		"monocdk.aws_route53.AddressRecordTarget",
		[]interface{}{values, aliasTarget},
		&j,
	)

	return &j
}

// Deprecated: Use RecordTarget.
func NewAddressRecordTarget_Override(a AddressRecordTarget, values *[]*string, aliasTarget IAliasRecordTarget) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53.AddressRecordTarget",
		[]interface{}{values, aliasTarget},
		a,
	)
}

// Use an alias as target.
// Deprecated: Use RecordTarget.
func AddressRecordTarget_FromAlias(aliasTarget IAliasRecordTarget) RecordTarget {
	_init_.Initialize()

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"monocdk.aws_route53.AddressRecordTarget",
		"fromAlias",
		[]interface{}{aliasTarget},
		&returns,
	)

	return returns
}

// Use ip addresses as target.
// Deprecated: Use RecordTarget.
func AddressRecordTarget_FromIpAddresses(ipAddresses ...*string) RecordTarget {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range ipAddresses {
		args = append(args, a)
	}

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"monocdk.aws_route53.AddressRecordTarget",
		"fromIpAddresses",
		args,
		&returns,
	)

	return returns
}

// Use string values as target.
// Deprecated: Use RecordTarget.
func AddressRecordTarget_FromValues(values ...*string) RecordTarget {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"monocdk.aws_route53.AddressRecordTarget",
		"fromValues",
		args,
		&returns,
	)

	return returns
}


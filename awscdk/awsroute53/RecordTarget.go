package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Type union for a record that accepts multiple types of target.
//
// Example:
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//
//   var myZone hostedZone
//   var distribution cloudFrontWebDistribution
//
//   route53.NewAaaaRecord(this, jsii.String("Alias"), &AaaaRecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
//   })
//
type RecordTarget interface {
	// alias for targets such as CloudFront distribution to route traffic to.
	AliasTarget() IAliasRecordTarget
	// correspond with the chosen record type (e.g. for 'A' Type, specify one or more IP addresses).
	Values() *[]*string
}

// The jsii proxy struct for RecordTarget
type jsiiProxy_RecordTarget struct {
	_ byte // padding
}

func (j *jsiiProxy_RecordTarget) AliasTarget() IAliasRecordTarget {
	var returns IAliasRecordTarget
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecordTarget) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


func NewRecordTarget(values *[]*string, aliasTarget IAliasRecordTarget) RecordTarget {
	_init_.Initialize()

	j := jsiiProxy_RecordTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.RecordTarget",
		[]interface{}{values, aliasTarget},
		&j,
	)

	return &j
}

func NewRecordTarget_Override(r RecordTarget, values *[]*string, aliasTarget IAliasRecordTarget) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.RecordTarget",
		[]interface{}{values, aliasTarget},
		r,
	)
}

// Use an alias as target.
func RecordTarget_FromAlias(aliasTarget IAliasRecordTarget) RecordTarget {
	_init_.Initialize()

	if err := validateRecordTarget_FromAliasParameters(aliasTarget); err != nil {
		panic(err)
	}
	var returns RecordTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordTarget",
		"fromAlias",
		[]interface{}{aliasTarget},
		&returns,
	)

	return returns
}

// Use ip addresses as target.
func RecordTarget_FromIpAddresses(ipAddresses ...*string) RecordTarget {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range ipAddresses {
		args = append(args, a)
	}

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordTarget",
		"fromIpAddresses",
		args,
		&returns,
	)

	return returns
}

// Use string values as target.
func RecordTarget_FromValues(values ...*string) RecordTarget {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns RecordTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.RecordTarget",
		"fromValues",
		args,
		&returns,
	)

	return returns
}


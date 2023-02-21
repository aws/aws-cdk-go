package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Classes that are valid alias record targets, like CloudFront distributions and load balancers, should implement this interface.
type IAliasRecordTarget interface {
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(record IRecordSet, zone IHostedZone) *AliasRecordTargetConfig
}

// The jsii proxy for IAliasRecordTarget
type jsiiProxy_IAliasRecordTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IAliasRecordTarget) Bind(record IRecordSet, zone IHostedZone) *AliasRecordTargetConfig {
	if err := i.validateBindParameters(record); err != nil {
		panic(err)
	}
	var returns *AliasRecordTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{record, zone},
		&returns,
	)

	return returns
}


package awsroute53targets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties the alias record target.
type IAliasRecordTargetProps interface {
	// Evaluate target health.
	// Default: - no health check configuration.
	//
	EvaluateTargetHealth() *bool
}

// The jsii proxy for IAliasRecordTargetProps
type jsiiProxy_IAliasRecordTargetProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IAliasRecordTargetProps) EvaluateTargetHealth() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"evaluateTargetHealth",
		&returns,
	)
	return returns
}


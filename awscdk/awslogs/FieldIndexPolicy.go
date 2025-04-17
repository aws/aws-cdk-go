package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Creates a field index policy for CloudWatch Logs log groups.
//
// Example:
//   fieldIndexPolicy := logs.NewFieldIndexPolicy(&FieldIndexPolicyProps{
//   	Fields: []*string{
//   		jsii.String("Operation"),
//   		jsii.String("RequestId"),
//   	},
//   })
//
//   logs.NewLogGroup(this, jsii.String("LogGroup"), &LogGroupProps{
//   	LogGroupName: jsii.String("cdkIntegLogGroup"),
//   	FieldIndexPolicies: []fieldIndexPolicy{
//   		fieldIndexPolicy,
//   	},
//   })
//
type FieldIndexPolicy interface {
}

// The jsii proxy struct for FieldIndexPolicy
type jsiiProxy_FieldIndexPolicy struct {
	_ byte // padding
}

func NewFieldIndexPolicy(props *FieldIndexPolicyProps) FieldIndexPolicy {
	_init_.Initialize()

	if err := validateNewFieldIndexPolicyParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FieldIndexPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.FieldIndexPolicy",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewFieldIndexPolicy_Override(f FieldIndexPolicy, props *FieldIndexPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.FieldIndexPolicy",
		[]interface{}{props},
		f,
	)
}


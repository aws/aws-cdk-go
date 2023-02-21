package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class to allow assigning a value or an auto-generated id to a partition key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionKeyStep := awscdk.Aws_appsync.NewPartitionKeyStep(jsii.String("key"))
//
type PartitionKeyStep interface {
	// Assign an auto-generated value to the partition key.
	Auto() PartitionKey
	// Assign an auto-generated value to the partition key.
	Is(val *string) PartitionKey
}

// The jsii proxy struct for PartitionKeyStep
type jsiiProxy_PartitionKeyStep struct {
	_ byte // padding
}

func NewPartitionKeyStep(key *string) PartitionKeyStep {
	_init_.Initialize()

	if err := validateNewPartitionKeyStepParameters(key); err != nil {
		panic(err)
	}
	j := jsiiProxy_PartitionKeyStep{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.PartitionKeyStep",
		[]interface{}{key},
		&j,
	)

	return &j
}

func NewPartitionKeyStep_Override(p PartitionKeyStep, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.PartitionKeyStep",
		[]interface{}{key},
		p,
	)
}

func (p *jsiiProxy_PartitionKeyStep) Auto() PartitionKey {
	var returns PartitionKey

	_jsii_.Invoke(
		p,
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartitionKeyStep) Is(val *string) PartitionKey {
	if err := p.validateIsParameters(val); err != nil {
		panic(err)
	}
	var returns PartitionKey

	_jsii_.Invoke(
		p,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}


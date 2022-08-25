package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
// Experimental.
type PartitionKeyStep interface {
	// Assign an auto-generated value to the partition key.
	// Experimental.
	Auto() PartitionKey
	// Assign an auto-generated value to the partition key.
	// Experimental.
	Is(val *string) PartitionKey
}

// The jsii proxy struct for PartitionKeyStep
type jsiiProxy_PartitionKeyStep struct {
	_ byte // padding
}

// Experimental.
func NewPartitionKeyStep(key *string) PartitionKeyStep {
	_init_.Initialize()

	j := jsiiProxy_PartitionKeyStep{}

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKeyStep",
		[]interface{}{key},
		&j,
	)

	return &j
}

// Experimental.
func NewPartitionKeyStep_Override(p PartitionKeyStep, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKeyStep",
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
	var returns PartitionKey

	_jsii_.Invoke(
		p,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}


package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the assignment to the partition key.
//
// It can be
// enhanced with the assignment of the sort key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign assign
//
//   partitionKey := awscdk.Aws_appsync.NewPartitionKey(assign)
//
// Experimental.
type PartitionKey interface {
	PrimaryKey
	// Experimental.
	Pkey() Assign
	// Renders the key assignment to a VTL string.
	// Experimental.
	RenderTemplate() *string
	// Allows assigning a value to the sort key.
	// Experimental.
	Sort(key *string) SortKeyStep
}

// The jsii proxy struct for PartitionKey
type jsiiProxy_PartitionKey struct {
	jsiiProxy_PrimaryKey
}

func (j *jsiiProxy_PartitionKey) Pkey() Assign {
	var returns Assign
	_jsii_.Get(
		j,
		"pkey",
		&returns,
	)
	return returns
}


// Experimental.
func NewPartitionKey(pkey Assign) PartitionKey {
	_init_.Initialize()

	if err := validateNewPartitionKeyParameters(pkey); err != nil {
		panic(err)
	}
	j := jsiiProxy_PartitionKey{}

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKey",
		[]interface{}{pkey},
		&j,
	)

	return &j
}

// Experimental.
func NewPartitionKey_Override(p PartitionKey, pkey Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKey",
		[]interface{}{pkey},
		p,
	)
}

// Allows assigning a value to the partition key.
// Experimental.
func PartitionKey_Partition(key *string) PartitionKeyStep {
	_init_.Initialize()

	if err := validatePartitionKey_PartitionParameters(key); err != nil {
		panic(err)
	}
	var returns PartitionKeyStep

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.PartitionKey",
		"partition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartitionKey) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PartitionKey) Sort(key *string) SortKeyStep {
	if err := p.validateSortParameters(key); err != nil {
		panic(err)
	}
	var returns SortKeyStep

	_jsii_.Invoke(
		p,
		"sort",
		[]interface{}{key},
		&returns,
	)

	return returns
}


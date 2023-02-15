package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class to allow assigning a value or an auto-generated id to a sort key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign assign
//
//   sortKeyStep := awscdk.Aws_appsync.NewSortKeyStep(assign, jsii.String("skey"))
//
type SortKeyStep interface {
	// Assign an auto-generated value to the sort key.
	Auto() PrimaryKey
	// Assign an auto-generated value to the sort key.
	Is(val *string) PrimaryKey
}

// The jsii proxy struct for SortKeyStep
type jsiiProxy_SortKeyStep struct {
	_ byte // padding
}

func NewSortKeyStep(pkey Assign, skey *string) SortKeyStep {
	_init_.Initialize()

	if err := validateNewSortKeyStepParameters(pkey, skey); err != nil {
		panic(err)
	}
	j := jsiiProxy_SortKeyStep{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.SortKeyStep",
		[]interface{}{pkey, skey},
		&j,
	)

	return &j
}

func NewSortKeyStep_Override(s SortKeyStep, pkey Assign, skey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.SortKeyStep",
		[]interface{}{pkey, skey},
		s,
	)
}

func (s *jsiiProxy_SortKeyStep) Auto() PrimaryKey {
	var returns PrimaryKey

	_jsii_.Invoke(
		s,
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SortKeyStep) Is(val *string) PrimaryKey {
	if err := s.validateIsParameters(val); err != nil {
		panic(err)
	}
	var returns PrimaryKey

	_jsii_.Invoke(
		s,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}


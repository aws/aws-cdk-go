package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Converts all fragments to strings and concats those.
//
// Drops 'undefined's.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   stringConcat := cdk.NewStringConcat()
//
type StringConcat interface {
	IFragmentConcatenator
	// Join the fragment on the left and on the right.
	Join(left interface{}, right interface{}) interface{}
}

// The jsii proxy struct for StringConcat
type jsiiProxy_StringConcat struct {
	jsiiProxy_IFragmentConcatenator
}

func NewStringConcat() StringConcat {
	_init_.Initialize()

	j := jsiiProxy_StringConcat{}

	_jsii_.Create(
		"aws-cdk-lib.StringConcat",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStringConcat_Override(s StringConcat) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.StringConcat",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StringConcat) Join(left interface{}, right interface{}) interface{} {
	if err := s.validateJoinParameters(left, right); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"join",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}


package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Define dynamic target parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//
//   targetParameter := pipes_alpha.NewTargetParameter()
//
// Experimental.
type TargetParameter interface {
}

// The jsii proxy struct for TargetParameter
type jsiiProxy_TargetParameter struct {
	_ byte // padding
}

// Experimental.
func NewTargetParameter() TargetParameter {
	_init_.Initialize()

	j := jsiiProxy_TargetParameter{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.TargetParameter",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTargetParameter_Override(t TargetParameter) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.TargetParameter",
		nil, // no parameters
		t,
	)
}

// Target parameter based on a jsonPath expression from the incoming event.
// Experimental.
func TargetParameter_FromJsonPath(jsonPath *string) *string {
	_init_.Initialize()

	if err := validateTargetParameter_FromJsonPathParameters(jsonPath); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-alpha.TargetParameter",
		"fromJsonPath",
		[]interface{}{jsonPath},
		&returns,
	)

	return returns
}


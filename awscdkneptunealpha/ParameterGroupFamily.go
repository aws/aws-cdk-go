// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The DB parameter group family that a DB parameter group is compatible with.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import neptune_alpha "github.com/aws/aws-cdk-go/awscdkneptunealpha"
//
//   parameterGroupFamily := neptune_alpha.NewParameterGroupFamily(jsii.String("family"))
//
// Experimental.
type ParameterGroupFamily interface {
	// the family of the parameter group Neptune.
	// Experimental.
	Family() *string
}

// The jsii proxy struct for ParameterGroupFamily
type jsiiProxy_ParameterGroupFamily struct {
	_ byte // padding
}

func (j *jsiiProxy_ParameterGroupFamily) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}


// Constructor for specifying a custom parameter group famil.
// Experimental.
func NewParameterGroupFamily(family *string) ParameterGroupFamily {
	_init_.Initialize()

	if err := validateNewParameterGroupFamilyParameters(family); err != nil {
		panic(err)
	}
	j := jsiiProxy_ParameterGroupFamily{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.ParameterGroupFamily",
		[]interface{}{family},
		&j,
	)

	return &j
}

// Constructor for specifying a custom parameter group famil.
// Experimental.
func NewParameterGroupFamily_Override(p ParameterGroupFamily, family *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.ParameterGroupFamily",
		[]interface{}{family},
		p,
	)
}

func ParameterGroupFamily_NEPTUNE_1() ParameterGroupFamily {
	_init_.Initialize()
	var returns ParameterGroupFamily
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.ParameterGroupFamily",
		"NEPTUNE_1",
		&returns,
	)
	return returns
}

func ParameterGroupFamily_NEPTUNE_1_2() ParameterGroupFamily {
	_init_.Initialize()
	var returns ParameterGroupFamily
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.ParameterGroupFamily",
		"NEPTUNE_1_2",
		&returns,
	)
	return returns
}


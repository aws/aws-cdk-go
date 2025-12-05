package core

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Selects constructs from a construct tree based on various criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   constructSelector := awscdkmixinspreview.Core.NewConstructSelector()
//
// Experimental.
type ConstructSelector interface {
}

// The jsii proxy struct for ConstructSelector
type jsiiProxy_ConstructSelector struct {
	_ byte // padding
}

// Experimental.
func NewConstructSelector() ConstructSelector {
	_init_.Initialize()

	j := jsiiProxy_ConstructSelector{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConstructSelector_Override(c ConstructSelector) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		nil, // no parameters
		c,
	)
}

// Selects all constructs in the tree.
// Experimental.
func ConstructSelector_All() IConstructSelector {
	_init_.Initialize()

	var returns IConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Selects constructs whose construct IDs match a pattern.
//
// Uses glob like matching.
// Experimental.
func ConstructSelector_ById(pattern *string) IConstructSelector {
	_init_.Initialize()

	if err := validateConstructSelector_ByIdParameters(pattern); err != nil {
		panic(err)
	}
	var returns IConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"byId",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Selects constructs whose construct paths match a pattern.
//
// Uses glob like matching.
// Experimental.
func ConstructSelector_ByPath(pattern *string) IConstructSelector {
	_init_.Initialize()

	if err := validateConstructSelector_ByPathParameters(pattern); err != nil {
		panic(err)
	}
	var returns IConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"byPath",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Selects CfnResource constructs or the default CfnResource child.
// Experimental.
func ConstructSelector_CfnResource() IConstructSelector {
	_init_.Initialize()

	var returns IConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"cfnResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Selects only the provided construct.
// Experimental.
func ConstructSelector_OnlyItself() IConstructSelector {
	_init_.Initialize()

	var returns IConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"onlyItself",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Selects constructs of a specific type.
// Experimental.
func ConstructSelector_ResourcesOfType(types ...*string) IConstructSelector {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range types {
		args = append(args, a)
	}

	var returns IConstructSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.core.ConstructSelector",
		"resourcesOfType",
		args,
		&returns,
	)

	return returns
}


package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Functions for devising unique names for constructs.
//
// For example, those can be
// used to allocate unique physical names for resources.
type Names interface {
}

// The jsii proxy struct for Names
type jsiiProxy_Names struct {
	_ byte // padding
}

// Returns a CloudFormation-compatible unique identifier for a construct based on its path.
//
// The identifier includes a human readable portion rendered
// from the path components and a hash suffix.
//
// TODO (v2): replace with API to use `constructs.Node`.
//
// Returns: a unique id based on the construct path.
func Names_NodeUniqueId(node constructs.Node) *string {
	_init_.Initialize()

	if err := validateNames_NodeUniqueIdParameters(node); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Names",
		"nodeUniqueId",
		[]interface{}{node},
		&returns,
	)

	return returns
}

// Returns a CloudFormation-compatible unique identifier for a construct based on its path.
//
// The identifier includes a human readable portion rendered
// from the path components and a hash suffix. uniqueId is not unique if multiple
// copies of the stack are deployed. Prefer using uniqueResourceName().
//
// Returns: a unique id based on the construct path.
func Names_UniqueId(construct constructs.IConstruct) *string {
	_init_.Initialize()

	if err := validateNames_UniqueIdParameters(construct); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Names",
		"uniqueId",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Returns a CloudFormation-compatible unique identifier for a construct based on its path.
//
// This function finds the stackName of the parent stack (non-nested)
// to the construct, and the ids of the components in the construct path.
//
// The user can define allowed special characters, a separator between the elements,
// and the maximum length of the resource name. The name includes a human readable portion rendered
// from the path components, with or without user defined separators, and a hash suffix.
// If the resource name is longer than the maximum length, it is trimmed in the middle.
//
// Returns: a unique resource name based on the construct path.
func Names_UniqueResourceName(construct constructs.IConstruct, options *UniqueResourceNameOptions) *string {
	_init_.Initialize()

	if err := validateNames_UniqueResourceNameParameters(construct, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Names",
		"uniqueResourceName",
		[]interface{}{construct, options},
		&returns,
	)

	return returns
}


// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// Functions for devising unique names for constructs.
//
// For example, those can be
// used to allocate unique physical names for resources.
// Experimental.
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
// Experimental.
func Names_NodeUniqueId(node ConstructNode) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.Names",
		"nodeUniqueId",
		[]interface{}{node},
		&returns,
	)

	return returns
}

// Returns a CloudFormation-compatible unique identifier for a construct based on its path.
//
// The identifier includes a human readable portion rendered
// from the path components and a hash suffix.
//
// Returns: a unique id based on the construct path.
// Experimental.
func Names_UniqueId(construct constructs.Construct) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.Names",
		"uniqueId",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


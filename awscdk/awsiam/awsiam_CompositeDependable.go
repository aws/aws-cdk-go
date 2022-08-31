package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam/internal"
)

// Composite dependable.
//
// Not as simple as eagerly getting the dependency roots from the
// inner dependables, as they may be mutable so we need to defer
// the query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dependable iDependable
//
//   compositeDependable := awscdk.Aws_iam.NewCompositeDependable(dependable)
//
// Experimental.
type CompositeDependable interface {
	awscdk.IDependable
}

// The jsii proxy struct for CompositeDependable
type jsiiProxy_CompositeDependable struct {
	internal.Type__awscdkIDependable
}

// Experimental.
func NewCompositeDependable(dependables ...awscdk.IDependable) CompositeDependable {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range dependables {
		args = append(args, a)
	}

	j := jsiiProxy_CompositeDependable{}

	_jsii_.Create(
		"monocdk.aws_iam.CompositeDependable",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewCompositeDependable_Override(c CompositeDependable, dependables ...awscdk.IDependable) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range dependables {
		args = append(args, a)
	}

	_jsii_.Create(
		"monocdk.aws_iam.CompositeDependable",
		args,
		c,
	)
}


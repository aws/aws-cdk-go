package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var dependable iDependable
//
//   compositeDependable := awscdk.Aws_iam.NewCompositeDependable(dependable)
//
type CompositeDependable interface {
	constructs.IDependable
}

// The jsii proxy struct for CompositeDependable
type jsiiProxy_CompositeDependable struct {
	internal.Type__constructsIDependable
}

func NewCompositeDependable(dependables ...constructs.IDependable) CompositeDependable {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range dependables {
		args = append(args, a)
	}

	j := jsiiProxy_CompositeDependable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.CompositeDependable",
		args,
		&j,
	)

	return &j
}

func NewCompositeDependable_Override(c CompositeDependable, dependables ...constructs.IDependable) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range dependables {
		args = append(args, a)
	}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.CompositeDependable",
		args,
		c,
	)
}


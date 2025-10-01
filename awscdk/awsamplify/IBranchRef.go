package awsamplify

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplify/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Branch.
// Experimental.
type IBranchRef interface {
	constructs.IConstruct
	// A reference to a Branch resource.
	// Experimental.
	BranchRef() *BranchReference
}

// The jsii proxy for IBranchRef
type jsiiProxy_IBranchRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBranchRef) BranchRef() *BranchReference {
	var returns *BranchReference
	_jsii_.Get(
		j,
		"branchRef",
		&returns,
	)
	return returns
}


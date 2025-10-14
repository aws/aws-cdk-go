package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FindingsFilter.
// Experimental.
type IFindingsFilterRef interface {
	constructs.IConstruct
	// A reference to a FindingsFilter resource.
	// Experimental.
	FindingsFilterRef() *FindingsFilterReference
}

// The jsii proxy for IFindingsFilterRef
type jsiiProxy_IFindingsFilterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFindingsFilterRef) FindingsFilterRef() *FindingsFilterReference {
	var returns *FindingsFilterReference
	_jsii_.Get(
		j,
		"findingsFilterRef",
		&returns,
	)
	return returns
}


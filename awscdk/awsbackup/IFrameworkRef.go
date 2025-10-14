package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Framework.
// Experimental.
type IFrameworkRef interface {
	constructs.IConstruct
	// A reference to a Framework resource.
	// Experimental.
	FrameworkRef() *FrameworkReference
}

// The jsii proxy for IFrameworkRef
type jsiiProxy_IFrameworkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFrameworkRef) FrameworkRef() *FrameworkReference {
	var returns *FrameworkReference
	_jsii_.Get(
		j,
		"frameworkRef",
		&returns,
	)
	return returns
}


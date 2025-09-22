package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxWindows.
// Experimental.
type ILocationFSxWindowsRef interface {
	constructs.IConstruct
	// A reference to a LocationFSxWindows resource.
	// Experimental.
	LocationFSxWindowsRef() *LocationFSxWindowsReference
}

// The jsii proxy for ILocationFSxWindowsRef
type jsiiProxy_ILocationFSxWindowsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocationFSxWindowsRef) LocationFSxWindowsRef() *LocationFSxWindowsReference {
	var returns *LocationFSxWindowsReference
	_jsii_.Get(
		j,
		"locationFSxWindowsRef",
		&returns,
	)
	return returns
}


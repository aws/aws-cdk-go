package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Portal.
// Experimental.
type IPortalRef interface {
	constructs.IConstruct
	// A reference to a Portal resource.
	// Experimental.
	PortalRef() *PortalReference
}

// The jsii proxy for IPortalRef
type jsiiProxy_IPortalRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPortalRef) PortalRef() *PortalReference {
	var returns *PortalReference
	_jsii_.Get(
		j,
		"portalRef",
		&returns,
	)
	return returns
}


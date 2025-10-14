package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Hub.
// Experimental.
type IHubRef interface {
	constructs.IConstruct
	// A reference to a Hub resource.
	// Experimental.
	HubRef() *HubReference
}

// The jsii proxy for IHubRef
type jsiiProxy_IHubRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHubRef) HubRef() *HubReference {
	var returns *HubReference
	_jsii_.Get(
		j,
		"hubRef",
		&returns,
	)
	return returns
}


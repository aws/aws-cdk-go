package awseventschemas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Discoverer.
// Experimental.
type IDiscovererRef interface {
	constructs.IConstruct
	// A reference to a Discoverer resource.
	// Experimental.
	DiscovererRef() *DiscovererReference
}

// The jsii proxy for IDiscovererRef
type jsiiProxy_IDiscovererRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDiscovererRef) DiscovererRef() *DiscovererReference {
	var returns *DiscovererReference
	_jsii_.Get(
		j,
		"discovererRef",
		&returns,
	)
	return returns
}


package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationSMB.
// Experimental.
type ILocationSMBRef interface {
	constructs.IConstruct
	// A reference to a LocationSMB resource.
	// Experimental.
	LocationSmbRef() *LocationSMBReference
}

// The jsii proxy for ILocationSMBRef
type jsiiProxy_ILocationSMBRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocationSMBRef) LocationSmbRef() *LocationSMBReference {
	var returns *LocationSMBReference
	_jsii_.Get(
		j,
		"locationSmbRef",
		&returns,
	)
	return returns
}


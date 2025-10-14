package awsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPoint.
// Experimental.
type IAccessPointRef interface {
	constructs.IConstruct
	// A reference to a AccessPoint resource.
	// Experimental.
	AccessPointRef() *AccessPointReference
}

// The jsii proxy for IAccessPointRef
type jsiiProxy_IAccessPointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessPointRef) AccessPointRef() *AccessPointReference {
	var returns *AccessPointReference
	_jsii_.Get(
		j,
		"accessPointRef",
		&returns,
	)
	return returns
}


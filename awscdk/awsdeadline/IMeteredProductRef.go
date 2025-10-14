package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MeteredProduct.
// Experimental.
type IMeteredProductRef interface {
	constructs.IConstruct
	// A reference to a MeteredProduct resource.
	// Experimental.
	MeteredProductRef() *MeteredProductReference
}

// The jsii proxy for IMeteredProductRef
type jsiiProxy_IMeteredProductRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMeteredProductRef) MeteredProductRef() *MeteredProductReference {
	var returns *MeteredProductReference
	_jsii_.Get(
		j,
		"meteredProductRef",
		&returns,
	)
	return returns
}


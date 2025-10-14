package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerCertificate.
// Experimental.
type IServerCertificateRef interface {
	constructs.IConstruct
	// A reference to a ServerCertificate resource.
	// Experimental.
	ServerCertificateRef() *ServerCertificateReference
}

// The jsii proxy for IServerCertificateRef
type jsiiProxy_IServerCertificateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServerCertificateRef) ServerCertificateRef() *ServerCertificateReference {
	var returns *ServerCertificateReference
	_jsii_.Get(
		j,
		"serverCertificateRef",
		&returns,
	)
	return returns
}


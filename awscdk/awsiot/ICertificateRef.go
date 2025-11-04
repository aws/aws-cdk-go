package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Certificate.
// Experimental.
type ICertificateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Certificate resource.
	// Experimental.
	CertificateRef() *CertificateReference
}

// The jsii proxy for ICertificateRef
type jsiiProxy_ICertificateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICertificateRef) CertificateRef() *CertificateReference {
	var returns *CertificateReference
	_jsii_.Get(
		j,
		"certificateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


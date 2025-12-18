package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsacmpca"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface which all CertificateAuthority based class must implement.
type ICertificateAuthority interface {
	interfacesawsacmpca.ICertificateAuthorityRef
	awscdk.IResource
	// The Amazon Resource Name of the Certificate.
	CertificateAuthorityArn() *string
}

// The jsii proxy for ICertificateAuthority
type jsiiProxy_ICertificateAuthority struct {
	internal.Type__interfacesawsacmpcaICertificateAuthorityRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICertificateAuthority) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ICertificateAuthority) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthority) CertificateAuthorityRef() *interfacesawsacmpca.CertificateAuthorityReference {
	var returns *interfacesawsacmpca.CertificateAuthorityReference
	_jsii_.Get(
		j,
		"certificateAuthorityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthority) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthority) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthority) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


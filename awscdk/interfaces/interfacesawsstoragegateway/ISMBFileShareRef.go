package interfacesawsstoragegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstoragegateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SMBFileShare.
// Experimental.
type ISMBFileShareRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SMBFileShare resource.
	// Experimental.
	SmbFileShareRef() *SMBFileShareReference
}

// The jsii proxy for ISMBFileShareRef
type jsiiProxy_ISMBFileShareRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISMBFileShareRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISMBFileShareRef) SmbFileShareRef() *SMBFileShareReference {
	var returns *SMBFileShareReference
	_jsii_.Get(
		j,
		"smbFileShareRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISMBFileShareRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISMBFileShareRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawssecuritylake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AwsLogSource.
// Experimental.
type IAwsLogSourceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AwsLogSource resource.
	// Experimental.
	AwsLogSourceRef() *AwsLogSourceReference
}

// The jsii proxy for IAwsLogSourceRef
type jsiiProxy_IAwsLogSourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAwsLogSourceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAwsLogSourceRef) AwsLogSourceRef() *AwsLogSourceReference {
	var returns *AwsLogSourceReference
	_jsii_.Get(
		j,
		"awsLogSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAwsLogSourceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAwsLogSourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


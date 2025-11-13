package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainNameV2.
// Experimental.
type IDomainNameV2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DomainNameV2 resource.
	// Experimental.
	DomainNameV2Ref() *DomainNameV2Reference
}

// The jsii proxy for IDomainNameV2Ref
type jsiiProxy_IDomainNameV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDomainNameV2Ref) DomainNameV2Ref() *DomainNameV2Reference {
	var returns *DomainNameV2Reference
	_jsii_.Get(
		j,
		"domainNameV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainNameV2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainNameV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package awsdirectoryservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdirectoryservice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MicrosoftAD.
// Experimental.
type IMicrosoftADRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MicrosoftAD resource.
	// Experimental.
	MicrosoftAdRef() *MicrosoftADReference
}

// The jsii proxy for IMicrosoftADRef
type jsiiProxy_IMicrosoftADRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMicrosoftADRef) MicrosoftAdRef() *MicrosoftADReference {
	var returns *MicrosoftADReference
	_jsii_.Get(
		j,
		"microsoftAdRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMicrosoftADRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMicrosoftADRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


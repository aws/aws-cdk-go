package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebExperience.
// Experimental.
type IWebExperienceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WebExperience resource.
	// Experimental.
	WebExperienceRef() *WebExperienceReference
}

// The jsii proxy for IWebExperienceRef
type jsiiProxy_IWebExperienceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IWebExperienceRef) WebExperienceRef() *WebExperienceReference {
	var returns *WebExperienceReference
	_jsii_.Get(
		j,
		"webExperienceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebExperienceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebExperienceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


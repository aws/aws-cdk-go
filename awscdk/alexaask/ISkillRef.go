package alexaask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/alexaask/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Skill.
// Experimental.
type ISkillRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Skill resource.
	// Experimental.
	SkillRef() *SkillReference
}

// The jsii proxy for ISkillRef
type jsiiProxy_ISkillRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISkillRef) SkillRef() *SkillReference {
	var returns *SkillReference
	_jsii_.Get(
		j,
		"skillRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISkillRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISkillRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


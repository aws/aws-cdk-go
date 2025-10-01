package alexaask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/alexaask/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Skill.
// Experimental.
type ISkillRef interface {
	constructs.IConstruct
	// A reference to a Skill resource.
	// Experimental.
	SkillRef() *SkillReference
}

// The jsii proxy for ISkillRef
type jsiiProxy_ISkillRef struct {
	internal.Type__constructsIConstruct
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


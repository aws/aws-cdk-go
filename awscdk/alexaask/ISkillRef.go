package alexaask

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/alexaask/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Skill.
// Experimental.
type ISkillRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISkillRef
type jsiiProxy_ISkillRef struct {
	internal.Type__constructsIConstruct
}


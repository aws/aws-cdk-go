package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebExperience.
// Experimental.
type IWebExperienceRef interface {
	constructs.IConstruct
	// A reference to a WebExperience resource.
	// Experimental.
	WebExperienceRef() *WebExperienceReference
}

// The jsii proxy for IWebExperienceRef
type jsiiProxy_IWebExperienceRef struct {
	internal.Type__constructsIConstruct
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


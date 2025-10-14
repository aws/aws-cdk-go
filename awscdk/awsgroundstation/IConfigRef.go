package awsgroundstation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Config.
// Experimental.
type IConfigRef interface {
	constructs.IConstruct
	// A reference to a Config resource.
	// Experimental.
	ConfigRef() *ConfigReference
}

// The jsii proxy for IConfigRef
type jsiiProxy_IConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigRef) ConfigRef() *ConfigReference {
	var returns *ConfigReference
	_jsii_.Get(
		j,
		"configRef",
		&returns,
	)
	return returns
}


package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Blueprint.
// Experimental.
type IBlueprintRef interface {
	constructs.IConstruct
	// A reference to a Blueprint resource.
	// Experimental.
	BlueprintRef() *BlueprintReference
}

// The jsii proxy for IBlueprintRef
type jsiiProxy_IBlueprintRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBlueprintRef) BlueprintRef() *BlueprintReference {
	var returns *BlueprintReference
	_jsii_.Get(
		j,
		"blueprintRef",
		&returns,
	)
	return returns
}


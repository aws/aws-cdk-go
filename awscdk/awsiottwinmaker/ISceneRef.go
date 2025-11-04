package awsiottwinmaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Scene.
// Experimental.
type ISceneRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Scene resource.
	// Experimental.
	SceneRef() *SceneReference
}

// The jsii proxy for ISceneRef
type jsiiProxy_ISceneRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISceneRef) SceneRef() *SceneReference {
	var returns *SceneReference
	_jsii_.Get(
		j,
		"sceneRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISceneRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISceneRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


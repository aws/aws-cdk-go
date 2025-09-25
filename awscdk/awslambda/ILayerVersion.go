package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

type ILayerVersion interface {
	ILayerVersionRef
	awscdk.IResource
	// Add permission for this layer version to specific entities.
	//
	// Usage within
	// the same account where the layer is defined is always allowed and does not
	// require calling this method. Note that the principal that creates the
	// Lambda function using the layer (for example, a CloudFormation changeset
	// execution role) also needs to have the ``lambda:GetLayerVersion``
	// permission on the layer version.
	AddPermission(id *string, permission *LayerVersionPermission)
	// The runtimes compatible with this Layer.
	// Default: - All supported runtimes. Setting this to Runtime.ALL is equivalent to leaving it undefined.
	//
	CompatibleRuntimes() *[]Runtime
	// The ARN of the Lambda Layer version that this Layer defines.
	LayerVersionArn() *string
}

// The jsii proxy for ILayerVersion
type jsiiProxy_ILayerVersion struct {
	jsiiProxy_ILayerVersionRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ILayerVersion) AddPermission(id *string, permission *LayerVersionPermission) {
	if err := i.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (i *jsiiProxy_ILayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ILayerVersion) CompatibleRuntimes() *[]Runtime {
	var returns *[]Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersion) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersion) LayerVersionRef() *LayerVersionReference {
	var returns *LayerVersionReference
	_jsii_.Get(
		j,
		"layerVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


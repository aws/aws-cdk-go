package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
)

type ILayerVersion interface {
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
	// Default: Runtime.All
	//
	CompatibleRuntimes() *[]Runtime
	// The ARN of the Lambda Layer version that this Layer defines.
	LayerVersionArn() *string
}

// The jsii proxy for ILayerVersion
type jsiiProxy_ILayerVersion struct {
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


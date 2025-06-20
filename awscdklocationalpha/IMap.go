package awscdklocationalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdklocationalpha/v2/internal"
)

// Represents the Amazon Location Service Map.
// Experimental.
type IMap interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the Map.
	// Experimental.
	MapArn() *string
	// The name of the map.
	// Experimental.
	MapName() *string
}

// The jsii proxy for IMap
type jsiiProxy_IMap struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IMap) MapArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMap) MapName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapName",
		&returns,
	)
	return returns
}


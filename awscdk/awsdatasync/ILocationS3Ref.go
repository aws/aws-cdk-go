package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationS3.
// Experimental.
type ILocationS3Ref interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LocationS3 resource.
	// Experimental.
	LocationS3Ref() *LocationS3Reference
}

// The jsii proxy for ILocationS3Ref
type jsiiProxy_ILocationS3Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILocationS3Ref) LocationS3Ref() *LocationS3Reference {
	var returns *LocationS3Reference
	_jsii_.Get(
		j,
		"locationS3Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationS3Ref) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationS3Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


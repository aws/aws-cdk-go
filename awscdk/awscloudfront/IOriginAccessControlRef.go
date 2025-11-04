package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginAccessControl.
// Experimental.
type IOriginAccessControlRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OriginAccessControl resource.
	// Experimental.
	OriginAccessControlRef() *OriginAccessControlReference
}

// The jsii proxy for IOriginAccessControlRef
type jsiiProxy_IOriginAccessControlRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IOriginAccessControlRef) OriginAccessControlRef() *OriginAccessControlReference {
	var returns *OriginAccessControlReference
	_jsii_.Get(
		j,
		"originAccessControlRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessControlRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessControlRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


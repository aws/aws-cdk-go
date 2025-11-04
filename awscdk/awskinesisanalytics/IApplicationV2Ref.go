package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Application.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type IApplicationV2Ref interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Application resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationRef() *ApplicationV2Reference
}

// The jsii proxy for IApplicationV2Ref
type jsiiProxy_IApplicationV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApplicationV2Ref) ApplicationRef() *ApplicationV2Reference {
	var returns *ApplicationV2Reference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationV2Ref) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationOutput.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type IApplicationOutputV2Ref interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApplicationOutput resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationOutputRef() *ApplicationOutputV2Reference
}

// The jsii proxy for IApplicationOutputV2Ref
type jsiiProxy_IApplicationOutputV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApplicationOutputV2Ref) ApplicationOutputRef() *ApplicationOutputV2Reference {
	var returns *ApplicationOutputV2Reference
	_jsii_.Get(
		j,
		"applicationOutputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationOutputV2Ref) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationOutputV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


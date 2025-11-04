package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FargateProfile.
// Experimental.
type IFargateProfileRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FargateProfile resource.
	// Experimental.
	FargateProfileRef() *FargateProfileReference
}

// The jsii proxy for IFargateProfileRef
type jsiiProxy_IFargateProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFargateProfileRef) FargateProfileRef() *FargateProfileReference {
	var returns *FargateProfileReference
	_jsii_.Get(
		j,
		"fargateProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFargateProfileRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFargateProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


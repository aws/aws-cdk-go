package awscontroltower

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LandingZone.
// Experimental.
type ILandingZoneRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LandingZone resource.
	// Experimental.
	LandingZoneRef() *LandingZoneReference
}

// The jsii proxy for ILandingZoneRef
type jsiiProxy_ILandingZoneRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILandingZoneRef) LandingZoneRef() *LandingZoneReference {
	var returns *LandingZoneReference
	_jsii_.Get(
		j,
		"landingZoneRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILandingZoneRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILandingZoneRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


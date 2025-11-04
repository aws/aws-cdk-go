package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SoftwarePackage.
// Experimental.
type ISoftwarePackageRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SoftwarePackage resource.
	// Experimental.
	SoftwarePackageRef() *SoftwarePackageReference
}

// The jsii proxy for ISoftwarePackageRef
type jsiiProxy_ISoftwarePackageRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISoftwarePackageRef) SoftwarePackageRef() *SoftwarePackageReference {
	var returns *SoftwarePackageReference
	_jsii_.Get(
		j,
		"softwarePackageRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoftwarePackageRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoftwarePackageRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


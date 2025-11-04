package awspanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspanorama/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Package.
// Experimental.
type IPackageRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Package resource.
	// Experimental.
	PackageRef() *PackageReference
}

// The jsii proxy for IPackageRef
type jsiiProxy_IPackageRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPackageRef) PackageRef() *PackageReference {
	var returns *PackageReference
	_jsii_.Get(
		j,
		"packageRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackageRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackageRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package awspanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspanorama/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Package.
// Experimental.
type IPackageRef interface {
	constructs.IConstruct
	// A reference to a Package resource.
	// Experimental.
	PackageRef() *PackageReference
}

// The jsii proxy for IPackageRef
type jsiiProxy_IPackageRef struct {
	internal.Type__constructsIConstruct
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


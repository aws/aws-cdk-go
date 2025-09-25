package awspanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspanorama/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackageVersion.
// Experimental.
type IPackageVersionRef interface {
	constructs.IConstruct
	// A reference to a PackageVersion resource.
	// Experimental.
	PackageVersionRef() *PackageVersionReference
}

// The jsii proxy for IPackageVersionRef
type jsiiProxy_IPackageVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPackageVersionRef) PackageVersionRef() *PackageVersionReference {
	var returns *PackageVersionReference
	_jsii_.Get(
		j,
		"packageVersionRef",
		&returns,
	)
	return returns
}


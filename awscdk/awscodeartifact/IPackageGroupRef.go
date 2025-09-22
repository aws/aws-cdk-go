package awscodeartifact

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeartifact/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackageGroup.
// Experimental.
type IPackageGroupRef interface {
	constructs.IConstruct
	// A reference to a PackageGroup resource.
	// Experimental.
	PackageGroupRef() *PackageGroupReference
}

// The jsii proxy for IPackageGroupRef
type jsiiProxy_IPackageGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPackageGroupRef) PackageGroupRef() *PackageGroupReference {
	var returns *PackageGroupReference
	_jsii_.Get(
		j,
		"packageGroupRef",
		&returns,
	)
	return returns
}


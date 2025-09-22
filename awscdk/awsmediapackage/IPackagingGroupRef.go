package awsmediapackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackagingGroup.
// Experimental.
type IPackagingGroupRef interface {
	constructs.IConstruct
	// A reference to a PackagingGroup resource.
	// Experimental.
	PackagingGroupRef() *PackagingGroupReference
}

// The jsii proxy for IPackagingGroupRef
type jsiiProxy_IPackagingGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPackagingGroupRef) PackagingGroupRef() *PackagingGroupReference {
	var returns *PackagingGroupReference
	_jsii_.Get(
		j,
		"packagingGroupRef",
		&returns,
	)
	return returns
}


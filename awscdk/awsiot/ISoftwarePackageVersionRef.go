package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SoftwarePackageVersion.
// Experimental.
type ISoftwarePackageVersionRef interface {
	constructs.IConstruct
	// A reference to a SoftwarePackageVersion resource.
	// Experimental.
	SoftwarePackageVersionRef() *SoftwarePackageVersionReference
}

// The jsii proxy for ISoftwarePackageVersionRef
type jsiiProxy_ISoftwarePackageVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISoftwarePackageVersionRef) SoftwarePackageVersionRef() *SoftwarePackageVersionReference {
	var returns *SoftwarePackageVersionReference
	_jsii_.Get(
		j,
		"softwarePackageVersionRef",
		&returns,
	)
	return returns
}


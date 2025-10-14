package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationS3.
// Experimental.
type ILocationS3Ref interface {
	constructs.IConstruct
	// A reference to a LocationS3 resource.
	// Experimental.
	LocationS3Ref() *LocationS3Reference
}

// The jsii proxy for ILocationS3Ref
type jsiiProxy_ILocationS3Ref struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocationS3Ref) LocationS3Ref() *LocationS3Reference {
	var returns *LocationS3Reference
	_jsii_.Get(
		j,
		"locationS3Ref",
		&returns,
	)
	return returns
}


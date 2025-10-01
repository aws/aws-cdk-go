package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Database.
// Experimental.
type IDatabaseRef interface {
	constructs.IConstruct
	// A reference to a Database resource.
	// Experimental.
	DatabaseRef() *DatabaseReference
}

// The jsii proxy for IDatabaseRef
type jsiiProxy_IDatabaseRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDatabaseRef) DatabaseRef() *DatabaseReference {
	var returns *DatabaseReference
	_jsii_.Get(
		j,
		"databaseRef",
		&returns,
	)
	return returns
}


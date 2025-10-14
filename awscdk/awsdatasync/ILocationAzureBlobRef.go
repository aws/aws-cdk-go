package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationAzureBlob.
// Experimental.
type ILocationAzureBlobRef interface {
	constructs.IConstruct
	// A reference to a LocationAzureBlob resource.
	// Experimental.
	LocationAzureBlobRef() *LocationAzureBlobReference
}

// The jsii proxy for ILocationAzureBlobRef
type jsiiProxy_ILocationAzureBlobRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocationAzureBlobRef) LocationAzureBlobRef() *LocationAzureBlobReference {
	var returns *LocationAzureBlobReference
	_jsii_.Get(
		j,
		"locationAzureBlobRef",
		&returns,
	)
	return returns
}


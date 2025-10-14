package awshealthlake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awshealthlake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FHIRDatastore.
// Experimental.
type IFHIRDatastoreRef interface {
	constructs.IConstruct
	// A reference to a FHIRDatastore resource.
	// Experimental.
	FhirDatastoreRef() *FHIRDatastoreReference
}

// The jsii proxy for IFHIRDatastoreRef
type jsiiProxy_IFHIRDatastoreRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFHIRDatastoreRef) FhirDatastoreRef() *FHIRDatastoreReference {
	var returns *FHIRDatastoreReference
	_jsii_.Get(
		j,
		"fhirDatastoreRef",
		&returns,
	)
	return returns
}


package awsappintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappintegrations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataIntegration.
// Experimental.
type IDataIntegrationRef interface {
	constructs.IConstruct
	// A reference to a DataIntegration resource.
	// Experimental.
	DataIntegrationRef() *DataIntegrationReference
}

// The jsii proxy for IDataIntegrationRef
type jsiiProxy_IDataIntegrationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataIntegrationRef) DataIntegrationRef() *DataIntegrationReference {
	var returns *DataIntegrationReference
	_jsii_.Get(
		j,
		"dataIntegrationRef",
		&returns,
	)
	return returns
}


package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IngestConfiguration.
// Experimental.
type IIngestConfigurationRef interface {
	constructs.IConstruct
	// A reference to a IngestConfiguration resource.
	// Experimental.
	IngestConfigurationRef() *IngestConfigurationReference
}

// The jsii proxy for IIngestConfigurationRef
type jsiiProxy_IIngestConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIngestConfigurationRef) IngestConfigurationRef() *IngestConfigurationReference {
	var returns *IngestConfigurationReference
	_jsii_.Get(
		j,
		"ingestConfigurationRef",
		&returns,
	)
	return returns
}


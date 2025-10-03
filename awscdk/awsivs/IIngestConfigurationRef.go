package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IngestConfiguration.
// Experimental.
type IIngestConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIngestConfigurationRef
type jsiiProxy_IIngestConfigurationRef struct {
	internal.Type__constructsIConstruct
}


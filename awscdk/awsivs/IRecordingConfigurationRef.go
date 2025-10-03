package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecordingConfiguration.
// Experimental.
type IRecordingConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRecordingConfigurationRef
type jsiiProxy_IRecordingConfigurationRef struct {
	internal.Type__constructsIConstruct
}


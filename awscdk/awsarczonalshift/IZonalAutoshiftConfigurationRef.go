package awsarczonalshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsarczonalshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ZonalAutoshiftConfiguration.
// Experimental.
type IZonalAutoshiftConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IZonalAutoshiftConfigurationRef
type jsiiProxy_IZonalAutoshiftConfigurationRef struct {
	internal.Type__constructsIConstruct
}


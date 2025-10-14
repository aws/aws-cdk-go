package awsarczonalshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsarczonalshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ZonalAutoshiftConfiguration.
// Experimental.
type IZonalAutoshiftConfigurationRef interface {
	constructs.IConstruct
	// A reference to a ZonalAutoshiftConfiguration resource.
	// Experimental.
	ZonalAutoshiftConfigurationRef() *ZonalAutoshiftConfigurationReference
}

// The jsii proxy for IZonalAutoshiftConfigurationRef
type jsiiProxy_IZonalAutoshiftConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IZonalAutoshiftConfigurationRef) ZonalAutoshiftConfigurationRef() *ZonalAutoshiftConfigurationReference {
	var returns *ZonalAutoshiftConfigurationReference
	_jsii_.Get(
		j,
		"zonalAutoshiftConfigurationRef",
		&returns,
	)
	return returns
}


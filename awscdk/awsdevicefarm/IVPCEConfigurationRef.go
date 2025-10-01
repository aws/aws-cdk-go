package awsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEConfiguration.
// Experimental.
type IVPCEConfigurationRef interface {
	constructs.IConstruct
	// A reference to a VPCEConfiguration resource.
	// Experimental.
	VpceConfigurationRef() *VPCEConfigurationReference
}

// The jsii proxy for IVPCEConfigurationRef
type jsiiProxy_IVPCEConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCEConfigurationRef) VpceConfigurationRef() *VPCEConfigurationReference {
	var returns *VPCEConfigurationReference
	_jsii_.Get(
		j,
		"vpceConfigurationRef",
		&returns,
	)
	return returns
}


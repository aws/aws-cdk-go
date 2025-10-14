package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IpAccessSettings.
// Experimental.
type IIpAccessSettingsRef interface {
	constructs.IConstruct
	// A reference to a IpAccessSettings resource.
	// Experimental.
	IpAccessSettingsRef() *IpAccessSettingsReference
}

// The jsii proxy for IIpAccessSettingsRef
type jsiiProxy_IIpAccessSettingsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIpAccessSettingsRef) IpAccessSettingsRef() *IpAccessSettingsReference {
	var returns *IpAccessSettingsReference
	_jsii_.Get(
		j,
		"ipAccessSettingsRef",
		&returns,
	)
	return returns
}


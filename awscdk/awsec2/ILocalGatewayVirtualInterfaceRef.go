package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayVirtualInterface.
// Experimental.
type ILocalGatewayVirtualInterfaceRef interface {
	constructs.IConstruct
	// A reference to a LocalGatewayVirtualInterface resource.
	// Experimental.
	LocalGatewayVirtualInterfaceRef() *LocalGatewayVirtualInterfaceReference
}

// The jsii proxy for ILocalGatewayVirtualInterfaceRef
type jsiiProxy_ILocalGatewayVirtualInterfaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocalGatewayVirtualInterfaceRef) LocalGatewayVirtualInterfaceRef() *LocalGatewayVirtualInterfaceReference {
	var returns *LocalGatewayVirtualInterfaceReference
	_jsii_.Get(
		j,
		"localGatewayVirtualInterfaceRef",
		&returns,
	)
	return returns
}


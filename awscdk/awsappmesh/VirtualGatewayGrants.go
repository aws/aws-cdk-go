package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
)

// Collection of grant methods for a IVirtualGatewayRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var virtualGatewayRef IVirtualGatewayRef
//
//   virtualGatewayGrants := awscdk.Aws_appmesh.VirtualGatewayGrants_FromVirtualGateway(virtualGatewayRef)
//
type VirtualGatewayGrants interface {
	Resource() interfacesawsappmesh.IVirtualGatewayRef
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	StreamAggregatedResources(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for VirtualGatewayGrants
type jsiiProxy_VirtualGatewayGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_VirtualGatewayGrants) Resource() interfacesawsappmesh.IVirtualGatewayRef {
	var returns interfacesawsappmesh.IVirtualGatewayRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for VirtualGatewayGrants.
func VirtualGatewayGrants_FromVirtualGateway(resource interfacesawsappmesh.IVirtualGatewayRef) VirtualGatewayGrants {
	_init_.Initialize()

	if err := validateVirtualGatewayGrants_FromVirtualGatewayParameters(resource); err != nil {
		panic(err)
	}
	var returns VirtualGatewayGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayGrants",
		"fromVirtualGateway",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualGatewayGrants) StreamAggregatedResources(grantee awsiam.IGrantable) awsiam.Grant {
	if err := v.validateStreamAggregatedResourcesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"streamAggregatedResources",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}


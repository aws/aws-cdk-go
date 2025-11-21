package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
)

// Collection of grant methods for a IVirtualNodeRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var virtualNodeRef IVirtualNodeRef
//
//   virtualNodeGrants := awscdk.Aws_appmesh.VirtualNodeGrants_FromVirtualNode(virtualNodeRef)
//
type VirtualNodeGrants interface {
	Resource() interfacesawsappmesh.IVirtualNodeRef
	// Grants streamAggregatedResources permissions.
	StreamAggregatedResources(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for VirtualNodeGrants
type jsiiProxy_VirtualNodeGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_VirtualNodeGrants) Resource() interfacesawsappmesh.IVirtualNodeRef {
	var returns interfacesawsappmesh.IVirtualNodeRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for VirtualNodeGrants.
func VirtualNodeGrants_FromVirtualNode(resource interfacesawsappmesh.IVirtualNodeRef) VirtualNodeGrants {
	_init_.Initialize()

	if err := validateVirtualNodeGrants_FromVirtualNodeParameters(resource); err != nil {
		panic(err)
	}
	var returns VirtualNodeGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualNodeGrants",
		"fromVirtualNode",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNodeGrants) StreamAggregatedResources(grantee awsiam.IGrantable) awsiam.Grant {
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


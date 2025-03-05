package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVirtualService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualServiceProps := &CfnVirtualServiceProps{
//   	MeshName: jsii.String("meshName"),
//   	Spec: &VirtualServiceSpecProperty{
//   		Provider: &VirtualServiceProviderProperty{
//   			VirtualNode: &VirtualNodeServiceProviderProperty{
//   				VirtualNodeName: jsii.String("virtualNodeName"),
//   			},
//   			VirtualRouter: &VirtualRouterServiceProviderProperty{
//   				VirtualRouterName: jsii.String("virtualRouterName"),
//   			},
//   		},
//   	},
//   	VirtualServiceName: jsii.String("virtualServiceName"),
//
//   	// the properties below are optional
//   	MeshOwner: jsii.String("meshOwner"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html
//
type CfnVirtualServiceProps struct {
	// The name of the service mesh to create the virtual service in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-meshname
	//
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The virtual service specification to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-spec
	//
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The name to use for the virtual service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-virtualservicename
	//
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-meshowner
	//
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual service to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


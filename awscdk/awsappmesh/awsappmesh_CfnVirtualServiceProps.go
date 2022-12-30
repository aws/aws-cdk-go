package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVirtualService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVirtualServiceProps := &cfnVirtualServiceProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &virtualServiceSpecProperty{
//   		provider: &virtualServiceProviderProperty{
//   			virtualNode: &virtualNodeServiceProviderProperty{
//   				virtualNodeName: jsii.String("virtualNodeName"),
//   			},
//   			virtualRouter: &virtualRouterServiceProviderProperty{
//   				virtualRouterName: jsii.String("virtualRouterName"),
//   			},
//   		},
//   	},
//   	virtualServiceName: jsii.String("virtualServiceName"),
//
//   	// the properties below are optional
//   	meshOwner: jsii.String("meshOwner"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVirtualServiceProps struct {
	// The name of the service mesh to create the virtual service in.
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The virtual service specification to apply.
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The name to use for the virtual service.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the virtual service to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


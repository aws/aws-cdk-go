package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMesh`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMeshProps := &CfnMeshProps{
//   	MeshName: jsii.String("meshName"),
//   	Spec: &MeshSpecProperty{
//   		EgressFilter: &EgressFilterProperty{
//   			Type: jsii.String("type"),
//   		},
//   		ServiceDiscovery: &MeshServiceDiscoveryProperty{
//   			IpPreference: jsii.String("ipPreference"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html
//
type CfnMeshProps struct {
	// The name to use for the service mesh.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html#cfn-appmesh-mesh-meshname
	//
	MeshName *string `field:"optional" json:"meshName" yaml:"meshName"`
	// The service mesh specification to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html#cfn-appmesh-mesh-spec
	//
	Spec interface{} `field:"optional" json:"spec" yaml:"spec"`
	// Optional metadata that you can apply to the service mesh to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html#cfn-appmesh-mesh-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


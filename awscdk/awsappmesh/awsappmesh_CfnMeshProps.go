package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnMesh`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMeshProps := &cfnMeshProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &meshSpecProperty{
//   		egressFilter: &egressFilterProperty{
//   			type: jsii.String("type"),
//   		},
//   		serviceDiscovery: &meshServiceDiscoveryProperty{
//   			ipPreference: jsii.String("ipPreference"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMeshProps struct {
	// The name to use for the service mesh.
	MeshName *string `field:"optional" json:"meshName" yaml:"meshName"`
	// The service mesh specification to apply.
	Spec interface{} `field:"optional" json:"spec" yaml:"spec"`
	// Optional metadata that you can apply to the service mesh to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


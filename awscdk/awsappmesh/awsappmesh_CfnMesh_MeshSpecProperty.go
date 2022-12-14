package awsappmesh


// An object that represents the specification of a service mesh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   meshSpecProperty := &meshSpecProperty{
//   	egressFilter: &egressFilterProperty{
//   		type: jsii.String("type"),
//   	},
//   	serviceDiscovery: &meshServiceDiscoveryProperty{
//   		ipPreference: jsii.String("ipPreference"),
//   	},
//   }
//
type CfnMesh_MeshSpecProperty struct {
	// The egress filter rules for the service mesh.
	EgressFilter interface{} `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// An object that represents the service discovery information for a service mesh.
	ServiceDiscovery interface{} `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}


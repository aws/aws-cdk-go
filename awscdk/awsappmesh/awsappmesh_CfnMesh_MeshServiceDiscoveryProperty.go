package awsappmesh


// An object that represents the service discovery information for a service mesh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   meshServiceDiscoveryProperty := &meshServiceDiscoveryProperty{
//   	ipPreference: jsii.String("ipPreference"),
//   }
//
type CfnMesh_MeshServiceDiscoveryProperty struct {
	// The IP version to use to control traffic within the mesh.
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}


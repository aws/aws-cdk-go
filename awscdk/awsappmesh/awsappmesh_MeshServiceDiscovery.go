package awsappmesh


// Properties for Mesh Service Discovery.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &meshProps{
//   	meshName: jsii.String("myAwsMesh"),
//   	serviceDiscovery: &meshServiceDiscovery{
//   		ipPreference: appmesh.ipPreference_IPV4_ONLY,
//   	},
//   })
//
type MeshServiceDiscovery struct {
	// IP preference applied to all Virtual Nodes in the Mesh.
	IpPreference IpPreference `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}


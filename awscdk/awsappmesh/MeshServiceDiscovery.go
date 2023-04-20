package awsappmesh


// Properties for Mesh Service Discovery.
//
// Example:
//   mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &MeshProps{
//   	MeshName: jsii.String("myAwsMesh"),
//   	ServiceDiscovery: &MeshServiceDiscovery{
//   		IpPreference: appmesh.IpPreference_IPV4_ONLY,
//   	},
//   })
//
type MeshServiceDiscovery struct {
	// IP preference applied to all Virtual Nodes in the Mesh.
	IpPreference IpPreference `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}


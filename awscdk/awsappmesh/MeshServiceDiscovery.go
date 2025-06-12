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
	// Default: - No IP preference is applied to any of the Virtual Nodes in the Mesh.
	// Virtual Nodes without an IP preference will have the following configured.
	// Envoy listeners are configured to bind only to IPv4.
	// Envoy will use IPv4 when sending traffic to a local application.
	// For DNS service discovery, the Envoy DNS resolver to prefer using IPv6 and fall back to IPv4.
	// For CloudMap service discovery, App Mesh will prefer using IPv4 and fall back to IPv6 for IPs returned by CloudMap.
	//
	IpPreference IpPreference `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}


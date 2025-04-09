package awsappmesh


// Enum of supported IP preferences.
//
// Used to dictate the IP version for mesh wide and virtual node service discovery.
// Also used to specify the IP version that a sidecar Envoy uses when sending traffic to a local application.
//
// Example:
//   mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &MeshProps{
//   	MeshName: jsii.String("myAwsMesh"),
//   	ServiceDiscovery: &MeshServiceDiscovery{
//   		IpPreference: appmesh.IpPreference_IPV4_ONLY,
//   	},
//   })
//
type IpPreference string

const (
	// Use IPv4 when sending traffic to a local application.
	//
	// Only use IPv4 for service discovery.
	IpPreference_IPV4_ONLY IpPreference = "IPV4_ONLY"
	// Use IPv4 when sending traffic to a local application.
	//
	// First attempt to use IPv4 and fall back to IPv6 for service discovery.
	IpPreference_IPV4_PREFERRED IpPreference = "IPV4_PREFERRED"
	// Use IPv6 when sending traffic to a local application.
	//
	// Only use IPv6 for service discovery.
	IpPreference_IPV6_ONLY IpPreference = "IPV6_ONLY"
	// Use IPv6 when sending traffic to a local application.
	//
	// First attempt to use IPv6 and fall back to IPv4 for service discovery.
	IpPreference_IPV6_PREFERRED IpPreference = "IPV6_PREFERRED"
)


package awsappmesh


// The set of properties used when creating a Mesh.
//
// Example:
//   var infraStack Stack
//   var appStack Stack
//
//
//   mesh := appmesh.NewMesh(infraStack, jsii.String("AppMesh"), &MeshProps{
//   	MeshName: jsii.String("myAwsMesh"),
//   	EgressFilter: appmesh.MeshFilterType_ALLOW_ALL,
//   })
//
//   // the VirtualRouter will belong to 'appStack',
//   // even though the Mesh belongs to 'infraStack'
//   router := appmesh.NewVirtualRouter(appStack, jsii.String("router"), &VirtualRouterProps{
//   	Mesh: Mesh,
//   	 // notice that mesh is a required property when creating a router with the 'new' statement
//   	Listeners: []VirtualRouterListener{
//   		appmesh.VirtualRouterListener_Http(jsii.Number(8081)),
//   	},
//   })
//
type MeshProps struct {
	// Egress filter to be applied to the Mesh.
	// Default: DROP_ALL.
	//
	EgressFilter MeshFilterType `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// The name of the Mesh being defined.
	// Default: - A name is automatically generated.
	//
	MeshName *string `field:"optional" json:"meshName" yaml:"meshName"`
	// Defines how upstream clients will discover VirtualNodes in the Mesh.
	// Default: - No Service Discovery.
	//
	ServiceDiscovery *MeshServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}


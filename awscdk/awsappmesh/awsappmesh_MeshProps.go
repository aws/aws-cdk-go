package awsappmesh


// The set of properties used when creating a Mesh.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var infraStack stack
//   var appStack stack
//
//
//   mesh := appmesh.NewMesh(infraStack, jsii.String("AppMesh"), &meshProps{
//   	meshName: jsii.String("myAwsMesh"),
//   	egressFilter: appmesh.meshFilterType_ALLOW_ALL,
//   })
//
//   // the VirtualRouter will belong to 'appStack',
//   // even though the Mesh belongs to 'infraStack'
//   router := appmesh.NewVirtualRouter(appStack, jsii.String("router"), &virtualRouterProps{
//   	mesh: mesh,
//   	 // notice that mesh is a required property when creating a router with the 'new' statement
//   	listeners: []virtualRouterListener{
//   		appmesh.*virtualRouterListener.http(jsii.Number(8081)),
//   	},
//   })
//
type MeshProps struct {
	// Egress filter to be applied to the Mesh.
	EgressFilter MeshFilterType `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// The name of the Mesh being defined.
	MeshName *string `field:"optional" json:"meshName" yaml:"meshName"`
	// Defines how upstream clients will discover VirtualNodes in the Mesh.
	ServiceDiscovery *MeshServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}


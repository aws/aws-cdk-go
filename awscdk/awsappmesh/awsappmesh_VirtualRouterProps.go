package awsappmesh


// The properties used when creating a new VirtualRouter.
//
// Example:
//   var infraStack stack
//   var appStack stack
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
//   	Listeners: []virtualRouterListener{
//   		appmesh.*virtualRouterListener_Http(jsii.Number(8081)),
//   	},
//   })
//
// Experimental.
type VirtualRouterProps struct {
	// Listener specification for the VirtualRouter.
	// Experimental.
	Listeners *[]VirtualRouterListener `field:"optional" json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	// Experimental.
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
	// The Mesh which the VirtualRouter belongs to.
	// Experimental.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
}


package awsappmesh


// The properties used when creating a new VirtualRouter.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type VirtualRouterProps struct {
	// Listener specification for the VirtualRouter.
	Listeners *[]VirtualRouterListener `field:"optional" json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
	// The Mesh which the VirtualRouter belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
}


package awsappmesh


// The properties used when creating a new VirtualRouter.
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
type VirtualRouterProps struct {
	// Listener specification for the VirtualRouter.
	Listeners *[]VirtualRouterListener `field:"optional" json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
	// The Mesh which the VirtualRouter belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
}


package awsappmesh


// Interface with base properties all routers willl inherit.
//
// Example:
//   var mesh Mesh
//
//   router := mesh.addVirtualRouter(jsii.String("router"), &VirtualRouterBaseProps{
//   	Listeners: []VirtualRouterListener{
//   		appmesh.VirtualRouterListener_Http(jsii.Number(8080)),
//   	},
//   })
//
type VirtualRouterBaseProps struct {
	// Listener specification for the VirtualRouter.
	// Default: - A listener on HTTP port 8080.
	//
	Listeners *[]VirtualRouterListener `field:"optional" json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	// Default: - A name is automatically determined.
	//
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
}


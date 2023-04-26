package awsappmesh


// Interface with base properties all routers willl inherit.
//
// Example:
//   var mesh mesh
//
//   router := mesh.addVirtualRouter(jsii.String("router"), &VirtualRouterBaseProps{
//   	Listeners: []virtualRouterListener{
//   		appmesh.*virtualRouterListener_Http(jsii.Number(8080)),
//   	},
//   })
//
// Experimental.
type VirtualRouterBaseProps struct {
	// Listener specification for the VirtualRouter.
	// Experimental.
	Listeners *[]VirtualRouterListener `field:"optional" json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	// Experimental.
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
}


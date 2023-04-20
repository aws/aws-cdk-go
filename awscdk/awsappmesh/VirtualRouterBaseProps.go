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
type VirtualRouterBaseProps struct {
	// Listener specification for the VirtualRouter.
	Listeners *[]VirtualRouterListener `field:"optional" json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
}


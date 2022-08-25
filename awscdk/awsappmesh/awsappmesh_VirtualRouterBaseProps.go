package awsappmesh


// Interface with base properties all routers willl inherit.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//   router := mesh.addVirtualRouter(jsii.String("router"), &virtualRouterBaseProps{
//   	listeners: []virtualRouterListener{
//   		appmesh.*virtualRouterListener.http(jsii.Number(8080)),
//   	},
//   })
//
type VirtualRouterBaseProps struct {
	// Listener specification for the VirtualRouter.
	Listeners *[]VirtualRouterListener `field:"optional" json:"listeners" yaml:"listeners"`
	// The name of the VirtualRouter.
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
}


package awsappmesh


// The properties applied to the VirtualService being defined.
//
// Example:
//   var mesh Mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
//   })
//
//   virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &VirtualServiceProps{
//   	VirtualServiceProvider: appmesh.VirtualServiceProvider_VirtualNode(node),
//   	VirtualServiceName: jsii.String("service1.domain.local"),
//   })
//
//   node.AddBackend(appmesh.Backend_VirtualService(virtualService))
//
type VirtualServiceProps struct {
	// The VirtualNode or VirtualRouter which the VirtualService uses as its provider.
	VirtualServiceProvider VirtualServiceProvider `field:"required" json:"virtualServiceProvider" yaml:"virtualServiceProvider"`
	// The name of the VirtualService.
	//
	// It is recommended this follows the fully-qualified domain name format,
	// such as "my-service.default.svc.cluster.local".
	//
	// Example value: `service.domain.local`
	// Default: - A name is automatically generated.
	//
	VirtualServiceName *string `field:"optional" json:"virtualServiceName" yaml:"virtualServiceName"`
}


package awsappmesh


// The properties applied to the VirtualService being defined.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
//   })
//
//   virtualService := appmesh.NewVirtualService(this, jsii.String("service-1"), &virtualServiceProps{
//   	virtualServiceProvider: appmesh.virtualServiceProvider.virtualNode(node),
//   	virtualServiceName: jsii.String("service1.domain.local"),
//   })
//
//   node.addBackend(appmesh.backend.virtualService(virtualService))
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
	VirtualServiceName *string `field:"optional" json:"virtualServiceName" yaml:"virtualServiceName"`
}


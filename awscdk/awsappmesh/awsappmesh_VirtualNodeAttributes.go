package awsappmesh


// Interface with properties necessary to import a reusable VirtualNode.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   virtualNodeName := "my-virtual-node"
//   appmesh.virtualNode.fromVirtualNodeAttributes(this, jsii.String("imported-virtual-node"), &virtualNodeAttributes{
//   	mesh: appmesh.mesh.fromMeshName(this, jsii.String("Mesh"), jsii.String("testMesh")),
//   	virtualNodeName: virtualNodeName,
//   })
//
type VirtualNodeAttributes struct {
	// The Mesh that the VirtualNode belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// The name of the VirtualNode.
	VirtualNodeName *string `field:"required" json:"virtualNodeName" yaml:"virtualNodeName"`
}


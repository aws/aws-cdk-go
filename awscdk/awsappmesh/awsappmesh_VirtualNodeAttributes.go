package awsappmesh


// Interface with properties necessary to import a reusable VirtualNode.
//
// Example:
//   virtualNodeName := "my-virtual-node"
//   appmesh.virtualNode.fromVirtualNodeAttributes(this, jsii.String("imported-virtual-node"), &virtualNodeAttributes{
//   	mesh: appmesh.mesh.fromMeshName(this, jsii.String("Mesh"), jsii.String("testMesh")),
//   	virtualNodeName: virtualNodeName,
//   })
//
// Experimental.
type VirtualNodeAttributes struct {
	// The Mesh that the VirtualNode belongs to.
	// Experimental.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// The name of the VirtualNode.
	// Experimental.
	VirtualNodeName *string `field:"required" json:"virtualNodeName" yaml:"virtualNodeName"`
}


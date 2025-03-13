package awsappmesh


// Interface with properties necessary to import a reusable VirtualNode.
//
// Example:
//   virtualNodeName := "my-virtual-node"
//   appmesh.VirtualNode_FromVirtualNodeAttributes(this, jsii.String("imported-virtual-node"), &VirtualNodeAttributes{
//   	Mesh: appmesh.Mesh_FromMeshName(this, jsii.String("Mesh"), jsii.String("testMesh")),
//   	VirtualNodeName: virtualNodeName,
//   })
//
type VirtualNodeAttributes struct {
	// The Mesh that the VirtualNode belongs to.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// The name of the VirtualNode.
	VirtualNodeName *string `field:"required" json:"virtualNodeName" yaml:"virtualNodeName"`
}


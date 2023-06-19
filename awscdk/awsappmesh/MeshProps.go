package awsappmesh


// The set of properties used when creating a Mesh.
//
// Example:
//   mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &MeshProps{
//   	MeshName: jsii.String("myAwsMesh"),
//   	EgressFilter: appmesh.MeshFilterType_ALLOW_ALL,
//   })
//
// Experimental.
type MeshProps struct {
	// Egress filter to be applied to the Mesh.
	// Experimental.
	EgressFilter MeshFilterType `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// The name of the Mesh being defined.
	// Experimental.
	MeshName *string `field:"optional" json:"meshName" yaml:"meshName"`
}


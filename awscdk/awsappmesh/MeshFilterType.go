package awsappmesh


// A utility enum defined for the egressFilter type property, the default of DROP_ALL, allows traffic only to other resources inside the mesh, or API calls to amazon resources.
//
// Example:
//   mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &MeshProps{
//   	MeshName: jsii.String("myAwsMesh"),
//   	EgressFilter: appmesh.MeshFilterType_ALLOW_ALL,
//   })
//
// Experimental.
type MeshFilterType string

const (
	// Allows all outbound traffic.
	// Experimental.
	MeshFilterType_ALLOW_ALL MeshFilterType = "ALLOW_ALL"
	// Allows traffic only to other resources inside the mesh, or API calls to amazon resources.
	// Experimental.
	MeshFilterType_DROP_ALL MeshFilterType = "DROP_ALL"
)


package awsappmesh


// A utility enum defined for the egressFilter type property, the default of DROP_ALL, allows traffic only to other resources inside the mesh, or API calls to amazon resources.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   mesh := appmesh.NewMesh(this, jsii.String("AppMesh"), &meshProps{
//   	meshName: jsii.String("myAwsMesh"),
//   	egressFilter: appmesh.meshFilterType_ALLOW_ALL,
//   })
//
type MeshFilterType string

const (
	// Allows all outbound traffic.
	MeshFilterType_ALLOW_ALL MeshFilterType = "ALLOW_ALL"
	// Allows traffic only to other resources inside the mesh, or API calls to amazon resources.
	MeshFilterType_DROP_ALL MeshFilterType = "DROP_ALL"
)


package awsimagebuilderalpha


// The type of the container being used in the container recipe.
// Experimental.
type ContainerType string

const (
	// Indicates the container recipe uses a Docker container.
	// Experimental.
	ContainerType_DOCKER ContainerType = "DOCKER"
)


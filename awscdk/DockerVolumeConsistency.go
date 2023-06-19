package awscdk


// Supported Docker volume consistency types.
//
// Only valid on macOS due to the way file storage works on Mac.
// Experimental.
type DockerVolumeConsistency string

const (
	// Read/write operations inside the Docker container are applied immediately on the mounted host machine volumes.
	// Experimental.
	DockerVolumeConsistency_CONSISTENT DockerVolumeConsistency = "CONSISTENT"
	// Read/write operations on mounted Docker volumes are first written inside the container and then synchronized to the host machine.
	// Experimental.
	DockerVolumeConsistency_DELEGATED DockerVolumeConsistency = "DELEGATED"
	// Read/write operations on mounted Docker volumes are first applied on the host machine and then synchronized to the container.
	// Experimental.
	DockerVolumeConsistency_CACHED DockerVolumeConsistency = "CACHED"
)


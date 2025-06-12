package awsbatch


// Determines when the image is pulled from the registry to launch a container.
type ImagePullPolicy string

const (
	// Every time the kubelet launches a container, the kubelet queries the container image registry to resolve the name to an image digest.
	//
	// If the kubelet has a container image with that exact digest cached locally,
	// the kubelet uses its cached image; otherwise, the kubelet pulls the image with the resolved digest,
	// and uses that image to launch the container.
	// See: https://docs.docker.com/engine/reference/commandline/pull/#pull-an-image-by-digest-immutable-identifier
	//
	ImagePullPolicy_ALWAYS ImagePullPolicy = "ALWAYS"
	// The image is pulled only if it is not already present locally.
	ImagePullPolicy_IF_NOT_PRESENT ImagePullPolicy = "IF_NOT_PRESENT"
	// The kubelet does not try fetching the image.
	//
	// If the image is somehow already present locally,
	// the kubelet attempts to start the container; otherwise, startup fails.
	// See pre-pulled images for more details.
	// See: https://kubernetes.io/docs/concepts/containers/images/#pre-pulled-images
	//
	ImagePullPolicy_NEVER ImagePullPolicy = "NEVER"
)


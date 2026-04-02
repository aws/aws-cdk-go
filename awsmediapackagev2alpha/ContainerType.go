package awsmediapackagev2alpha


// The type of container to attach to this origin endpoint.
//
// A container type is a file format that encapsulates one or more media streams, such as audio and video, into a single file. You can't change the container type after you create the endpoint.
// Experimental.
type ContainerType string

const (
	// TS Container Type.
	// Experimental.
	ContainerType_TS ContainerType = "TS"
	// CMAF Container Type.
	// Experimental.
	ContainerType_CMAF ContainerType = "CMAF"
	// ISM Container Type (Microsoft Smooth Streaming).
	// Experimental.
	ContainerType_ISM ContainerType = "ISM"
)


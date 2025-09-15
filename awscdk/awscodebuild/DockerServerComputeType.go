package awscodebuild


// Docker server compute type.
// See: https://docs.aws.amazon.com/codebuild/latest/APIReference/API_DockerServer.html
//
type DockerServerComputeType string

const (
	// BUILD_GENERAL1_SMALL.
	DockerServerComputeType_SMALL DockerServerComputeType = "SMALL"
	// BUILD_GENERAL1_MEDIUM.
	DockerServerComputeType_MEDIUM DockerServerComputeType = "MEDIUM"
	// BUILD_GENERAL1_LARGE.
	DockerServerComputeType_LARGE DockerServerComputeType = "LARGE"
	// BUILD_GENERAL1_XLARGE.
	DockerServerComputeType_X_LARGE DockerServerComputeType = "X_LARGE"
	// BUILD_GENERAL1_2XLARGE.
	DockerServerComputeType_X2_LARGE DockerServerComputeType = "X2_LARGE"
)


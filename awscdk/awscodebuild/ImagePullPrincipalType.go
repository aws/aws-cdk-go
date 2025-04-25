package awscodebuild


// The type of principal CodeBuild will use to pull your build Docker image.
type ImagePullPrincipalType string

const (
	// CODEBUILD specifies that CodeBuild uses its own identity when pulling the image.
	//
	// This means the resource policy of the ECR repository that hosts the image will be modified to trust
	// CodeBuild's service principal.
	// This is the required principal type when using CodeBuild's pre-defined images.
	ImagePullPrincipalType_CODEBUILD ImagePullPrincipalType = "CODEBUILD"
	// SERVICE_ROLE specifies that AWS CodeBuild uses the project's role when pulling the image.
	//
	// The role will be granted pull permissions on the ECR repository hosting the image.
	ImagePullPrincipalType_SERVICE_ROLE ImagePullPrincipalType = "SERVICE_ROLE"
)


package cloudassemblyschema


// Type of cloud artifact.
// Experimental.
type ArtifactType string

const (
	// Stub required because of JSII.
	// Experimental.
	ArtifactType_NONE ArtifactType = "NONE"
	// The artifact is an AWS CloudFormation stack.
	// Experimental.
	ArtifactType_AWS_CLOUDFORMATION_STACK ArtifactType = "AWS_CLOUDFORMATION_STACK"
	// The artifact contains the CDK application's construct tree.
	// Experimental.
	ArtifactType_CDK_TREE ArtifactType = "CDK_TREE"
	// Manifest for all assets in the Cloud Assembly.
	// Experimental.
	ArtifactType_ASSET_MANIFEST ArtifactType = "ASSET_MANIFEST"
	// Nested Cloud Assembly.
	// Experimental.
	ArtifactType_NESTED_CLOUD_ASSEMBLY ArtifactType = "NESTED_CLOUD_ASSEMBLY"
)


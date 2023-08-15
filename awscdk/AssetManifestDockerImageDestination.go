package awscdk


// The destination for a docker image asset, when it is given to the AssetManifestBuilder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   assetManifestDockerImageDestination := &AssetManifestDockerImageDestination{
//   	RepositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	DockerTagPrefix: jsii.String("dockerTagPrefix"),
//   	Role: &RoleOptions{
//   		AssumeRoleArn: jsii.String("assumeRoleArn"),
//
//   		// the properties below are optional
//   		AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   	},
//   }
//
type AssetManifestDockerImageDestination struct {
	// Repository name where the docker image asset should be written.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Prefix to add to the asset hash to make the Docker image tag.
	// Default: ''.
	//
	DockerTagPrefix *string `field:"optional" json:"dockerTagPrefix" yaml:"dockerTagPrefix"`
	// Role to use to perform the upload.
	// Default: - No role.
	//
	Role *RoleOptions `field:"optional" json:"role" yaml:"role"`
}


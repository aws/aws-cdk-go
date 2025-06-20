package appstagingsynthesizeralpha


// Roles that are included in the Staging Stack (for access to Staging Resources).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//
//   var bootstrapRole bootstrapRole
//
//   stagingRoles := &StagingRoles{
//   	DockerAssetPublishingRole: bootstrapRole,
//   	FileAssetPublishingRole: bootstrapRole,
//   }
//
// Experimental.
type StagingRoles struct {
	// Docker Asset Publishing Role.
	// Default: - staging stack creates a docker asset publishing role.
	//
	// Experimental.
	DockerAssetPublishingRole BootstrapRole `field:"optional" json:"dockerAssetPublishingRole" yaml:"dockerAssetPublishingRole"`
	// File Asset Publishing Role.
	// Default: - staging stack creates a file asset publishing role.
	//
	// Experimental.
	FileAssetPublishingRole BootstrapRole `field:"optional" json:"fileAssetPublishingRole" yaml:"fileAssetPublishingRole"`
}


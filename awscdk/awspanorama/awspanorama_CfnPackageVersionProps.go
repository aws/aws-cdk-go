package awspanorama


// Properties for defining a `CfnPackageVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackageVersionProps := &cfnPackageVersionProps{
//   	packageId: jsii.String("packageId"),
//   	packageVersion: jsii.String("packageVersion"),
//   	patchVersion: jsii.String("patchVersion"),
//
//   	// the properties below are optional
//   	markLatest: jsii.Boolean(false),
//   	ownerAccount: jsii.String("ownerAccount"),
//   	updatedLatestPatchVersion: jsii.String("updatedLatestPatchVersion"),
//   }
//
type CfnPackageVersionProps struct {
	// A package ID.
	PackageId *string `field:"required" json:"packageId" yaml:"packageId"`
	// A package version.
	PackageVersion *string `field:"required" json:"packageVersion" yaml:"packageVersion"`
	// A patch version.
	PatchVersion *string `field:"required" json:"patchVersion" yaml:"patchVersion"`
	// Whether to mark the new version as the latest version.
	MarkLatest interface{} `field:"optional" json:"markLatest" yaml:"markLatest"`
	// An owner account.
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// If the version was marked latest, the new version to maker as latest.
	UpdatedLatestPatchVersion *string `field:"optional" json:"updatedLatestPatchVersion" yaml:"updatedLatestPatchVersion"`
}


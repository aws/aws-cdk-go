package awspanorama


// Properties for defining a `CfnPackageVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackageVersionProps := &CfnPackageVersionProps{
//   	PackageId: jsii.String("packageId"),
//   	PackageVersion: jsii.String("packageVersion"),
//   	PatchVersion: jsii.String("patchVersion"),
//
//   	// the properties below are optional
//   	MarkLatest: jsii.Boolean(false),
//   	OwnerAccount: jsii.String("ownerAccount"),
//   	UpdatedLatestPatchVersion: jsii.String("updatedLatestPatchVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html
//
type CfnPackageVersionProps struct {
	// A package ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-packageid
	//
	PackageId *string `field:"required" json:"packageId" yaml:"packageId"`
	// A package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-packageversion
	//
	PackageVersion *string `field:"required" json:"packageVersion" yaml:"packageVersion"`
	// A patch version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-patchversion
	//
	PatchVersion *string `field:"required" json:"patchVersion" yaml:"patchVersion"`
	// Whether to mark the new version as the latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-marklatest
	//
	MarkLatest interface{} `field:"optional" json:"markLatest" yaml:"markLatest"`
	// An owner account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-owneraccount
	//
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// If the version was marked latest, the new version to maker as latest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-updatedlatestpatchversion
	//
	UpdatedLatestPatchVersion *string `field:"optional" json:"updatedLatestPatchVersion" yaml:"updatedLatestPatchVersion"`
}


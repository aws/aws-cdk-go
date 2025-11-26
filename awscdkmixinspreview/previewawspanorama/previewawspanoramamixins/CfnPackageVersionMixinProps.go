package previewawspanoramamixins


// Properties for CfnPackageVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPackageVersionMixinProps := &CfnPackageVersionMixinProps{
//   	MarkLatest: jsii.Boolean(false),
//   	OwnerAccount: jsii.String("ownerAccount"),
//   	PackageId: jsii.String("packageId"),
//   	PackageVersion: jsii.String("packageVersion"),
//   	PatchVersion: jsii.String("patchVersion"),
//   	UpdatedLatestPatchVersion: jsii.String("updatedLatestPatchVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html
//
type CfnPackageVersionMixinProps struct {
	// Whether to mark the new version as the latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-marklatest
	//
	MarkLatest interface{} `field:"optional" json:"markLatest" yaml:"markLatest"`
	// An owner account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-owneraccount
	//
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// A package ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-packageid
	//
	PackageId *string `field:"optional" json:"packageId" yaml:"packageId"`
	// A package version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-packageversion
	//
	PackageVersion *string `field:"optional" json:"packageVersion" yaml:"packageVersion"`
	// A patch version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-patchversion
	//
	PatchVersion *string `field:"optional" json:"patchVersion" yaml:"patchVersion"`
	// If the version was marked latest, the new version to maker as latest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-packageversion.html#cfn-panorama-packageversion-updatedlatestpatchversion
	//
	UpdatedLatestPatchVersion *string `field:"optional" json:"updatedLatestPatchVersion" yaml:"updatedLatestPatchVersion"`
}


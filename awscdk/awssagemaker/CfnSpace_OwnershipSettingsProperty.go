package awssagemaker


// The collection of ownership settings for a space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ownershipSettingsProperty := &OwnershipSettingsProperty{
//   	OwnerUserProfileName: jsii.String("ownerUserProfileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-ownershipsettings.html
//
type CfnSpace_OwnershipSettingsProperty struct {
	// The user profile who is the owner of the private space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-ownershipsettings.html#cfn-sagemaker-space-ownershipsettings-owneruserprofilename
	//
	OwnerUserProfileName *string `field:"required" json:"ownerUserProfileName" yaml:"ownerUserProfileName"`
}


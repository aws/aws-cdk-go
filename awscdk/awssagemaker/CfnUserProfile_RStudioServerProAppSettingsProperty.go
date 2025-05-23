package awssagemaker


// A collection of settings that configure user interaction with the `RStudioServerPro` app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rStudioServerProAppSettingsProperty := &RStudioServerProAppSettingsProperty{
//   	AccessStatus: jsii.String("accessStatus"),
//   	UserGroup: jsii.String("userGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-rstudioserverproappsettings.html
//
type CfnUserProfile_RStudioServerProAppSettingsProperty struct {
	// Indicates whether the current user has access to the `RStudioServerPro` app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-rstudioserverproappsettings.html#cfn-sagemaker-userprofile-rstudioserverproappsettings-accessstatus
	//
	AccessStatus *string `field:"optional" json:"accessStatus" yaml:"accessStatus"`
	// The level of permissions that the user has within the `RStudioServerPro` app.
	//
	// This value defaults to `User`. The `Admin` value allows the user access to the RStudio Administrative Dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-rstudioserverproappsettings.html#cfn-sagemaker-userprofile-rstudioserverproappsettings-usergroup
	//
	UserGroup *string `field:"optional" json:"userGroup" yaml:"userGroup"`
}


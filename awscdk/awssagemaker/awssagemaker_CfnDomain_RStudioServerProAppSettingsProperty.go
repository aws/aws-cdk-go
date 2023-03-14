package awssagemaker


// A collection of settings that configure user interaction with the `RStudioServerPro` app.
//
// `RStudioServerProAppSettings` cannot be updated. The `RStudioServerPro` app must be deleted and a new one created to make any changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rStudioServerProAppSettingsProperty := &rStudioServerProAppSettingsProperty{
//   	accessStatus: jsii.String("accessStatus"),
//   	userGroup: jsii.String("userGroup"),
//   }
//
type CfnDomain_RStudioServerProAppSettingsProperty struct {
	// Indicates whether the current user has access to the `RStudioServerPro` app.
	AccessStatus *string `field:"optional" json:"accessStatus" yaml:"accessStatus"`
	// The level of permissions that the user has within the `RStudioServerPro` app.
	//
	// This value defaults to `User`. The `Admin` value allows the user access to the RStudio Administrative Dashboard.
	UserGroup *string `field:"optional" json:"userGroup" yaml:"userGroup"`
}


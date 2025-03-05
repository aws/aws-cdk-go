package awssagemaker


// A collection of space sharing settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceSharingSettingsProperty := &SpaceSharingSettingsProperty{
//   	SharingType: jsii.String("sharingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesharingsettings.html
//
type CfnSpace_SpaceSharingSettingsProperty struct {
	// Specifies the sharing type of the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesharingsettings.html#cfn-sagemaker-space-spacesharingsettings-sharingtype
	//
	SharingType *string `field:"required" json:"sharingType" yaml:"sharingType"`
}


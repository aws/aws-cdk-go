package awsimagebuilder


// The deletion settings of the image, indicating whether to delete the underlying resources in addition to the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deletionSettingsProperty := &DeletionSettingsProperty{
//   	ExecutionRole: jsii.String("executionRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-deletionsettings.html
//
type CfnImage_DeletionSettingsProperty struct {
	// The execution role to use for deleting the image, as well as underlying resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-deletionsettings.html#cfn-imagebuilder-image-deletionsettings-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
}


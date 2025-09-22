package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageCreatorProperty := &ModelPackageCreatorProperty{
//   	UserProfileName: jsii.String("userProfileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagecreator.html
//
type CfnModelCard_ModelPackageCreatorProperty struct {
	// The name of the user's profile in Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modelpackagecreator.html#cfn-sagemaker-modelcard-modelpackagecreator-userprofilename
	//
	UserProfileName *string `field:"optional" json:"userProfileName" yaml:"userProfileName"`
}


package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hiddenSageMakerImageProperty := &HiddenSageMakerImageProperty{
//   	SageMakerImageName: jsii.String("sageMakerImageName"),
//   	VersionAliases: []*string{
//   		jsii.String("versionAliases"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-hiddensagemakerimage.html
//
type CfnUserProfile_HiddenSageMakerImageProperty struct {
	// The SageMaker image name that you are hiding from the Studio user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-hiddensagemakerimage.html#cfn-sagemaker-userprofile-hiddensagemakerimage-sagemakerimagename
	//
	SageMakerImageName *string `field:"optional" json:"sageMakerImageName" yaml:"sageMakerImageName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-hiddensagemakerimage.html#cfn-sagemaker-userprofile-hiddensagemakerimage-versionaliases
	//
	VersionAliases *[]*string `field:"optional" json:"versionAliases" yaml:"versionAliases"`
}


package awssagemaker


// The SageMaker images that are hidden from the Studio user interface.
//
// You must specify the SageMaker image name and version aliases.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-hiddensagemakerimage.html
//
type CfnDomain_HiddenSageMakerImageProperty struct {
	// The SageMaker image name that you are hiding from the Studio user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-hiddensagemakerimage.html#cfn-sagemaker-domain-hiddensagemakerimage-sagemakerimagename
	//
	SageMakerImageName *string `field:"optional" json:"sageMakerImageName" yaml:"sageMakerImageName"`
	// The version aliases you are hiding from the Studio user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-hiddensagemakerimage.html#cfn-sagemaker-domain-hiddensagemakerimage-versionaliases
	//
	VersionAliases *[]*string `field:"optional" json:"versionAliases" yaml:"versionAliases"`
}


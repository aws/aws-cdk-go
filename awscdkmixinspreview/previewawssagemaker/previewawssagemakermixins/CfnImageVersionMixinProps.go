package previewawssagemakermixins


// Properties for CfnImageVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnImageVersionMixinProps := &CfnImageVersionMixinProps{
//   	Alias: jsii.String("alias"),
//   	Aliases: []*string{
//   		jsii.String("aliases"),
//   	},
//   	BaseImage: jsii.String("baseImage"),
//   	Horovod: jsii.Boolean(false),
//   	ImageName: jsii.String("imageName"),
//   	JobType: jsii.String("jobType"),
//   	MlFramework: jsii.String("mlFramework"),
//   	Processor: jsii.String("processor"),
//   	ProgrammingLang: jsii.String("programmingLang"),
//   	ReleaseNotes: jsii.String("releaseNotes"),
//   	VendorGuidance: jsii.String("vendorGuidance"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html
//
type CfnImageVersionMixinProps struct {
	// The alias of the image version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// List of aliases for the image version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-aliases
	//
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// The container image that the SageMaker image version is based on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-baseimage
	//
	BaseImage *string `field:"optional" json:"baseImage" yaml:"baseImage"`
	// Indicates Horovod compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-horovod
	//
	Horovod interface{} `field:"optional" json:"horovod" yaml:"horovod"`
	// The name of the parent image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-imagename
	//
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Indicates SageMaker job type compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-jobtype
	//
	JobType *string `field:"optional" json:"jobType" yaml:"jobType"`
	// The machine learning framework vended in the image version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-mlframework
	//
	MlFramework *string `field:"optional" json:"mlFramework" yaml:"mlFramework"`
	// Indicates CPU or GPU compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-processor
	//
	Processor *string `field:"optional" json:"processor" yaml:"processor"`
	// The supported programming language and its version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-programminglang
	//
	ProgrammingLang *string `field:"optional" json:"programmingLang" yaml:"programmingLang"`
	// The maintainer description of the image version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-releasenotes
	//
	ReleaseNotes *string `field:"optional" json:"releaseNotes" yaml:"releaseNotes"`
	// The availability of the image version specified by the maintainer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-vendorguidance
	//
	VendorGuidance *string `field:"optional" json:"vendorGuidance" yaml:"vendorGuidance"`
}


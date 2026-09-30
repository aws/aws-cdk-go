package awselasticbeanstalk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   imageBuildConfigurationProperty := &ImageBuildConfigurationProperty{
//   	Architecture: jsii.String("architecture"),
//   	Buildpack: jsii.String("buildpack"),
//   	CodeBuildServiceRole: jsii.String("codeBuildServiceRole"),
//   	ComputeType: jsii.String("computeType"),
//   	DockerfileLocation: jsii.String("dockerfileLocation"),
//   	TimeoutInMinutes: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html
//
type CfnApplicationVersionPropsMixin_ImageBuildConfigurationProperty struct {
	// The target architecture for the built container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html#cfn-elasticbeanstalk-applicationversion-imagebuildconfiguration-architecture
	//
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// The buildpack to use for building the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html#cfn-elasticbeanstalk-applicationversion-imagebuildconfiguration-buildpack
	//
	Buildpack *string `field:"optional" json:"buildpack" yaml:"buildpack"`
	// The ARN of the IAM role that AWS CodeBuild assumes to build the application version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html#cfn-elasticbeanstalk-applicationversion-imagebuildconfiguration-codebuildservicerole
	//
	CodeBuildServiceRole *string `field:"optional" json:"codeBuildServiceRole" yaml:"codeBuildServiceRole"`
	// The compute type for the CodeBuild build environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html#cfn-elasticbeanstalk-applicationversion-imagebuildconfiguration-computetype
	//
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// The path to the Dockerfile, relative to the source root.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html#cfn-elasticbeanstalk-applicationversion-imagebuildconfiguration-dockerfilelocation
	//
	DockerfileLocation *string `field:"optional" json:"dockerfileLocation" yaml:"dockerfileLocation"`
	// The timeout for the CodeBuild build, in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html#cfn-elasticbeanstalk-applicationversion-imagebuildconfiguration-timeoutinminutes
	//
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// The type of image build: docker or buildpack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagebuildconfiguration.html#cfn-elasticbeanstalk-applicationversion-imagebuildconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


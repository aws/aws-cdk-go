package awselasticbeanstalk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buildConfigurationProperty := &BuildConfigurationProperty{
//   	CodeBuildServiceRole: jsii.String("codeBuildServiceRole"),
//   	Image: jsii.String("image"),
//
//   	// the properties below are optional
//   	ArtifactName: jsii.String("artifactName"),
//   	ComputeType: jsii.String("computeType"),
//   	TimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-buildconfiguration.html
//
type CfnApplicationVersion_BuildConfigurationProperty struct {
	// The ARN of the IAM role that AWS CodeBuild assumes to build the application version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-buildconfiguration.html#cfn-elasticbeanstalk-applicationversion-buildconfiguration-codebuildservicerole
	//
	CodeBuildServiceRole *string `field:"required" json:"codeBuildServiceRole" yaml:"codeBuildServiceRole"`
	// The CodeBuild image used for the build environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-buildconfiguration.html#cfn-elasticbeanstalk-applicationversion-buildconfiguration-image
	//
	Image *string `field:"required" json:"image" yaml:"image"`
	// The name of the build artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-buildconfiguration.html#cfn-elasticbeanstalk-applicationversion-buildconfiguration-artifactname
	//
	ArtifactName *string `field:"optional" json:"artifactName" yaml:"artifactName"`
	// The compute type for the CodeBuild build environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-buildconfiguration.html#cfn-elasticbeanstalk-applicationversion-buildconfiguration-computetype
	//
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// The timeout for the CodeBuild build, in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-buildconfiguration.html#cfn-elasticbeanstalk-applicationversion-buildconfiguration-timeoutinminutes
	//
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}


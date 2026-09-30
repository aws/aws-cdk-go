package awselasticbeanstalk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigurationProperty := &ImageConfigurationProperty{
//   	BuildProperty: &ImageBuildConfigurationProperty{
//   		CodeBuildServiceRole: jsii.String("codeBuildServiceRole"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Architecture: jsii.String("architecture"),
//   		Buildpack: jsii.String("buildpack"),
//   		ComputeType: jsii.String("computeType"),
//   		DockerfileLocation: jsii.String("dockerfileLocation"),
//   		TimeoutInMinutes: jsii.Number(123),
//   	},
//   	Source: &ImageSourceProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imageconfiguration.html
//
type CfnApplicationVersion_ImageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imageconfiguration.html#cfn-elasticbeanstalk-applicationversion-imageconfiguration-build
	//
	BuildProperty interface{} `field:"optional" json:"buildProperty" yaml:"buildProperty"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imageconfiguration.html#cfn-elasticbeanstalk-applicationversion-imageconfiguration-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}


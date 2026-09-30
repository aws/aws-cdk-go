package awselasticbeanstalk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   imageConfigurationProperty := &ImageConfigurationProperty{
//   	BuildProperty: &ImageBuildConfigurationProperty{
//   		Architecture: jsii.String("architecture"),
//   		Buildpack: jsii.String("buildpack"),
//   		CodeBuildServiceRole: jsii.String("codeBuildServiceRole"),
//   		ComputeType: jsii.String("computeType"),
//   		DockerfileLocation: jsii.String("dockerfileLocation"),
//   		TimeoutInMinutes: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	Source: &ImageSourceProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imageconfiguration.html
//
type CfnApplicationVersionPropsMixin_ImageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imageconfiguration.html#cfn-elasticbeanstalk-applicationversion-imageconfiguration-build
	//
	BuildProperty interface{} `field:"optional" json:"buildProperty" yaml:"buildProperty"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imageconfiguration.html#cfn-elasticbeanstalk-applicationversion-imageconfiguration-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}


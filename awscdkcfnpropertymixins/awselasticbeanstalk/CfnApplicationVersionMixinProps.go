package awselasticbeanstalk


// Properties for CfnApplicationVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnApplicationVersionMixinProps := &CfnApplicationVersionMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	BuildConfiguration: &BuildConfigurationProperty{
//   		ArtifactName: jsii.String("artifactName"),
//   		CodeBuildServiceRole: jsii.String("codeBuildServiceRole"),
//   		ComputeType: jsii.String("computeType"),
//   		Image: jsii.String("image"),
//   		TimeoutInMinutes: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	ImageConfiguration: &ImageConfigurationProperty{
//   		BuildProperty: &ImageBuildConfigurationProperty{
//   			Architecture: jsii.String("architecture"),
//   			Buildpack: jsii.String("buildpack"),
//   			CodeBuildServiceRole: jsii.String("codeBuildServiceRole"),
//   			ComputeType: jsii.String("computeType"),
//   			DockerfileLocation: jsii.String("dockerfileLocation"),
//   			TimeoutInMinutes: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   		Source: &ImageSourceProperty{
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	Process: jsii.Boolean(false),
//   	SourceBundle: &SourceBundleProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html
//
type CfnApplicationVersionMixinProps struct {
	// The name of the Elastic Beanstalk application that is associated with this application version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-applicationname
	//
	ApplicationName interface{} `field:"optional" json:"applicationName" yaml:"applicationName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-buildconfiguration
	//
	BuildConfiguration interface{} `field:"optional" json:"buildConfiguration" yaml:"buildConfiguration"`
	// A description of this application version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-imageconfiguration
	//
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
	// Pre-process and validate the environment manifest (`env.yaml`) and configuration files in the source bundle. Leave unset for the service default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-process
	//
	Process interface{} `field:"optional" json:"process" yaml:"process"`
	// The Amazon S3 bucket and key that identify the location of the source bundle for this version.
	//
	// > The Amazon S3 bucket must be in the same region as the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	//
	SourceBundle interface{} `field:"optional" json:"sourceBundle" yaml:"sourceBundle"`
}


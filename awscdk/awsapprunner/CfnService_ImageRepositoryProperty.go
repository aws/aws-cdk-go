package awsapprunner


// Describes a source image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageRepositoryProperty := &ImageRepositoryProperty{
//   	ImageIdentifier: jsii.String("imageIdentifier"),
//   	ImageRepositoryType: jsii.String("imageRepositoryType"),
//
//   	// the properties below are optional
//   	ImageConfiguration: &ImageConfigurationProperty{
//   		Port: jsii.String("port"),
//   		RuntimeEnvironmentSecrets: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RuntimeEnvironmentVariables: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		StartCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html
//
type CfnService_ImageRepositoryProperty struct {
	// The identifier of an image.
	//
	// For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the image name format, see [Pulling an image](https://docs.aws.amazon.com/AmazonECR/latest/userguide/docker-pull-ecr-image.html) in the *Amazon ECR User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html#cfn-apprunner-service-imagerepository-imageidentifier
	//
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether the repository is private or public.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html#cfn-apprunner-service-imagerepository-imagerepositorytype
	//
	ImageRepositoryType *string `field:"required" json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html#cfn-apprunner-service-imagerepository-imageconfiguration
	//
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}


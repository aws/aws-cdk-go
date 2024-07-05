package awscdkapprunneralpha


// Describes a source image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var secret secret
//
//   imageRepository := &ImageRepository{
//   	ImageIdentifier: jsii.String("imageIdentifier"),
//   	ImageRepositoryType: apprunner_alpha.ImageRepositoryType_ECR_PUBLIC,
//
//   	// the properties below are optional
//   	ImageConfiguration: &ImageConfiguration{
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		EnvironmentSecrets: map[string]*secret{
//   			"environmentSecretsKey": secret,
//   		},
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		Port: jsii.Number(123),
//   		StartCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html
//
// Experimental.
type ImageRepository struct {
	// The identifier of the image.
	//
	// For `ECR_PUBLIC` imageRepositoryType, the identifier domain should
	// always be `public.ecr.aws`. For `ECR`, the pattern should be
	// `([0-9]{12}.dkr.ecr.[a-z\-]+-[0-9]{1}.amazonaws.com\/.*)`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html
	//
	// Experimental.
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// The type of the image repository.
	//
	// This reflects the repository provider and whether
	// the repository is private or public.
	// Experimental.
	ImageRepositoryType ImageRepositoryType `field:"required" json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Configuration for running the identified image.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Default: - no image configuration will be passed. The default `port` will be 8080.
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}


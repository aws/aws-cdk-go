package awsapprunner


// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// Example:
//   import assets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   imageAsset := assets.NewDockerImageAsset(this, jsii.String("ImageAssets"), &DockerImageAssetProps{
//   	Directory: path.join(__dirname, jsii.String("./docker.assets")),
//   })
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromAsset(&AssetProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		Asset: imageAsset,
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html
//
// Experimental.
type ImageConfiguration struct {
	// Environment variables that are available to your running App Runner service.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	// Experimental.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}


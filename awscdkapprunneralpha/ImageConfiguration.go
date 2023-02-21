// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	IpAddresses: ec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &VpcConnectorProps{
//   	Vpc: Vpc,
//   	VpcSubnets: vpc.selectSubnets(&SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	}),
//   	VpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	VpcConnector: VpcConnector,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html
//
// Experimental.
type ImageConfiguration struct {
	// Environment variables that are available to your running App Runner service.
	// Deprecated: use environmentVariables.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Environment secrets that are available to your running App Runner service.
	// Experimental.
	EnvironmentSecrets *map[string]Secret `field:"optional" json:"environmentSecrets" yaml:"environmentSecrets"`
	// Environment variables that are available to your running App Runner service.
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The port that your application listens to in the container.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	// Experimental.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}


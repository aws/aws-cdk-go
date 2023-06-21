package awscdkapprunneralpha


// Properties of the image repository for `Source.fromEcrPublic()`.
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
// Experimental.
type EcrPublicProps struct {
	// The ECR Public image URI.
	// Experimental.
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// The image configuration for the image from ECR Public.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}


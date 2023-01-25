// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// Properties of the image repository for `Source.fromEcrPublic()`.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
//   	ipAddresses: ec2.ipAddresses.cidr(jsii.String("10.0.0.0/16")),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &vpcConnectorProps{
//   	vpc: vpc,
//   	vpcSubnets: vpc.selectSubnets(&subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	}),
//   	vpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	vpcConnector: vpcConnector,
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


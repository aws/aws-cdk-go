package awscdkapprunneralpha


// Properties of the image repository for `Source.fromEcrPublic()`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   service := apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
//   service.AddToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Effect: iam.Effect_ALLOW,
//   	Actions: []*string{
//   		jsii.String("s3:GetObject"),
//   	},
//   	Resources: []*string{
//   		jsii.String("*"),
//   	},
//   }))
//
// Experimental.
type EcrPublicProps struct {
	// The ECR Public image URI.
	// Experimental.
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// The image configuration for the image from ECR Public.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imageconfiguration.html#cfn-apprunner-service-imageconfiguration-port
	//
	// Default: - no image configuration will be passed. The default `port` will be 8080.
	//
	// Experimental.
	ImageConfiguration *ImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}


package awsevidently


// This is a structure that defines the configuration of how your application integrates with AWS AppConfig to run client-side evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appConfigResourceObjectProperty := &appConfigResourceObjectProperty{
//   	applicationId: jsii.String("applicationId"),
//   	environmentId: jsii.String("environmentId"),
//   }
//
type CfnProject_AppConfigResourceObjectProperty struct {
	// The ID of the AWS AppConfig application to use for client-side evaluation.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ID of the AWS AppConfig environment to use for client-side evaluation.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}


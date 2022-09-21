package awsevidently


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
	// `CfnProject.AppConfigResourceObjectProperty.ApplicationId`.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// `CfnProject.AppConfigResourceObjectProperty.EnvironmentId`.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}


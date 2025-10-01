package awselasticbeanstalk


// A reference to a ApplicationVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationVersionReference := &ApplicationVersionReference{
//   	ApplicationName: jsii.String("applicationName"),
//   	ApplicationVersionId: jsii.String("applicationVersionId"),
//   }
//
type ApplicationVersionReference struct {
	// The ApplicationName of the ApplicationVersion resource.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The Id of the ApplicationVersion resource.
	ApplicationVersionId *string `field:"required" json:"applicationVersionId" yaml:"applicationVersionId"`
}


package awsqbusiness


// A reference to a WebExperience resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webExperienceReference := &WebExperienceReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	WebExperienceArn: jsii.String("webExperienceArn"),
//   	WebExperienceId: jsii.String("webExperienceId"),
//   }
//
type WebExperienceReference struct {
	// The ApplicationId of the WebExperience resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ARN of the WebExperience resource.
	WebExperienceArn *string `field:"required" json:"webExperienceArn" yaml:"webExperienceArn"`
	// The WebExperienceId of the WebExperience resource.
	WebExperienceId *string `field:"required" json:"webExperienceId" yaml:"webExperienceId"`
}


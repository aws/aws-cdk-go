package awsproton


// A reference to a EnvironmentTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentTemplateReference := &EnvironmentTemplateReference{
//   	EnvironmentTemplateArn: jsii.String("environmentTemplateArn"),
//   }
//
type EnvironmentTemplateReference struct {
	// The Arn of the EnvironmentTemplate resource.
	EnvironmentTemplateArn *string `field:"required" json:"environmentTemplateArn" yaml:"environmentTemplateArn"`
}


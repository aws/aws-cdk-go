package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationLocationProperty := &applicationLocationProperty{
//   	applicationId: jsii.String("applicationId"),
//   	semanticVersion: jsii.String("semanticVersion"),
//   }
//
type CfnApplication_ApplicationLocationProperty struct {
	// `CfnApplication.ApplicationLocationProperty.ApplicationId`.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// `CfnApplication.ApplicationLocationProperty.SemanticVersion`.
	SemanticVersion *string `field:"required" json:"semanticVersion" yaml:"semanticVersion"`
}


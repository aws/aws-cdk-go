package awsinspectorv2


// A reference to a CodeSecurityIntegration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeSecurityIntegrationReference := &CodeSecurityIntegrationReference{
//   	CodeSecurityIntegrationArn: jsii.String("codeSecurityIntegrationArn"),
//   }
//
type CodeSecurityIntegrationReference struct {
	// The Arn of the CodeSecurityIntegration resource.
	CodeSecurityIntegrationArn *string `field:"required" json:"codeSecurityIntegrationArn" yaml:"codeSecurityIntegrationArn"`
}


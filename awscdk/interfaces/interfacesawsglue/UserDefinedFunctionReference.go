package interfacesawsglue


// A reference to a UserDefinedFunction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userDefinedFunctionReference := &UserDefinedFunctionReference{
//   	UserDefinedFunctionArn: jsii.String("userDefinedFunctionArn"),
//   }
//
type UserDefinedFunctionReference struct {
	// The Arn of the UserDefinedFunction resource.
	UserDefinedFunctionArn *string `field:"required" json:"userDefinedFunctionArn" yaml:"userDefinedFunctionArn"`
}


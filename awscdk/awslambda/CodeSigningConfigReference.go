package awslambda


// A reference to a CodeSigningConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeSigningConfigReference := &CodeSigningConfigReference{
//   	CodeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   }
//
type CodeSigningConfigReference struct {
	// The CodeSigningConfigArn of the CodeSigningConfig resource.
	CodeSigningConfigArn *string `field:"required" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
}


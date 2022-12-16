// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Customization properties for an Intrinsic token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   intrinsicProps := &intrinsicProps{
//   	stackTrace: jsii.Boolean(false),
//   	typeHint: cdk.resolutionTypeHint_STRING,
//   }
//
type IntrinsicProps struct {
	// Capture the stack trace of where this token is created.
	StackTrace *bool `field:"optional" json:"stackTrace" yaml:"stackTrace"`
	// Type that this token is expected to evaluate to.
	TypeHint ResolutionTypeHint `field:"optional" json:"typeHint" yaml:"typeHint"`
}


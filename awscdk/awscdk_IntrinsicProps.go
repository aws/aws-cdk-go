// An experiment to bundle the entire CDK into a single module
package awscdk


// Customization properties for an Intrinsic token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   intrinsicProps := &IntrinsicProps{
//   	StackTrace: jsii.Boolean(false),
//   }
//
// Experimental.
type IntrinsicProps struct {
	// Capture the stack trace of where this token is created.
	// Experimental.
	StackTrace *bool `field:"optional" json:"stackTrace" yaml:"stackTrace"`
}


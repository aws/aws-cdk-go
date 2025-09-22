package awschatbot


// A reference to a CustomAction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionReference := &CustomActionReference{
//   	CustomActionArn: jsii.String("customActionArn"),
//   }
//
type CustomActionReference struct {
	// The CustomActionArn of the CustomAction resource.
	CustomActionArn *string `field:"required" json:"customActionArn" yaml:"customActionArn"`
}


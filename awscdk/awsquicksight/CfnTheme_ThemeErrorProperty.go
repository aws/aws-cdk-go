package awsquicksight


// Theme error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   themeErrorProperty := &ThemeErrorProperty{
//   	Message: jsii.String("message"),
//   	Type: jsii.String("type"),
//   }
//
type CfnTheme_ThemeErrorProperty struct {
	// The error message.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The type of error.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


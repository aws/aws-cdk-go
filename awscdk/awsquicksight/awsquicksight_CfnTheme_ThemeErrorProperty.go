package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   themeErrorProperty := &themeErrorProperty{
//   	message: jsii.String("message"),
//   	type: jsii.String("type"),
//   }
//
type CfnTheme_ThemeErrorProperty struct {
	// `CfnTheme.ThemeErrorProperty.Message`.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// `CfnTheme.ThemeErrorProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


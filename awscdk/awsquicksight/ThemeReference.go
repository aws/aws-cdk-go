package awsquicksight


// A reference to a Theme resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   themeReference := &ThemeReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	ThemeArn: jsii.String("themeArn"),
//   	ThemeId: jsii.String("themeId"),
//   }
//
type ThemeReference struct {
	// The AwsAccountId of the Theme resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the Theme resource.
	ThemeArn *string `field:"required" json:"themeArn" yaml:"themeArn"`
	// The ThemeId of the Theme resource.
	ThemeId *string `field:"required" json:"themeId" yaml:"themeId"`
}


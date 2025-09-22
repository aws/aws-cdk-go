package awsamplifyuibuilder


// A reference to a Theme resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   themeReference := &ThemeReference{
//   	AppId: jsii.String("appId"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	ThemeId: jsii.String("themeId"),
//   }
//
type ThemeReference struct {
	// The AppId of the Theme resource.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The EnvironmentName of the Theme resource.
	EnvironmentName *string `field:"required" json:"environmentName" yaml:"environmentName"`
	// The Id of the Theme resource.
	ThemeId *string `field:"required" json:"themeId" yaml:"themeId"`
}


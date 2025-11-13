package interfacesawsamplifyuibuilder


// A reference to a Form resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formReference := &FormReference{
//   	AppId: jsii.String("appId"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	FormId: jsii.String("formId"),
//   }
//
type FormReference struct {
	// The AppId of the Form resource.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The EnvironmentName of the Form resource.
	EnvironmentName *string `field:"required" json:"environmentName" yaml:"environmentName"`
	// The Id of the Form resource.
	FormId *string `field:"required" json:"formId" yaml:"formId"`
}


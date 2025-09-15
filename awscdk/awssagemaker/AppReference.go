package awssagemaker


// A reference to a App resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appReference := &AppReference{
//   	AppArn: jsii.String("appArn"),
//   	AppName: jsii.String("appName"),
//   	AppType: jsii.String("appType"),
//   	DomainId: jsii.String("domainId"),
//   	UserProfileName: jsii.String("userProfileName"),
//   }
//
type AppReference struct {
	// The ARN of the App resource.
	AppArn *string `field:"required" json:"appArn" yaml:"appArn"`
	// The AppName of the App resource.
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The AppType of the App resource.
	AppType *string `field:"required" json:"appType" yaml:"appType"`
	// The DomainId of the App resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The UserProfileName of the App resource.
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
}


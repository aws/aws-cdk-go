package interfacesawschime


// A reference to a AppInstanceBot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appInstanceBotReference := &AppInstanceBotReference{
//   	AppInstanceBotArn: jsii.String("appInstanceBotArn"),
//   }
//
type AppInstanceBotReference struct {
	// The AppInstanceBotArn of the AppInstanceBot resource.
	AppInstanceBotArn *string `field:"required" json:"appInstanceBotArn" yaml:"appInstanceBotArn"`
}


package interfacesawsiot


// A reference to a Logging resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingReference := &LoggingReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type LoggingReference struct {
	// The AccountId of the Logging resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}


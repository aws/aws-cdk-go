package interfacesawsmacie


// A reference to a Session resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionReference := &SessionReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   }
//
type SessionReference struct {
	// The AwsAccountId of the Session resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
}


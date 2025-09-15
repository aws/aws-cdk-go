package awsapigateway


// A reference to a Account resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountReference := &AccountReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type AccountReference struct {
	// The Id of the Account resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}


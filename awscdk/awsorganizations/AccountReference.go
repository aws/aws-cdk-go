package awsorganizations


// A reference to a Account resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountReference := &AccountReference{
//   	AccountArn: jsii.String("accountArn"),
//   	AccountId: jsii.String("accountId"),
//   }
//
type AccountReference struct {
	// The ARN of the Account resource.
	AccountArn *string `field:"required" json:"accountArn" yaml:"accountArn"`
	// The AccountId of the Account resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}


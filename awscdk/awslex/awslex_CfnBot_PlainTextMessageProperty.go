package awslex


// Defines an ASCII text message to send to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   plainTextMessageProperty := &plainTextMessageProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_PlainTextMessageProperty struct {
	// The message to send to the user.
	Value *string `field:"required" json:"value" yaml:"value"`
}


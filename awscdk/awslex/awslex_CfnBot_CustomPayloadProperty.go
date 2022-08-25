package awslex


// A custom response string that Amazon Lex sends to your application.
//
// You define the content and structure of the string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPayloadProperty := &customPayloadProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_CustomPayloadProperty struct {
	// The string that is sent to your application.
	Value *string `field:"required" json:"value" yaml:"value"`
}


package awslex


// Defines a Speech Synthesis Markup Language (SSML) prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSMLMessageProperty := &sSMLMessageProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_SSMLMessageProperty struct {
	// The SSML text that defines the prompt.
	Value *string `field:"required" json:"value" yaml:"value"`
}


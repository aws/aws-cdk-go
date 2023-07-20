package awscdkiotactionsalpha


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_actions_alpha "github.com/aws/aws-cdk-go/awscdkiotactionsalpha"
//
//   httpActionHeader := &HttpActionHeader{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// Experimental.
type HttpActionHeader struct {
	// The HTTP header key.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The HTTP header value.
	//
	// Substitution templates are supported.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}


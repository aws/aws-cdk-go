package awsapigateway


// A reference to a RequestValidator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestValidatorReference := &RequestValidatorReference{
//   	RequestValidatorId: jsii.String("requestValidatorId"),
//   }
//
type RequestValidatorReference struct {
	// The RequestValidatorId of the RequestValidator resource.
	RequestValidatorId *string `field:"required" json:"requestValidatorId" yaml:"requestValidatorId"`
}


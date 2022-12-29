package awsses


// AddHeaderAction configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addHeaderActionConfig := &addHeaderActionConfig{
//   	headerName: jsii.String("headerName"),
//   	headerValue: jsii.String("headerValue"),
//   }
//
type AddHeaderActionConfig struct {
	// The name of the header that you want to add to the incoming message.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The content that you want to include in the header.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}


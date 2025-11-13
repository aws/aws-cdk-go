package interfacesawsapigatewayv2


// A reference to a Api resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiReference := &ApiReference{
//   	ApiId: jsii.String("apiId"),
//   }
//
type ApiReference struct {
	// The ApiId of the Api resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
}


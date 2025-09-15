package awsapigatewayv2


// The attributes used to import existing ApiMapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiMappingAttributes := &ApiMappingAttributes{
//   	ApiMappingId: jsii.String("apiMappingId"),
//   }
//
type ApiMappingAttributes struct {
	// The API mapping ID.
	ApiMappingId *string `field:"required" json:"apiMappingId" yaml:"apiMappingId"`
}


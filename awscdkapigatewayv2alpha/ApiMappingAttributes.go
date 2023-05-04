// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// The attributes used to import existing ApiMapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   apiMappingAttributes := &ApiMappingAttributes{
//   	ApiMappingId: jsii.String("apiMappingId"),
//   }
//
// Experimental.
type ApiMappingAttributes struct {
	// The API mapping ID.
	// Experimental.
	ApiMappingId *string `field:"required" json:"apiMappingId" yaml:"apiMappingId"`
}


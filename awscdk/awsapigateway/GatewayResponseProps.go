package awsapigateway


// Properties for a new gateway response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseType responseType
//   var restApi restApi
//
//   gatewayResponseProps := &GatewayResponseProps{
//   	RestApi: restApi,
//   	Type: responseType,
//
//   	// the properties below are optional
//   	ResponseHeaders: map[string]*string{
//   		"responseHeadersKey": jsii.String("responseHeaders"),
//   	},
//   	StatusCode: jsii.String("statusCode"),
//   	Templates: map[string]*string{
//   		"templatesKey": jsii.String("templates"),
//   	},
//   }
//
// Experimental.
type GatewayResponseProps struct {
	// Response type to associate with gateway response.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html
	//
	// Experimental.
	Type ResponseType `field:"required" json:"type" yaml:"type"`
	// Custom headers parameters for response.
	// Experimental.
	ResponseHeaders *map[string]*string `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
	// Http status code for response.
	// Experimental.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Custom templates to get mapped as response.
	// Experimental.
	Templates *map[string]*string `field:"optional" json:"templates" yaml:"templates"`
	// Rest api resource to target.
	// Experimental.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
}


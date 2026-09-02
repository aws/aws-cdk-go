package interfacesawsbedrockagentcore


// A reference to a GatewayRateLimit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRateLimitReference := &GatewayRateLimitReference{
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	RateLimitId: jsii.String("rateLimitId"),
//   }
//
type GatewayRateLimitReference struct {
	// The GatewayIdentifier of the GatewayRateLimit resource.
	GatewayIdentifier *string `field:"required" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// The RateLimitId of the GatewayRateLimit resource.
	RateLimitId *string `field:"required" json:"rateLimitId" yaml:"rateLimitId"`
}


package awselasticloadbalancingv2


// Options for `ListenerAction.authenticateJwt()`.
//
// Example:
//   var lb ApplicationLoadBalancer
//   var certificate IListenerCertificate
//   var myTargetGroup ApplicationTargetGroup
//
//
//   // JWT authentication requires HTTPS
//   listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Protocol: elbv2.ApplicationProtocol_HTTPS,
//   	Port: jsii.Number(443),
//   	Certificates: []IListenerCertificate{
//   		certificate,
//   	},
//   	DefaultAction: elbv2.ListenerAction_AuthenticateJwt(&AuthenticateJwtOptions{
//   		Issuer: jsii.String("https://issuer.example.com"),
//   		JwksEndpoint: jsii.String("https://issuer.example.com/.well-known/jwks.json"),
//   		Next: elbv2.ListenerAction_Forward([]IApplicationTargetGroup{
//   			myTargetGroup,
//   		}),
//   	}),
//   })
//
type AuthenticateJwtOptions struct {
	// The issuer of the JWT token.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	//
	// Example:
	//   "https://issuer.example.com"
	//
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The JWKS (JSON Web Key Set) endpoint URL.
	//
	// The endpoint must be publicly accessible and return the public keys used to verify JWT signatures.
	//
	// Example:
	//   "https://issuer.example.com/jwks"
	//
	JwksEndpoint *string `field:"required" json:"jwksEndpoint" yaml:"jwksEndpoint"`
	// What action to execute next.
	//
	// Multiple actions form a linked chain; the chain must always terminate in a
	// (weighted)forward, fixedResponse or redirect action.
	Next ListenerAction `field:"required" json:"next" yaml:"next"`
}


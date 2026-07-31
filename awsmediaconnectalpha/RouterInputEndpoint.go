package awsmediaconnectalpha


// A single ingest endpoint where the router input listens for upstream content.
//
// Failover and merge configurations expose two endpoints; standard configurations
// expose one. Returned by {@link IRouterInput.endpoints}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   routerInputEndpoint := &RouterInputEndpoint{
//   	Port: jsii.Number(123),
//   	Url: jsii.String("url"),
//   }
//
// Experimental.
type RouterInputEndpoint struct {
	// The listening port at which the router input accepts upstream content.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The full ingest URL combining protocol scheme, IP, and port.
	//
	// For example:
	// `srt://203.0.113.10:5000` or `rtp://203.0.113.10:5001`.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
}


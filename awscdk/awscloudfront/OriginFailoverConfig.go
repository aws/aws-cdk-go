package awscloudfront


// The failover configuration used for Origin Groups, returned in `OriginBindConfig.failoverConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var origin IOrigin
//
//   originFailoverConfig := &OriginFailoverConfig{
//   	FailoverOrigin: origin,
//
//   	// the properties below are optional
//   	StatusCodes: []*f64{
//   		jsii.Number(123),
//   	},
//   }
//
type OriginFailoverConfig struct {
	// The origin to use as the fallback origin.
	FailoverOrigin IOrigin `field:"required" json:"failoverOrigin" yaml:"failoverOrigin"`
	// The HTTP status codes of the response that trigger querying the failover Origin.
	// Default: - 500, 502, 503 and 504.
	//
	StatusCodes *[]*float64 `field:"optional" json:"statusCodes" yaml:"statusCodes"`
}


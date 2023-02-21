package awsappsync


// The caching configuration for a resolver that has caching activated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cachingConfigProperty := &cachingConfigProperty{
//   	ttl: jsii.Number(123),
//
//   	// the properties below are optional
//   	cachingKeys: []*string{
//   		jsii.String("cachingKeys"),
//   	},
//   }
//
type CfnResolver_CachingConfigProperty struct {
	// The TTL in seconds for a resolver that has caching activated.
	//
	// Valid values are 1–3,600 seconds.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// The caching keys for a resolver that has caching activated.
	//
	// Valid values are entries from the `$context.arguments` , `$context.source` , and `$context.identity` maps.
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
}


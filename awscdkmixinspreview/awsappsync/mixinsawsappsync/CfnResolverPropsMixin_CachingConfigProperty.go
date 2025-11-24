package mixinsawsappsync


// The caching configuration for a resolver that has caching activated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cachingConfigProperty := &CachingConfigProperty{
//   	CachingKeys: []*string{
//   		jsii.String("cachingKeys"),
//   	},
//   	Ttl: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-cachingconfig.html
//
type CfnResolverPropsMixin_CachingConfigProperty struct {
	// The caching keys for a resolver that has caching activated.
	//
	// Valid values are entries from the `$context.arguments` , `$context.source` , and `$context.identity` maps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-cachingconfig.html#cfn-appsync-resolver-cachingconfig-cachingkeys
	//
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
	// The TTL in seconds for a resolver that has caching activated.
	//
	// Valid values are 1–3,600 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-cachingconfig.html#cfn-appsync-resolver-cachingconfig-ttl
	//
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}


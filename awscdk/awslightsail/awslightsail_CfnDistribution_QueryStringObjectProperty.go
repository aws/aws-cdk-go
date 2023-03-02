package awslightsail


// `QueryStringObject` is a property of the [CacheSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachesettings.html) property. It describes the query string parameters that an Amazon Lightsail content delivery network (CDN) distribution to bases caching on.
//
// For the query strings that you specify, your distribution caches separate versions of the specified content based on the query string values in viewer requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringObjectProperty := &queryStringObjectProperty{
//   	option: jsii.Boolean(false),
//   	queryStringsAllowList: []*string{
//   		jsii.String("queryStringsAllowList"),
//   	},
//   }
//
type CfnDistribution_QueryStringObjectProperty struct {
	// Indicates whether the distribution forwards and caches based on query strings.
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// The specific query strings that the distribution forwards to the origin.
	//
	// Your distribution caches content based on the specified query strings.
	//
	// If the `option` parameter is true, then your distribution forwards all query strings, regardless of what you specify using the `QueryStringsAllowList` parameter.
	QueryStringsAllowList *[]*string `field:"optional" json:"queryStringsAllowList" yaml:"queryStringsAllowList"`
}


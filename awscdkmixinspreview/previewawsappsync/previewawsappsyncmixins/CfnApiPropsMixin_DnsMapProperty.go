package previewawsappsyncmixins


// A map of DNS names for the Api.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dnsMapProperty := &DnsMapProperty{
//   	Http: jsii.String("http"),
//   	Realtime: jsii.String("realtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-dnsmap.html
//
type CfnApiPropsMixin_DnsMapProperty struct {
	// The domain name of the Api's HTTP endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-dnsmap.html#cfn-appsync-api-dnsmap-http
	//
	Http *string `field:"optional" json:"http" yaml:"http"`
	// The domain name of the Api's real-time endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-dnsmap.html#cfn-appsync-api-dnsmap-realtime
	//
	Realtime *string `field:"optional" json:"realtime" yaml:"realtime"`
}


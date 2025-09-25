package awsroute53


// Properties of an HTTPS ServiceMode record.
//
// Example:
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//
//   var myZone hostedZone
//   var distribution cloudFrontWebDistribution
//
//   // Alias to CloudFront target
//   // Alias to CloudFront target
//   route53.NewHttpsRecord(this, jsii.String("HttpsRecord-CloudFrontAlias"), &HttpsRecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
//   })
//   // ServiceMode (priority >= 1)
//   // ServiceMode (priority >= 1)
//   route53.NewHttpsRecord(this, jsii.String("HttpsRecord-ServiceMode"), &HttpsRecordProps{
//   	Zone: myZone,
//   	Values: []httpsRecordValue{
//   		route53.*httpsRecordValue_Service(&HttpsRecordServiceModeProps{
//   			Alpn: []alpn{
//   				route53.*alpn_H3(),
//   				route53.*alpn_H2(),
//   			},
//   		}),
//   	},
//   })
//   // AliasMode (priority = 0)
//   // AliasMode (priority = 0)
//   route53.NewHttpsRecord(this, jsii.String("HttpsRecord-AliasMode"), &HttpsRecordProps{
//   	Zone: myZone,
//   	Values: []*httpsRecordValue{
//   		route53.*httpsRecordValue_Alias(jsii.String("service.example.com")),
//   	},
//   })
//
type HttpsRecordServiceModeProps struct {
	// Indicates the set of Application-Layer Protocol Negotiation (ALPN) protocol identifiers and associated transport protocols supported by this service endpoint.
	// Default: - No ALPN protocol identifiers.
	//
	Alpn *[]Alpn `field:"optional" json:"alpn" yaml:"alpn"`
	// Conveys that clients may use to reach the service.
	// Default: - No hints.
	//
	Ipv4hint *[]*string `field:"optional" json:"ipv4hint" yaml:"ipv4hint"`
	// Conveys that clients may use to reach the service.
	// Default: - No hints.
	//
	Ipv6hint *[]*string `field:"optional" json:"ipv6hint" yaml:"ipv6hint"`
	// Indicates mandatory keys.
	// Default: - No mandatory keys.
	//
	Mandatory *[]*string `field:"optional" json:"mandatory" yaml:"mandatory"`
	// Indicates no default ALPN protocol identifiers.
	//
	// The `alpn` parameter must be supplied together.
	// Default: false.
	//
	NoDefaultAlpn *bool `field:"optional" json:"noDefaultAlpn" yaml:"noDefaultAlpn"`
	// The alternative port number.
	// Default: - Use the default port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The priority.
	// Default: 1.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The domain name of the alternative endpoint.
	// Default: '.' - The record name of the record itself
	//
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
}


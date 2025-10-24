package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The ALPN protocol identifier.
//
// Example:
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//
//   var myZone HostedZone
//   var distribution CloudFrontWebDistribution
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
//   	Values: []HttpsRecordValue{
//   		route53.HttpsRecordValue_Service(&HttpsRecordServiceModeProps{
//   			Alpn: []Alpn{
//   				route53.Alpn_H3(),
//   				route53.Alpn_H2(),
//   			},
//   		}),
//   	},
//   })
//   // AliasMode (priority = 0)
//   // AliasMode (priority = 0)
//   route53.NewHttpsRecord(this, jsii.String("HttpsRecord-AliasMode"), &HttpsRecordProps{
//   	Zone: myZone,
//   	Values: []HttpsRecordValue{
//   		route53.HttpsRecordValue_Alias(jsii.String("service.example.com")),
//   	},
//   })
//
type Alpn interface {
	// The ALPN protocol identifier.
	Protocol() *string
}

// The jsii proxy struct for Alpn
type jsiiProxy_Alpn struct {
	_ byte // padding
}

func (j *jsiiProxy_Alpn) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}


// A custom ALPN protocol identifier.
func Alpn_Of(protocol *string) Alpn {
	_init_.Initialize()

	if err := validateAlpn_OfParameters(protocol); err != nil {
		panic(err)
	}
	var returns Alpn

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.Alpn",
		"of",
		[]interface{}{protocol},
		&returns,
	)

	return returns
}

func Alpn_H2() Alpn {
	_init_.Initialize()
	var returns Alpn
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.Alpn",
		"H2",
		&returns,
	)
	return returns
}

func Alpn_H3() Alpn {
	_init_.Initialize()
	var returns Alpn
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.Alpn",
		"H3",
		&returns,
	)
	return returns
}

func Alpn_HTTP1_1() Alpn {
	_init_.Initialize()
	var returns Alpn
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.Alpn",
		"HTTP1_1",
		&returns,
	)
	return returns
}


package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an HTTPS record value.
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
type HttpsRecordValue interface {
	// Returns the string representation of SVCB and HTTPS record value.
	ToString() *string
}

// The jsii proxy struct for HttpsRecordValue
type jsiiProxy_HttpsRecordValue struct {
	_ byte // padding
}

// An HTTPS AliasMode record value.
func HttpsRecordValue_Alias(targetName *string) HttpsRecordValue {
	_init_.Initialize()

	if err := validateHttpsRecordValue_AliasParameters(targetName); err != nil {
		panic(err)
	}
	var returns HttpsRecordValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HttpsRecordValue",
		"alias",
		[]interface{}{targetName},
		&returns,
	)

	return returns
}

// An HTTPS ServiceMode record value.
func HttpsRecordValue_Service(props *HttpsRecordServiceModeProps) HttpsRecordValue {
	_init_.Initialize()

	if err := validateHttpsRecordValue_ServiceParameters(props); err != nil {
		panic(err)
	}
	var returns HttpsRecordValue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.HttpsRecordValue",
		"service",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpsRecordValue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


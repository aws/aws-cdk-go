package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// Viewer certificate configuration class.
//
// Example:
//   s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))
//
//   distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []SourceConfiguration{
//   		&SourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: *S3BucketSource,
//   			},
//   			Behaviors: []Behavior{
//   				&Behavior{
//   					IsDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	ViewerCertificate: cloudfront.ViewerCertificate_FromIamCertificate(jsii.String("certificateId"), &ViewerCertificateOptions{
//   		Aliases: []*string{
//   			jsii.String("example.com"),
//   		},
//   		SecurityPolicy: cloudfront.SecurityPolicyProtocol_SSL_V3,
//   		 // default
//   		SslMethod: cloudfront.SSLMethod_SNI,
//   	}),
//   })
//
type ViewerCertificate interface {
	Aliases() *[]*string
	Props() *CfnDistribution_ViewerCertificateProperty
}

// The jsii proxy struct for ViewerCertificate
type jsiiProxy_ViewerCertificate struct {
	_ byte // padding
}

func (j *jsiiProxy_ViewerCertificate) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ViewerCertificate) Props() *CfnDistribution_ViewerCertificateProperty {
	var returns *CfnDistribution_ViewerCertificateProperty
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Generate an AWS Certificate Manager (ACM) viewer certificate configuration.
func ViewerCertificate_FromAcmCertificate(certificate awscertificatemanager.ICertificate, options *ViewerCertificateOptions) ViewerCertificate {
	_init_.Initialize()

	if err := validateViewerCertificate_FromAcmCertificateParameters(certificate, options); err != nil {
		panic(err)
	}
	var returns ViewerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificate",
		"fromAcmCertificate",
		[]interface{}{certificate, options},
		&returns,
	)

	return returns
}

// Generate a viewer certificate configuration using the CloudFront default certificate (e.g. d111111abcdef8.cloudfront.net) and a `SecurityPolicyProtocol.TLS_V1` security policy.
func ViewerCertificate_FromCloudFrontDefaultCertificate(aliases ...*string) ViewerCertificate {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range aliases {
		args = append(args, a)
	}

	var returns ViewerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificate",
		"fromCloudFrontDefaultCertificate",
		args,
		&returns,
	)

	return returns
}

// Generate an IAM viewer certificate configuration.
func ViewerCertificate_FromIamCertificate(iamCertificateId *string, options *ViewerCertificateOptions) ViewerCertificate {
	_init_.Initialize()

	if err := validateViewerCertificate_FromIamCertificateParameters(iamCertificateId, options); err != nil {
		panic(err)
	}
	var returns ViewerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificate",
		"fromIamCertificate",
		[]interface{}{iamCertificateId, options},
		&returns,
	)

	return returns
}


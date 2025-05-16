package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for how CloudFront signs requests.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   oac := cloudfront.NewS3OriginAccessControl(this, jsii.String("MyOAC"), &S3OriginAccessControlProps{
//   	Signing: cloudfront.Signing_SIGV4_NO_OVERRIDE(),
//   })
//   s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
//   	OriginAccessControl: oac,
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: s3Origin,
//   	},
//   })
//
type Signing interface {
	// Which requests CloudFront signs.
	Behavior() SigningBehavior
	// The signing protocol.
	Protocol() SigningProtocol
}

// The jsii proxy struct for Signing
type jsiiProxy_Signing struct {
	_ byte // padding
}

func (j *jsiiProxy_Signing) Behavior() SigningBehavior {
	var returns SigningBehavior
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Signing) Protocol() SigningProtocol {
	var returns SigningProtocol
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}


func NewSigning(protocol SigningProtocol, behavior SigningBehavior) Signing {
	_init_.Initialize()

	if err := validateNewSigningParameters(protocol, behavior); err != nil {
		panic(err)
	}
	j := jsiiProxy_Signing{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Signing",
		[]interface{}{protocol, behavior},
		&j,
	)

	return &j
}

func NewSigning_Override(s Signing, protocol SigningProtocol, behavior SigningBehavior) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Signing",
		[]interface{}{protocol, behavior},
		s,
	)
}

func Signing_NEVER() Signing {
	_init_.Initialize()
	var returns Signing
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.Signing",
		"NEVER",
		&returns,
	)
	return returns
}

func Signing_SIGV4_ALWAYS() Signing {
	_init_.Initialize()
	var returns Signing
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.Signing",
		"SIGV4_ALWAYS",
		&returns,
	)
	return returns
}

func Signing_SIGV4_NO_OVERRIDE() Signing {
	_init_.Initialize()
	var returns Signing
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.Signing",
		"SIGV4_NO_OVERRIDE",
		&returns,
	)
	return returns
}


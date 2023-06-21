package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin that is backed by an S3 bucket.
//
// If the bucket is configured for website hosting, this origin will be configured to use the bucket as an
// HTTP server origin and will use the bucket's configured website redirects and error handling. Otherwise,
// the origin is created as a bucket origin and will use CloudFront's redirect and error handling.
//
// Example:
//   // Adding an existing Lambda@Edge function created in a different stack
//   // to a CloudFront distribution.
//   var s3Bucket bucket
//
//   functionVersion := lambda.Version_FromVersionArn(this, jsii.String("Version"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:functionName:1"))
//
//   cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(s3Bucket),
//   		EdgeLambdas: []edgeLambda{
//   			&edgeLambda{
//   				FunctionVersion: *FunctionVersion,
//   				EventType: cloudfront.LambdaEdgeEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type S3Origin interface {
	awscloudfront.IOrigin
	// The method called when a given Origin is added (for the first time) to a Distribution.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
}

// The jsii proxy struct for S3Origin
type jsiiProxy_S3Origin struct {
	internal.Type__awscloudfrontIOrigin
}

func NewS3Origin(bucket awss3.IBucket, props *S3OriginProps) S3Origin {
	_init_.Initialize()

	if err := validateNewS3OriginParameters(bucket, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Origin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.S3Origin",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

func NewS3Origin_Override(s S3Origin, bucket awss3.IBucket, props *S3OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.S3Origin",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3Origin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := s.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}


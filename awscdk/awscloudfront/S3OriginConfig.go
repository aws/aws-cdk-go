package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// S3 origin configuration for CloudFront.
//
// Example:
//   var sourceBucket Bucket
//
//   viewerCertificate := cloudfront.ViewerCertificate_FromIamCertificate(jsii.String("MYIAMROLEIDENTIFIER"), &ViewerCertificateOptions{
//   	Aliases: []*string{
//   		jsii.String("MYALIAS"),
//   	},
//   })
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []SourceConfiguration{
//   		&SourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: sourceBucket,
//   			},
//   			Behaviors: []Behavior{
//   				&Behavior{
//   					IsDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	ViewerCertificate: viewerCertificate,
//   })
//
type S3OriginConfig struct {
	// The source bucket to serve content from.
	S3BucketSource awss3.IBucket `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// The optional Origin Access Identity of the origin identity cloudfront will use when calling your s3 bucket.
	// Default: No Origin Access Identity which requires the S3 bucket to be public accessible.
	//
	OriginAccessIdentity IOriginAccessIdentity `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
	// Any additional headers to pass to the origin.
	// Default: - No additional headers are passed.
	//
	OriginHeaders *map[string]*string `field:"optional" json:"originHeaders" yaml:"originHeaders"`
	// The relative path to the origin root to use for sources.
	// Default: /.
	//
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// Default: - origin shield not enabled.
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
}


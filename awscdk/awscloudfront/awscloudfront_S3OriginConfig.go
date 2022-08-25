package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// S3 origin configuration for CloudFront.
//
// Example:
//   var sourceBucket bucket
//
//   viewerCertificate := cloudfront.viewerCertificate.fromIamCertificate(jsii.String("MYIAMROLEIDENTIFIER"), &viewerCertificateOptions{
//   	aliases: []*string{
//   		jsii.String("MYALIAS"),
//   	},
//   })
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
//   	originConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			s3OriginSource: &s3OriginConfig{
//   				s3BucketSource: sourceBucket,
//   			},
//   			behaviors: []behavior{
//   				&behavior{
//   					isDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	viewerCertificate: viewerCertificate,
//   })
//
type S3OriginConfig struct {
	// The source bucket to serve content from.
	S3BucketSource awss3.IBucket `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// The optional Origin Access Identity of the origin identity cloudfront will use when calling your s3 bucket.
	OriginAccessIdentity IOriginAccessIdentity `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
	// Any additional headers to pass to the origin.
	OriginHeaders *map[string]*string `field:"optional" json:"originHeaders" yaml:"originHeaders"`
	// The relative path to the origin root to use for sources.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
}


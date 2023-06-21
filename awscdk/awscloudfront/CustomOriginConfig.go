package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A custom origin configuration.
//
// Example:
//   var sourceBucket bucket
//   var oai originAccessIdentity
//
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: sourceBucket,
//   				OriginAccessIdentity: oai,
//   			},
//   			Behaviors: []behavior{
//   				&behavior{
//   					IsDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   		&sourceConfiguration{
//   			CustomOriginSource: &CustomOriginConfig{
//   				DomainName: jsii.String("MYALIAS"),
//   			},
//   			Behaviors: []*behavior{
//   				&behavior{
//   					PathPattern: jsii.String("/somewhere"),
//   				},
//   			},
//   		},
//   	},
//   })
//
type CustomOriginConfig struct {
	// The domain name of the custom origin.
	//
	// Should not include the path - that should be in the parent SourceConfiguration.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The SSL versions to use when interacting with the origin.
	AllowedOriginSSLVersions *[]OriginSslPolicy `field:"optional" json:"allowedOriginSSLVersions" yaml:"allowedOriginSSLVersions"`
	// The origin HTTP port.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The origin HTTPS port.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Any additional headers to pass to the origin.
	OriginHeaders *map[string]*string `field:"optional" json:"originHeaders" yaml:"originHeaders"`
	// The keep alive timeout when making calls in seconds.
	OriginKeepaliveTimeout awscdk.Duration `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// The relative path to the origin root to use for sources.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// The protocol (http or https) policy to use when interacting with the origin.
	OriginProtocolPolicy OriginProtocolPolicy `field:"optional" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// The read timeout when calling the origin in seconds.
	OriginReadTimeout awscdk.Duration `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
}


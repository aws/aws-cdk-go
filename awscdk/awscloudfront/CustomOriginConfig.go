package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A custom origin configuration.
//
// Example:
//   var sourceBucket Bucket
//   var oai OriginAccessIdentity
//
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &CloudFrontWebDistributionProps{
//   	OriginConfigs: []SourceConfiguration{
//   		&SourceConfiguration{
//   			S3OriginSource: &S3OriginConfig{
//   				S3BucketSource: sourceBucket,
//   				OriginAccessIdentity: oai,
//   			},
//   			Behaviors: []Behavior{
//   				&Behavior{
//   					IsDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   		&SourceConfiguration{
//   			CustomOriginSource: &CustomOriginConfig{
//   				DomainName: jsii.String("MYALIAS"),
//   			},
//   			Behaviors: []Behavior{
//   				&Behavior{
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
	// Default: OriginSslPolicy.TLS_V1_2
	//
	AllowedOriginSSLVersions *[]OriginSslPolicy `field:"optional" json:"allowedOriginSSLVersions" yaml:"allowedOriginSSLVersions"`
	// The origin HTTP port.
	// Default: 80.
	//
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The origin HTTPS port.
	// Default: 443.
	//
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Any additional headers to pass to the origin.
	// Default: - No additional headers are passed.
	//
	OriginHeaders *map[string]*string `field:"optional" json:"originHeaders" yaml:"originHeaders"`
	// The keep alive timeout when making calls in seconds.
	// Default: Duration.seconds(5)
	//
	OriginKeepaliveTimeout awscdk.Duration `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// The relative path to the origin root to use for sources.
	// Default: /.
	//
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// The protocol (http or https) policy to use when interacting with the origin.
	// Default: OriginProtocolPolicy.HttpsOnly
	//
	OriginProtocolPolicy OriginProtocolPolicy `field:"optional" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// The read timeout when calling the origin in seconds.
	// Default: Duration.seconds(30)
	//
	OriginReadTimeout awscdk.Duration `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// Default: - origin shield not enabled.
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
}


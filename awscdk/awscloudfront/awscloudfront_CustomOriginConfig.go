package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// A custom origin configuration.
//
// Example:
//   var sourceBucket bucket
//   var oai originAccessIdentity
//
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
//   	originConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			s3OriginSource: &s3OriginConfig{
//   				s3BucketSource: sourceBucket,
//   				originAccessIdentity: oai,
//   			},
//   			behaviors: []behavior{
//   				&behavior{
//   					isDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   		&sourceConfiguration{
//   			customOriginSource: &customOriginConfig{
//   				domainName: jsii.String("MYALIAS"),
//   			},
//   			behaviors: []*behavior{
//   				&behavior{
//   					pathPattern: jsii.String("/somewhere"),
//   				},
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type CustomOriginConfig struct {
	// The domain name of the custom origin.
	//
	// Should not include the path - that should be in the parent SourceConfiguration.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The SSL versions to use when interacting with the origin.
	// Experimental.
	AllowedOriginSSLVersions *[]OriginSslPolicy `field:"optional" json:"allowedOriginSSLVersions" yaml:"allowedOriginSSLVersions"`
	// The origin HTTP port.
	// Experimental.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The origin HTTPS port.
	// Experimental.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Any additional headers to pass to the origin.
	// Experimental.
	OriginHeaders *map[string]*string `field:"optional" json:"originHeaders" yaml:"originHeaders"`
	// The keep alive timeout when making calls in seconds.
	// Experimental.
	OriginKeepaliveTimeout awscdk.Duration `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// The relative path to the origin root to use for sources.
	// Experimental.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// The protocol (http or https) policy to use when interacting with the origin.
	// Experimental.
	OriginProtocolPolicy OriginProtocolPolicy `field:"optional" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// The read timeout when calling the origin in seconds.
	// Experimental.
	OriginReadTimeout awscdk.Duration `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// Experimental.
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
}


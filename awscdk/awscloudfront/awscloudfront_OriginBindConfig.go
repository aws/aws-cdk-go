package awscloudfront


// The struct returned from {@link IOrigin.bind}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var origin iOrigin
//
//   originBindConfig := &originBindConfig{
//   	failoverConfig: &originFailoverConfig{
//   		failoverOrigin: origin,
//
//   		// the properties below are optional
//   		statusCodes: []*f64{
//   			jsii.Number(123),
//   		},
//   	},
//   	originProperty: &originProperty{
//   		domainName: jsii.String("domainName"),
//   		id: jsii.String("id"),
//
//   		// the properties below are optional
//   		connectionAttempts: jsii.Number(123),
//   		connectionTimeout: jsii.Number(123),
//   		customOriginConfig: &customOriginConfigProperty{
//   			originProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   			// the properties below are optional
//   			httpPort: jsii.Number(123),
//   			httpsPort: jsii.Number(123),
//   			originKeepaliveTimeout: jsii.Number(123),
//   			originReadTimeout: jsii.Number(123),
//   			originSslProtocols: []*string{
//   				jsii.String("originSslProtocols"),
//   			},
//   		},
//   		originAccessControlId: jsii.String("originAccessControlId"),
//   		originCustomHeaders: []interface{}{
//   			&originCustomHeaderProperty{
//   				headerName: jsii.String("headerName"),
//   				headerValue: jsii.String("headerValue"),
//   			},
//   		},
//   		originPath: jsii.String("originPath"),
//   		originShield: &originShieldProperty{
//   			enabled: jsii.Boolean(false),
//   			originShieldRegion: jsii.String("originShieldRegion"),
//   		},
//   		s3OriginConfig: &s3OriginConfigProperty{
//   			originAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   	},
//   }
//
type OriginBindConfig struct {
	// The failover configuration for this Origin.
	FailoverConfig *OriginFailoverConfig `field:"optional" json:"failoverConfig" yaml:"failoverConfig"`
	// The CloudFormation OriginProperty configuration for this Origin.
	OriginProperty *CfnDistribution_OriginProperty `field:"optional" json:"originProperty" yaml:"originProperty"`
}


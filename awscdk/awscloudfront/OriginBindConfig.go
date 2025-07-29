package awscloudfront


// The struct returned from `IOrigin.bind`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var origin iOrigin
//
//   originBindConfig := &OriginBindConfig{
//   	FailoverConfig: &OriginFailoverConfig{
//   		FailoverOrigin: origin,
//
//   		// the properties below are optional
//   		StatusCodes: []*f64{
//   			jsii.Number(123),
//   		},
//   	},
//   	OriginProperty: &OriginProperty{
//   		DomainName: jsii.String("domainName"),
//   		Id: jsii.String("id"),
//
//   		// the properties below are optional
//   		ConnectionAttempts: jsii.Number(123),
//   		ConnectionTimeout: jsii.Number(123),
//   		CustomOriginConfig: &CustomOriginConfigProperty{
//   			OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   			// the properties below are optional
//   			HttpPort: jsii.Number(123),
//   			HttpsPort: jsii.Number(123),
//   			OriginKeepaliveTimeout: jsii.Number(123),
//   			OriginReadTimeout: jsii.Number(123),
//   			OriginSslProtocols: []*string{
//   				jsii.String("originSslProtocols"),
//   			},
//   		},
//   		OriginAccessControlId: jsii.String("originAccessControlId"),
//   		OriginCustomHeaders: []interface{}{
//   			&OriginCustomHeaderProperty{
//   				HeaderName: jsii.String("headerName"),
//   				HeaderValue: jsii.String("headerValue"),
//   			},
//   		},
//   		OriginPath: jsii.String("originPath"),
//   		OriginShield: &OriginShieldProperty{
//   			Enabled: jsii.Boolean(false),
//   			OriginShieldRegion: jsii.String("originShieldRegion"),
//   		},
//   		ResponseCompletionTimeout: jsii.Number(123),
//   		S3OriginConfig: &S3OriginConfigProperty{
//   			OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   			OriginReadTimeout: jsii.Number(123),
//   		},
//   		VpcOriginConfig: &VpcOriginConfigProperty{
//   			VpcOriginId: jsii.String("vpcOriginId"),
//
//   			// the properties below are optional
//   			OriginKeepaliveTimeout: jsii.Number(123),
//   			OriginReadTimeout: jsii.Number(123),
//   		},
//   	},
//   	SelectionCriteria: awscdk.Aws_cloudfront.OriginSelectionCriteria_DEFAULT,
//   }
//
type OriginBindConfig struct {
	// The failover configuration for this Origin.
	// Default: - nothing is returned.
	//
	FailoverConfig *OriginFailoverConfig `field:"optional" json:"failoverConfig" yaml:"failoverConfig"`
	// The CloudFormation OriginProperty configuration for this Origin.
	// Default: - nothing is returned.
	//
	OriginProperty *CfnDistribution_OriginProperty `field:"optional" json:"originProperty" yaml:"originProperty"`
	// The selection criteria for how your origins are selected.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/high_availability_origin_failover.html#concept_origin_groups.creating
	//
	// Default: - OriginSelectionCriteria.DEFAULT
	//
	SelectionCriteria OriginSelectionCriteria `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}


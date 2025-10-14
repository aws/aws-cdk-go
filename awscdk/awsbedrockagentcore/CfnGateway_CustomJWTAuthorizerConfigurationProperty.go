package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customJWTAuthorizerConfigurationProperty := &CustomJWTAuthorizerConfigurationProperty{
//   	DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   	// the properties below are optional
//   	AllowedAudience: []*string{
//   		jsii.String("allowedAudience"),
//   	},
//   	AllowedClients: []*string{
//   		jsii.String("allowedClients"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html
//
type CfnGateway_CustomJWTAuthorizerConfigurationProperty struct {
	// The discovery URL for the authorizer configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-discoveryurl
	//
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// The allowed audience authorized for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-allowedaudience
	//
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.html#cfn-bedrockagentcore-gateway-customjwtauthorizerconfiguration-allowedclients
	//
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
}


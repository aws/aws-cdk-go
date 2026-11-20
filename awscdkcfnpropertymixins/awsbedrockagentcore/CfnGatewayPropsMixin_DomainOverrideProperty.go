package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   domainOverrideProperty := &DomainOverrideProperty{
//   	AlternateDomains: []*string{
//   		jsii.String("alternateDomains"),
//   	},
//   	DefaultDomain: jsii.String("defaultDomain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-domainoverride.html
//
type CfnGatewayPropsMixin_DomainOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-domainoverride.html#cfn-bedrockagentcore-gateway-domainoverride-alternatedomains
	//
	AlternateDomains *[]*string `field:"optional" json:"alternateDomains" yaml:"alternateDomains"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-domainoverride.html#cfn-bedrockagentcore-gateway-domainoverride-defaultdomain
	//
	DefaultDomain *string `field:"optional" json:"defaultDomain" yaml:"defaultDomain"`
}


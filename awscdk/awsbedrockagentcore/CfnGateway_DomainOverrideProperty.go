package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainOverrideProperty := &DomainOverrideProperty{
//   	DefaultDomain: jsii.String("defaultDomain"),
//
//   	// the properties below are optional
//   	AlternateDomains: []*string{
//   		jsii.String("alternateDomains"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-domainoverride.html
//
type CfnGateway_DomainOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-domainoverride.html#cfn-bedrockagentcore-gateway-domainoverride-defaultdomain
	//
	DefaultDomain *string `field:"required" json:"defaultDomain" yaml:"defaultDomain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-domainoverride.html#cfn-bedrockagentcore-gateway-domainoverride-alternatedomains
	//
	AlternateDomains *[]*string `field:"optional" json:"alternateDomains" yaml:"alternateDomains"`
}


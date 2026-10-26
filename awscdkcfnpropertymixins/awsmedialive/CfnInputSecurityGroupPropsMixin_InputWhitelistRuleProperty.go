package awsmedialive


// Whitelist rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inputWhitelistRuleProperty := &InputWhitelistRuleProperty{
//   	Cidr: jsii.String("cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-inputsecuritygroup-inputwhitelistrule.html
//
type CfnInputSecurityGroupPropsMixin_InputWhitelistRuleProperty struct {
	// The IPv4 CIDR that's whitelisted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-inputsecuritygroup-inputwhitelistrule.html#cfn-medialive-inputsecuritygroup-inputwhitelistrule-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}


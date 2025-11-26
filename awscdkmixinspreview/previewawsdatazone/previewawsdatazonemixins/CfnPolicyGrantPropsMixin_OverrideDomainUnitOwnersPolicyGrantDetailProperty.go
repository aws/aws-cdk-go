package previewawsdatazonemixins


// The grant details of the override domain unit owners policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   overrideDomainUnitOwnersPolicyGrantDetailProperty := &OverrideDomainUnitOwnersPolicyGrantDetailProperty{
//   	IncludeChildDomainUnits: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-overridedomainunitownerspolicygrantdetail.html
//
type CfnPolicyGrantPropsMixin_OverrideDomainUnitOwnersPolicyGrantDetailProperty struct {
	// Specifies whether the policy is inherited by child domain units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-overridedomainunitownerspolicygrantdetail.html#cfn-datazone-policygrant-overridedomainunitownerspolicygrantdetail-includechilddomainunits
	//
	IncludeChildDomainUnits interface{} `field:"optional" json:"includeChildDomainUnits" yaml:"includeChildDomainUnits"`
}


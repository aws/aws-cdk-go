package awsdatazone


// The details of the policy grant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createGlossaryPolicyGrantDetailProperty := &CreateGlossaryPolicyGrantDetailProperty{
//   	IncludeChildDomainUnits: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createglossarypolicygrantdetail.html
//
type CfnPolicyGrant_CreateGlossaryPolicyGrantDetailProperty struct {
	// Specifies whether the policy grant is applied to child domain units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createglossarypolicygrantdetail.html#cfn-datazone-policygrant-createglossarypolicygrantdetail-includechilddomainunits
	//
	IncludeChildDomainUnits interface{} `field:"optional" json:"includeChildDomainUnits" yaml:"includeChildDomainUnits"`
}


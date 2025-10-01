package awsdatazone


// The details of the policy grant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createFormTypePolicyGrantDetailProperty := &CreateFormTypePolicyGrantDetailProperty{
//   	IncludeChildDomainUnits: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createformtypepolicygrantdetail.html
//
type CfnPolicyGrant_CreateFormTypePolicyGrantDetailProperty struct {
	// Specifies whether the policy grant is applied to child domain units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createformtypepolicygrantdetail.html#cfn-datazone-policygrant-createformtypepolicygrantdetail-includechilddomainunits
	//
	IncludeChildDomainUnits interface{} `field:"optional" json:"includeChildDomainUnits" yaml:"includeChildDomainUnits"`
}


package awsdatazone


// The details of the policy grant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createAssetTypePolicyGrantDetailProperty := &CreateAssetTypePolicyGrantDetailProperty{
//   	IncludeChildDomainUnits: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createassettypepolicygrantdetail.html
//
type CfnPolicyGrant_CreateAssetTypePolicyGrantDetailProperty struct {
	// Specifies whether the policy grant is applied to child domain units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createassettypepolicygrantdetail.html#cfn-datazone-policygrant-createassettypepolicygrantdetail-includechilddomainunits
	//
	IncludeChildDomainUnits interface{} `field:"optional" json:"includeChildDomainUnits" yaml:"includeChildDomainUnits"`
}


package previewawsdatazonemixins


// The details of the policy grant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createEnvironmentProfilePolicyGrantDetailProperty := &CreateEnvironmentProfilePolicyGrantDetailProperty{
//   	DomainUnitId: jsii.String("domainUnitId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createenvironmentprofilepolicygrantdetail.html
//
type CfnPolicyGrantPropsMixin_CreateEnvironmentProfilePolicyGrantDetailProperty struct {
	// The ID of the domain unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-createenvironmentprofilepolicygrantdetail.html#cfn-datazone-policygrant-createenvironmentprofilepolicygrantdetail-domainunitid
	//
	DomainUnitId *string `field:"optional" json:"domainUnitId" yaml:"domainUnitId"`
}


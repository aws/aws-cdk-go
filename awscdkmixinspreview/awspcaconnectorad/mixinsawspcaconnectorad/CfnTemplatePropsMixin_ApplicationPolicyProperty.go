package mixinsawspcaconnectorad


// Application policies describe what the certificate can be used for.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationPolicyProperty := &ApplicationPolicyProperty{
//   	PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   	PolicyType: jsii.String("policyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicy.html
//
type CfnTemplatePropsMixin_ApplicationPolicyProperty struct {
	// The object identifier (OID) of an application policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicy.html#cfn-pcaconnectorad-template-applicationpolicy-policyobjectidentifier
	//
	PolicyObjectIdentifier *string `field:"optional" json:"policyObjectIdentifier" yaml:"policyObjectIdentifier"`
	// The type of application policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicy.html#cfn-pcaconnectorad-template-applicationpolicy-policytype
	//
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
}


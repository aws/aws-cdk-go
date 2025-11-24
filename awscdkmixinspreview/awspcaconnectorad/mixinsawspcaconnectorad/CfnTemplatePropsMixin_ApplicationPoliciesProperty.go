package mixinsawspcaconnectorad


// Application policies describe what the certificate can be used for.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationPoliciesProperty := &ApplicationPoliciesProperty{
//   	Critical: jsii.Boolean(false),
//   	Policies: []interface{}{
//   		&ApplicationPolicyProperty{
//   			PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   			PolicyType: jsii.String("policyType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicies.html
//
type CfnTemplatePropsMixin_ApplicationPoliciesProperty struct {
	// Marks the application policy extension as critical.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicies.html#cfn-pcaconnectorad-template-applicationpolicies-critical
	//
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
	// Application policies describe what the certificate can be used for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicies.html#cfn-pcaconnectorad-template-applicationpolicies-policies
	//
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
}


package awspcaconnectorad


// Application policies describe what the certificate can be used for.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationPoliciesProperty := &ApplicationPoliciesProperty{
//   	Policies: []interface{}{
//   		&ApplicationPolicyProperty{
//   			PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   			PolicyType: jsii.String("policyType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Critical: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicies.html
//
type CfnTemplate_ApplicationPoliciesProperty struct {
	// Application policies describe what the certificate can be used for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicies.html#cfn-pcaconnectorad-template-applicationpolicies-policies
	//
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
	// Marks the application policy extension as critical.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-applicationpolicies.html#cfn-pcaconnectorad-template-applicationpolicies-critical
	//
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
}


package mixinsawsecr


// Properties for CfnRegistryPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyText interface{}
//
//   cfnRegistryPolicyMixinProps := &CfnRegistryPolicyMixinProps{
//   	PolicyText: policyText,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html
//
type CfnRegistryPolicyMixinProps struct {
	// The JSON policy text for your registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html#cfn-ecr-registrypolicy-policytext
	//
	PolicyText interface{} `field:"optional" json:"policyText" yaml:"policyText"`
}


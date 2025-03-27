package awsecr


// Properties for defining a `CfnRegistryPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyText interface{}
//
//   cfnRegistryPolicyProps := &CfnRegistryPolicyProps{
//   	PolicyText: policyText,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html
//
type CfnRegistryPolicyProps struct {
	// The JSON policy text for your registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-registrypolicy.html#cfn-ecr-registrypolicy-policytext
	//
	PolicyText interface{} `field:"required" json:"policyText" yaml:"policyText"`
}


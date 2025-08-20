package awseventschemas


// Properties for defining a `CfnRegistryPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnRegistryPolicyProps := &CfnRegistryPolicyProps{
//   	Policy: policy,
//   	RegistryName: jsii.String("registryName"),
//
//   	// the properties below are optional
//   	RevisionId: jsii.String("revisionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registrypolicy.html
//
type CfnRegistryPolicyProps struct {
	// A resource-based policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registrypolicy.html#cfn-eventschemas-registrypolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The name of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registrypolicy.html#cfn-eventschemas-registrypolicy-registryname
	//
	RegistryName *string `field:"required" json:"registryName" yaml:"registryName"`
	// The revision ID of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-registrypolicy.html#cfn-eventschemas-registrypolicy-revisionid
	//
	RevisionId *string `field:"optional" json:"revisionId" yaml:"revisionId"`
}


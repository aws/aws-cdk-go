package awsfms


// The resource tags that AWS Firewall Manager uses to determine if a particular resource should be included or excluded from the AWS Firewall Manager policy.
//
// Tags enable you to categorize your AWS resources in different ways, for example, by purpose, owner, or environment. Each tag consists of a key and an optional value. Firewall Manager combines the tags with "AND" so that, if you add more than one tag to a policy scope, a resource must have all the specified tags to be included or excluded. For more information, see [Working with Tag Editor](https://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/tag-editor.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceTagProperty := &resourceTagProperty{
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	value: jsii.String("value"),
//   }
//
type CfnPolicy_ResourceTagProperty struct {
	// The resource tag key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The resource tag value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


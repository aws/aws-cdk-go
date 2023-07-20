package awsfms


// A collection of key:value pairs associated with an AWS resource.
//
// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyTagProperty := &PolicyTagProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policytag.html
//
type CfnPolicy_PolicyTagProperty struct {
	// Part of the key:value pair that defines a tag.
	//
	// You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policytag.html#cfn-fms-policy-policytag-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// Part of the key:value pair that defines a tag.
	//
	// You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policytag.html#cfn-fms-policy-policytag-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


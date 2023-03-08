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
//   policyTagProperty := &policyTagProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnPolicy_PolicyTagProperty struct {
	// Part of the key:value pair that defines a tag.
	//
	// You can use a tag key to describe a category of information, such as "customer." Tag keys are case-sensitive.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Part of the key:value pair that defines a tag.
	//
	// You can use a tag value to describe a specific value within a category, such as "companyA" or "companyB." Tag values are case-sensitive.
	Value *string `field:"required" json:"value" yaml:"value"`
}


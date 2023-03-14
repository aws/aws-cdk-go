package awsconfig


// Defines which resources trigger an evaluation for an AWS Config rule.
//
// The scope can include one or more resource types, a combination of a tag key and value, or a combination of one resource type and one resource ID. Specify a scope to constrain which resources trigger an evaluation for a rule. Otherwise, evaluations for the rule are triggered when any resource in your recording group changes in configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scopeProperty := &scopeProperty{
//   	complianceResourceId: jsii.String("complianceResourceId"),
//   	complianceResourceTypes: []*string{
//   		jsii.String("complianceResourceTypes"),
//   	},
//   	tagKey: jsii.String("tagKey"),
//   	tagValue: jsii.String("tagValue"),
//   }
//
type CfnConfigRule_ScopeProperty struct {
	// The ID of the only AWS resource that you want to trigger an evaluation for the rule.
	//
	// If you specify a resource ID, you must specify one resource type for `ComplianceResourceTypes` .
	ComplianceResourceId *string `field:"optional" json:"complianceResourceId" yaml:"complianceResourceId"`
	// The resource types of only those AWS resources that you want to trigger an evaluation for the rule.
	//
	// You can only specify one type if you also specify a resource ID for `ComplianceResourceId` .
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// The tag key that is applied to only those AWS resources that you want to trigger an evaluation for the rule.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule.
	//
	// If you specify a value for `TagValue` , you must also specify a value for `TagKey` .
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}


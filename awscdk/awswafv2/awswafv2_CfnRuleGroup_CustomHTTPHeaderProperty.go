package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customHTTPHeaderProperty := &customHTTPHeaderProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnRuleGroup_CustomHTTPHeaderProperty struct {
	// `CfnRuleGroup.CustomHTTPHeaderProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnRuleGroup.CustomHTTPHeaderProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}


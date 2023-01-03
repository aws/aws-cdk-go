package awswafv2


// Inspect one query argument in the web request, identified by name, for example *UserName* or *SalesRegion* .
//
// The name isn't case sensitive.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example JSON: `"SingleQueryArgument": { "Name": "myArgument" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleQueryArgumentProperty := &singleQueryArgumentProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnRuleGroup_SingleQueryArgumentProperty struct {
	// The name of the query argument to inspect.
	Name *string `field:"required" json:"name" yaml:"name"`
}


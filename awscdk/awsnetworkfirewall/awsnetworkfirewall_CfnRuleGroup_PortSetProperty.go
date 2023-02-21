package awsnetworkfirewall


// A set of port ranges for use in the rules in a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portSetProperty := &portSetProperty{
//   	definition: []*string{
//   		jsii.String("definition"),
//   	},
//   }
//
type CfnRuleGroup_PortSetProperty struct {
	// The set of port ranges.
	Definition *[]*string `field:"optional" json:"definition" yaml:"definition"`
}


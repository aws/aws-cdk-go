package awsbudgets


// The Amazon EC2 Systems Manager ( SSM ) action definition details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmActionDefinitionProperty := &ssmActionDefinitionProperty{
//   	instanceIds: []*string{
//   		jsii.String("instanceIds"),
//   	},
//   	region: jsii.String("region"),
//   	subtype: jsii.String("subtype"),
//   }
//
type CfnBudgetsAction_SsmActionDefinitionProperty struct {
	// The EC2 and RDS instance IDs.
	InstanceIds *[]*string `field:"required" json:"instanceIds" yaml:"instanceIds"`
	// The Region to run the ( SSM ) document.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The action subType.
	Subtype *string `field:"required" json:"subtype" yaml:"subtype"`
}


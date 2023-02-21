package awsbudgets


// The definition is where you specify all of the type-specific parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   definitionProperty := &DefinitionProperty{
//   	IamActionDefinition: &IamActionDefinitionProperty{
//   		PolicyArn: jsii.String("policyArn"),
//
//   		// the properties below are optional
//   		Groups: []*string{
//   			jsii.String("groups"),
//   		},
//   		Roles: []*string{
//   			jsii.String("roles"),
//   		},
//   		Users: []*string{
//   			jsii.String("users"),
//   		},
//   	},
//   	ScpActionDefinition: &ScpActionDefinitionProperty{
//   		PolicyId: jsii.String("policyId"),
//   		TargetIds: []*string{
//   			jsii.String("targetIds"),
//   		},
//   	},
//   	SsmActionDefinition: &SsmActionDefinitionProperty{
//   		InstanceIds: []*string{
//   			jsii.String("instanceIds"),
//   		},
//   		Region: jsii.String("region"),
//   		Subtype: jsii.String("subtype"),
//   	},
//   }
//
type CfnBudgetsAction_DefinitionProperty struct {
	// The AWS Identity and Access Management ( IAM ) action definition details.
	IamActionDefinition interface{} `field:"optional" json:"iamActionDefinition" yaml:"iamActionDefinition"`
	// The service control policies (SCP) action definition details.
	ScpActionDefinition interface{} `field:"optional" json:"scpActionDefinition" yaml:"scpActionDefinition"`
	// The Amazon EC2 Systems Manager ( SSM ) action definition details.
	SsmActionDefinition interface{} `field:"optional" json:"ssmActionDefinition" yaml:"ssmActionDefinition"`
}


package awsbudgets


// The definition is where you specify all of the type-specific parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   definitionProperty := &definitionProperty{
//   	iamActionDefinition: &iamActionDefinitionProperty{
//   		policyArn: jsii.String("policyArn"),
//
//   		// the properties below are optional
//   		groups: []*string{
//   			jsii.String("groups"),
//   		},
//   		roles: []*string{
//   			jsii.String("roles"),
//   		},
//   		users: []*string{
//   			jsii.String("users"),
//   		},
//   	},
//   	scpActionDefinition: &scpActionDefinitionProperty{
//   		policyId: jsii.String("policyId"),
//   		targetIds: []*string{
//   			jsii.String("targetIds"),
//   		},
//   	},
//   	ssmActionDefinition: &ssmActionDefinitionProperty{
//   		instanceIds: []*string{
//   			jsii.String("instanceIds"),
//   		},
//   		region: jsii.String("region"),
//   		subtype: jsii.String("subtype"),
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


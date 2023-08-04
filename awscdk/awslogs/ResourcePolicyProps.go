package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties to define Cloudwatch log group resource policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyStatement policyStatement
//
//   resourcePolicyProps := &ResourcePolicyProps{
//   	PolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	ResourcePolicyName: jsii.String("resourcePolicyName"),
//   }
//
type ResourcePolicyProps struct {
	// Initial statements to add to the resource policy.
	// Default: - No statements.
	//
	PolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"policyStatements" yaml:"policyStatements"`
	// Name of the log group resource policy.
	// Default: - Uses a unique id based on the construct path.
	//
	ResourcePolicyName *string `field:"optional" json:"resourcePolicyName" yaml:"resourcePolicyName"`
}


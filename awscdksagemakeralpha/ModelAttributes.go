package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a Model resource defined outside this stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var securityGroup securityGroup
//
//   modelAttributes := &ModelAttributes{
//   	ModelArn: jsii.String("modelArn"),
//
//   	// the properties below are optional
//   	Role: role,
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// Experimental.
type ModelAttributes struct {
	// The ARN of this model.
	// Experimental.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// The IAM execution role associated with this model.
	// Default: - When not provided, any role-related operations will no-op.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The security groups for this model, if in a VPC.
	// Default: - When not provided, the connections to/from this model cannot be managed.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}


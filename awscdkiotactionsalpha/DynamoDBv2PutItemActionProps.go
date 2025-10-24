package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for the dynamodb table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_actions_alpha "github.com/aws/aws-cdk-go/awscdkiotactionsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   dynamoDBv2PutItemActionProps := &DynamoDBv2PutItemActionProps{
//   	Role: role,
//   }
//
// Experimental.
type DynamoDBv2PutItemActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


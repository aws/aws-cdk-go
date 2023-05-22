package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for a CrossAccountDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   crossAccountDestinationProps := &CrossAccountDestinationProps{
//   	Role: role,
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	DestinationName: jsii.String("destinationName"),
//   }
//
// Experimental.
type CrossAccountDestinationProps struct {
	// The role to assume that grants permissions to write to 'target'.
	//
	// The role must be assumable by 'logs.{REGION}.amazonaws.com'.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// The log destination target's ARN.
	// Experimental.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The name of the log destination.
	// Experimental.
	DestinationName *string `field:"optional" json:"destinationName" yaml:"destinationName"`
}


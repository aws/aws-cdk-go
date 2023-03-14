// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a Script content defined outside of this stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   scriptAttributes := &ScriptAttributes{
//   	ScriptArn: jsii.String("scriptArn"),
//
//   	// the properties below are optional
//   	Role: role,
//   }
//
// Experimental.
type ScriptAttributes struct {
	// The ARN of the realtime server script.
	// Experimental.
	ScriptArn *string `field:"required" json:"scriptArn" yaml:"scriptArn"`
	// The IAM role assumed by GameLift to access server script in S3.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


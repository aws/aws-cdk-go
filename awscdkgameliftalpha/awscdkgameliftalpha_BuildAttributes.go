// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a Build content defined outside of this stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   buildAttributes := &buildAttributes{
//   	buildId: jsii.String("buildId"),
//
//   	// the properties below are optional
//   	role: role,
//   }
//
// Experimental.
type BuildAttributes struct {
	// The identifier of the build.
	// Experimental.
	BuildId *string `field:"required" json:"buildId" yaml:"buildId"`
	// The IAM role assumed by GameLift to access server build in S3.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


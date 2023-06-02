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
//   buildAttributes := &BuildAttributes{
//   	BuildArn: jsii.String("buildArn"),
//   	BuildId: jsii.String("buildId"),
//   	Role: role,
//   }
//
// Experimental.
type BuildAttributes struct {
	// The ARN of the build.
	//
	// At least one of `buildArn` and `buildId` must be provided.
	// Experimental.
	BuildArn *string `field:"optional" json:"buildArn" yaml:"buildArn"`
	// The identifier of the build.
	//
	// At least one of `buildId` and `buildArn`  must be provided.
	// Experimental.
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// The IAM role assumed by GameLift to access server build in S3.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


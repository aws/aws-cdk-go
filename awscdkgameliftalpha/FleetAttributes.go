package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A full specification of a fleet that can be used to import it fluently into the CDK application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   fleetAttributes := &FleetAttributes{
//   	FleetArn: jsii.String("fleetArn"),
//   	FleetId: jsii.String("fleetId"),
//   	Role: role,
//   }
//
// Experimental.
type FleetAttributes struct {
	// The ARN of the fleet.
	//
	// At least one of `fleetArn` and `fleetId` must be provided.
	// Default: derived from `fleetId`.
	//
	// Experimental.
	FleetArn *string `field:"optional" json:"fleetArn" yaml:"fleetArn"`
	// The identifier of the fleet.
	//
	// At least one of `fleetId` and `fleetArn`  must be provided.
	// Default: derived from `fleetArn`.
	//
	// Experimental.
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// The IAM role assumed by GameLift fleet instances to access AWS ressources.
	// Default: the imported fleet cannot be granted access to other resources as an `iam.IGrantable`.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


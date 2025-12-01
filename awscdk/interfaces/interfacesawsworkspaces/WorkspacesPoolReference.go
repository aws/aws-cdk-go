package interfacesawsworkspaces


// A reference to a WorkspacesPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspacesPoolReference := &WorkspacesPoolReference{
//   	PoolArn: jsii.String("poolArn"),
//   	PoolId: jsii.String("poolId"),
//   }
//
type WorkspacesPoolReference struct {
	// The ARN of the WorkspacesPool resource.
	PoolArn *string `field:"required" json:"poolArn" yaml:"poolArn"`
	// The PoolId of the WorkspacesPool resource.
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
}


package awsworkspacesinstances


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceMaintenanceOptionsRequestProperty := &InstanceMaintenanceOptionsRequestProperty{
//   	AutoRecovery: jsii.String("autoRecovery"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemaintenanceoptionsrequest.html
//
type CfnWorkspaceInstance_InstanceMaintenanceOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemaintenanceoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemaintenanceoptionsrequest-autorecovery
	//
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}


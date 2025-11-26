package previewawsworkspacesmixins


// Describes the user capacity for the pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityProperty := &CapacityProperty{
//   	DesiredUserSessions: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-capacity.html
//
type CfnWorkspacesPoolPropsMixin_CapacityProperty struct {
	// The desired number of user sessions for the WorkSpaces in the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-capacity.html#cfn-workspaces-workspacespool-capacity-desiredusersessions
	//
	DesiredUserSessions *float64 `field:"optional" json:"desiredUserSessions" yaml:"desiredUserSessions"`
}


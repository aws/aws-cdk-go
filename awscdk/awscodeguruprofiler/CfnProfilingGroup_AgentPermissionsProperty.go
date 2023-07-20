package awscodeguruprofiler


// The agent permissions attached to this profiling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentPermissionsProperty := &AgentPermissionsProperty{
//   	Principals: []*string{
//   		jsii.String("principals"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeguruprofiler-profilinggroup-agentpermissions.html
//
type CfnProfilingGroup_AgentPermissionsProperty struct {
	// The principals for the agent permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeguruprofiler-profilinggroup-agentpermissions.html#cfn-codeguruprofiler-profilinggroup-agentpermissions-principals
	//
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
}


package awscodeguruprofiler


// The agent permissions attached to this profiling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   agentPermissionsProperty := &AgentPermissionsProperty{
//   	Principals: []interface{}{
//   		jsii.String("principals"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeguruprofiler-profilinggroup-agentpermissions.html
//
type CfnProfilingGroupPropsMixin_AgentPermissionsProperty struct {
	// The principals for the agent permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeguruprofiler-profilinggroup-agentpermissions.html#cfn-codeguruprofiler-profilinggroup-agentpermissions-principals
	//
	Principals *[]interface{} `field:"optional" json:"principals" yaml:"principals"`
}


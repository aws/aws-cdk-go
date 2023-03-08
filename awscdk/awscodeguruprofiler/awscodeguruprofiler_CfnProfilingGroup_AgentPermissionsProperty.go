package awscodeguruprofiler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentPermissionsProperty := &agentPermissionsProperty{
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   }
//
type CfnProfilingGroup_AgentPermissionsProperty struct {
	// `CfnProfilingGroup.AgentPermissionsProperty.Principals`.
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
}


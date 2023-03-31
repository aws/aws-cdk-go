package awsevents


// This parameter contains the criteria (either InstanceIds or a tag) used to specify which EC2 instances are to be sent the command.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runCommandParametersProperty := &runCommandParametersProperty{
//   	runCommandTargets: []interface{}{
//   		&runCommandTargetProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
type CfnRule_RunCommandParametersProperty struct {
	// Currently, we support including only one RunCommandTarget block, which specifies either an array of InstanceIds or a tag.
	RunCommandTargets interface{} `field:"required" json:"runCommandTargets" yaml:"runCommandTargets"`
}


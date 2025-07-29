package awsevents


// This parameter contains the criteria (either InstanceIds or a tag) used to specify which EC2 instances are to be sent the command.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runCommandParametersProperty := &RunCommandParametersProperty{
//   	RunCommandTargets: []interface{}{
//   		&RunCommandTargetProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandparameters.html
//
type CfnRule_RunCommandParametersProperty struct {
	// Currently, we support including only one RunCommandTarget block, which specifies either an array of InstanceIds or a tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandparameters.html#cfn-events-rule-runcommandparameters-runcommandtargets
	//
	RunCommandTargets interface{} `field:"required" json:"runCommandTargets" yaml:"runCommandTargets"`
}


package awsevents


// Information about the EC2 instances that are to be sent the command, specified as key-value pairs.
//
// Each `RunCommandTarget` block can include only one key, but this key may specify multiple values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runCommandTargetProperty := &RunCommandTargetProperty{
//   	Key: jsii.String("key"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html
//
type CfnRule_RunCommandTargetProperty struct {
	// Can be either `tag:` *tag-key* or `InstanceIds` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html#cfn-events-rule-runcommandtarget-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// If `Key` is `tag:` *tag-key* , `Values` is a list of tag values.
	//
	// If `Key` is `InstanceIds` , `Values` is a list of Amazon EC2 instance IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html#cfn-events-rule-runcommandtarget-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


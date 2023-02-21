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
//   runCommandTargetProperty := &runCommandTargetProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnRule_RunCommandTargetProperty struct {
	// Can be either `tag:` *tag-key* or `InstanceIds` .
	Key *string `field:"required" json:"key" yaml:"key"`
	// If `Key` is `tag:` *tag-key* , `Values` is a list of tag values.
	//
	// If `Key` is `InstanceIds` , `Values` is a list of Amazon EC2 instance IDs.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


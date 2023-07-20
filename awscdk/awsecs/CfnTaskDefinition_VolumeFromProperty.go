package awsecs


// Details on a data volume from another container in the same task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeFromProperty := &VolumeFromProperty{
//   	ReadOnly: jsii.Boolean(false),
//   	SourceContainer: jsii.String("sourceContainer"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumefrom.html
//
type CfnTaskDefinition_VolumeFromProperty struct {
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumefrom.html#cfn-ecs-taskdefinition-volumefrom-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of another container within the same task definition to mount volumes from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumefrom.html#cfn-ecs-taskdefinition-volumefrom-sourcecontainer
	//
	SourceContainer *string `field:"optional" json:"sourceContainer" yaml:"sourceContainer"`
}


package awsstepfunctions


// Properties for CfnMapRunPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMapRunMixinProps := &CfnMapRunMixinProps{
//   	ExecutionArn: jsii.String("executionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-maprun.html
//
type CfnMapRunMixinProps struct {
	// The Amazon Resource Name (ARN) that identifies the execution in which the Map Run was started.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-maprun.html#cfn-stepfunctions-maprun-executionarn
	//
	ExecutionArn *string `field:"optional" json:"executionArn" yaml:"executionArn"`
}


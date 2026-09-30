package awssagemaker


// Priority class configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   priorityClassProperty := &PriorityClassProperty{
//   	Name: jsii.String("name"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-clusterschedulerconfig-priorityclass.html
//
type CfnClusterSchedulerConfigPropsMixin_PriorityClassProperty struct {
	// Name of the priority class.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-clusterschedulerconfig-priorityclass.html#cfn-sagemaker-clusterschedulerconfig-priorityclass-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Weight of the priority class.
	//
	// Range 0-100, default 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-clusterschedulerconfig-priorityclass.html#cfn-sagemaker-clusterschedulerconfig-priorityclass-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


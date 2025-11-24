package mixinsawsbatch


// Properties for CfnServiceEnvironmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceEnvironmentMixinProps := &CfnServiceEnvironmentMixinProps{
//   	CapacityLimits: []interface{}{
//   		&CapacityLimitProperty{
//   			CapacityUnit: jsii.String("capacityUnit"),
//   			MaxCapacity: jsii.Number(123),
//   		},
//   	},
//   	ServiceEnvironmentName: jsii.String("serviceEnvironmentName"),
//   	ServiceEnvironmentType: jsii.String("serviceEnvironmentType"),
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-serviceenvironment.html
//
type CfnServiceEnvironmentMixinProps struct {
	// The capacity limits for the service environment.
	//
	// This defines the maximum resources that can be used by service jobs in this environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-serviceenvironment.html#cfn-batch-serviceenvironment-capacitylimits
	//
	CapacityLimits interface{} `field:"optional" json:"capacityLimits" yaml:"capacityLimits"`
	// The name of the service environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-serviceenvironment.html#cfn-batch-serviceenvironment-serviceenvironmentname
	//
	ServiceEnvironmentName *string `field:"optional" json:"serviceEnvironmentName" yaml:"serviceEnvironmentName"`
	// The type of service environment.
	//
	// For SageMaker Training jobs, this value is `SAGEMAKER_TRAINING` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-serviceenvironment.html#cfn-batch-serviceenvironment-serviceenvironmenttype
	//
	ServiceEnvironmentType *string `field:"optional" json:"serviceEnvironmentType" yaml:"serviceEnvironmentType"`
	// The state of the service environment.
	//
	// Valid values are `ENABLED` and `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-serviceenvironment.html#cfn-batch-serviceenvironment-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags associated with the service environment.
	//
	// Each tag consists of a key and an optional value. For more information, see [Tagging your AWS Batch resources](https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-serviceenvironment.html#cfn-batch-serviceenvironment-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


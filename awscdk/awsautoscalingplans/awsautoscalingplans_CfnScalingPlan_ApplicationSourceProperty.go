package awsautoscalingplans


// `ApplicationSource` is a property of [ScalingPlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html) that specifies the application source to use with AWS Auto Scaling ( Auto Scaling Plans ). You can create one scaling plan per application source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSourceProperty := &applicationSourceProperty{
//   	cloudFormationStackArn: jsii.String("cloudFormationStackArn"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
type CfnScalingPlan_ApplicationSourceProperty struct {
	// The Amazon Resource Name (ARN) of a CloudFormation stack.
	//
	// You must specify either a `CloudFormationStackARN` or `TagFilters` .
	CloudFormationStackArn *string `field:"optional" json:"cloudFormationStackArn" yaml:"cloudFormationStackArn"`
	// A set of tag filters (keys and values).
	//
	// Each tag filter specified must contain a key with values as optional. Each scaling plan can include up to 50 keys, and each key can include up to 20 values.
	//
	// Tags are part of the syntax that you use to specify the resources you want returned when configuring a scaling plan from the AWS Auto Scaling console. You do not need to specify valid tag filter values when you create a scaling plan with CloudFormation. The `Key` and `Values` properties can accept any value as long as the combination of values is unique across scaling plans. However, if you also want to use the AWS Auto Scaling console to edit the scaling plan, then you must specify valid values.
	//
	// You must specify either a `CloudFormationStackARN` or `TagFilters` .
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}


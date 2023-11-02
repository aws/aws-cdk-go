package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the amount of read and write operations supported by a DynamoDB table.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-2"),
//   	},
//   })
//
//   globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Billing: dynamodb.Billing_Provisioned(&ThroughputProps{
//   		ReadCapacity: dynamodb.Capacity_Fixed(jsii.Number(10)),
//   		WriteCapacity: dynamodb.Capacity_Autoscaled(&AutoscaledCapacityOptions{
//   			MaxCapacity: jsii.Number(15),
//   		}),
//   	}),
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   			ReadCapacity: dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
//   				MaxCapacity: jsii.Number(20),
//   				TargetUtilizationPercent: jsii.Number(50),
//   			}),
//   		},
//   	},
//   })
//
type Capacity interface {
	Mode() CapacityMode
}

// The jsii proxy struct for Capacity
type jsiiProxy_Capacity struct {
	_ byte // padding
}

func (j *jsiiProxy_Capacity) Mode() CapacityMode {
	var returns CapacityMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// Dynamically adjusts provisioned throughput capacity on your behalf in response to actual traffic patterns.
func Capacity_Autoscaled(options *AutoscaledCapacityOptions) Capacity {
	_init_.Initialize()

	if err := validateCapacity_AutoscaledParameters(options); err != nil {
		panic(err)
	}
	var returns Capacity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.Capacity",
		"autoscaled",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Provisioned throughput capacity is configured with fixed capacity units.
//
// Note: You cannot configure write capacity using fixed capacity mode.
func Capacity_Fixed(iops *float64) Capacity {
	_init_.Initialize()

	if err := validateCapacity_FixedParameters(iops); err != nil {
		panic(err)
	}
	var returns Capacity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.Capacity",
		"fixed",
		[]interface{}{iops},
		&returns,
	)

	return returns
}


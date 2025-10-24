package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents how capacity is managed and how you are charged for read and write throughput for a DynamoDB table.
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
//   	Replicas: []ReplicaTableProps{
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   			ReadCapacity: dynamodb.Capacity_*Autoscaled(&AutoscaledCapacityOptions{
//   				MaxCapacity: jsii.Number(20),
//   				TargetUtilizationPercent: jsii.Number(50),
//   			}),
//   		},
//   	},
//   })
//
type Billing interface {
	Mode() BillingMode
}

// The jsii proxy struct for Billing
type jsiiProxy_Billing struct {
	_ byte // padding
}

func (j *jsiiProxy_Billing) Mode() BillingMode {
	var returns BillingMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// Flexible billing option capable of serving requests without capacity planning.
//
// Note: Billing mode will be PAY_PER_REQUEST.
func Billing_OnDemand(props *MaxThroughputProps) Billing {
	_init_.Initialize()

	if err := validateBilling_OnDemandParameters(props); err != nil {
		panic(err)
	}
	var returns Billing

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.Billing",
		"onDemand",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Specify the number of reads and writes per second that you need for your application.
func Billing_Provisioned(props *ThroughputProps) Billing {
	_init_.Initialize()

	if err := validateBilling_ProvisionedParameters(props); err != nil {
		panic(err)
	}
	var returns Billing

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.Billing",
		"provisioned",
		[]interface{}{props},
		&returns,
	)

	return returns
}


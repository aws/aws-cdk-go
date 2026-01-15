package previewawss3events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Bucket aws.s3@ObjectCreated event.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//   var fn Function
//
//
//   // Works with L2 constructs
//   bucket := s3.NewBucket(scope, jsii.String("Bucket"))
//   bucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(bucket)
//
//   events.NewRule(scope, jsii.String("Rule"), &RuleProps{
//   	EventPattern: bucketEvents.ObjectCreatedPattern(&ObjectCreatedProps{
//   		Object: &ObjectType{
//   			Key: events.Match_Wildcard(jsii.String("uploads/*")),
//   		},
//   	}),
//   	Targets: []IRuleTarget{
//   		targets.NewLambdaFunction(fn),
//   	},
//   })
//
//   // Also works with L1 constructs
//   cfnBucket := s3.NewCfnBucket(scope, jsii.String("CfnBucket"))
//   cfnBucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(cfnBucket)
//
//   events.NewCfnRule(scope, jsii.String("CfnRule"), &CfnRuleProps{
//   	State: jsii.String("ENABLED"),
//   	EventPattern: cfnBucketEvents.*ObjectCreatedPattern(&ObjectCreatedProps{
//   		Object: &ObjectType{
//   			Key: events.Match_*Wildcard(jsii.String("uploads/*")),
//   		},
//   	}),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Arn: fn.FunctionArn,
//   			Id: jsii.String("L1"),
//   		},
//   	},
//   })
//
// Experimental.
type BucketEvents_ObjectCreated_ObjectCreatedProps struct {
	// bucket property.
	//
	// Specify an array of string values to match this event if the actual value of bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bucket *BucketEvents_ObjectCreated_Bucket `field:"optional" json:"bucket" yaml:"bucket"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// object property.
	//
	// Specify an array of string values to match this event if the actual value of object is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Object *BucketEvents_ObjectCreated_ObjectType `field:"optional" json:"object" yaml:"object"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
	// requester property.
	//
	// Specify an array of string values to match this event if the actual value of requester is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Requester *[]*string `field:"optional" json:"requester" yaml:"requester"`
	// request-id property.
	//
	// Specify an array of string values to match this event if the actual value of request-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// source-ip-address property.
	//
	// Specify an array of string values to match this event if the actual value of source-ip-address is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}


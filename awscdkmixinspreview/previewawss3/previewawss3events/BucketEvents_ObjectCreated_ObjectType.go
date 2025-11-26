package previewawss3events


// Type definition for Object.
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
//   			Key: []*string{
//   				jsii.String("uploads/*"),
//   			},
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
//   			Key: []*string{
//   				jsii.String("uploads/*"),
//   			},
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
type BucketEvents_ObjectCreated_ObjectType struct {
	// etag property.
	//
	// Specify an array of string values to match this event if the actual value of etag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Etag *[]*string `field:"optional" json:"etag" yaml:"etag"`
	// key property.
	//
	// Specify an array of string values to match this event if the actual value of key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
	// sequencer property.
	//
	// Specify an array of string values to match this event if the actual value of sequencer is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Sequencer *[]*string `field:"optional" json:"sequencer" yaml:"sequencer"`
	// size property.
	//
	// Specify an array of string values to match this event if the actual value of size is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Size *[]*string `field:"optional" json:"size" yaml:"size"`
	// version-id property.
	//
	// Specify an array of string values to match this event if the actual value of version-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VersionId *[]*string `field:"optional" json:"versionId" yaml:"versionId"`
}


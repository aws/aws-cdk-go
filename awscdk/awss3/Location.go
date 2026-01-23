package awss3


// An interface that represents the location of a specific object in an S3 Bucket.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("memoryBucket"), &BucketProps{
//   	BucketName: jsii.String("test-memory"),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//
//   topic := sns.NewTopic(this, jsii.String("topic"))
//
//   // Create a custom semantic memory strategy
//   selfManagedStrategy := agentcore.MemoryStrategy_UsingSelfManaged(&SelfManagedStrategyProps{
//   	Name: jsii.String("selfManagedStrategy"),
//   	Description: jsii.String("self managed memory strategy"),
//   	HistoricalContextWindowSize: jsii.Number(5),
//   	InvocationConfiguration: &InvocationConfiguration{
//   		Topic: topic,
//   		S3Location: &Location{
//   			BucketName: bucket.bucketName,
//   			ObjectKey: jsii.String("memory/"),
//   		},
//   	},
//   	TriggerConditions: &TriggerConditions{
//   		MessageBasedTrigger: jsii.Number(1),
//   		TimeBasedTrigger: cdk.Duration_Seconds(jsii.Number(10)),
//   		TokenBasedTrigger: jsii.Number(100),
//   	},
//   })
//
//   // Create memory with custom strategy
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my-custom-memory"),
//   	Description: jsii.String("Memory with custom strategy"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		selfManagedStrategy,
//   	},
//   })
//
type Location struct {
	// The name of the S3 Bucket the object is in.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The path inside the Bucket where the object is located at.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 object version.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}


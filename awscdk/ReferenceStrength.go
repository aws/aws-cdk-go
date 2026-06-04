package awscdk


// Controls how cross-stack references to a resource are resolved.
//
// Example:
//   var bucket Bucket
//   var consumer Stack
//
//
//   awscdk.NewCfnOutput(consumer, jsii.String("BucketArn"), &CfnOutputProps{
//   	Value: awscdk.*stack_ConsumeReference(bucket.bucketArn, awscdk.ReferenceStrength_WEAK),
//   })
//
type ReferenceStrength string

const (
	// Strong reference: uses CloudFormation Export/Import (same region) or ExportWriter/ExportReader custom resources (cross-region).
	//
	// The producing stack cannot be deleted while consumers exist.
	ReferenceStrength_STRONG ReferenceStrength = "STRONG"
	// Weak reference: uses Fn::GetStackOutput to read an output directly from the producing stack.
	//
	// The producing stack or resource can be deleted independently of consumers.
	// This will cause infrastructure in consuming stacks to temporarily reference a nonexistant
	// resource until the consumers are updated as well, causing any accesses in that time
	// frame to fail.
	//
	// Strong references prevent this.
	ReferenceStrength_WEAK ReferenceStrength = "WEAK"
	// Both strong and weak mechanisms are created (transitional state).
	//
	// Use this when migrating from strong to weak. The producer keeps the
	// strong-side artifacts and also adds a plain Output. The consumer
	// switches to Fn::GetStackOutput.
	ReferenceStrength_BOTH ReferenceStrength = "BOTH"
)


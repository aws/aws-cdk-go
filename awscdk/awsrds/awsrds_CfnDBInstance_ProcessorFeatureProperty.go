package awsrds


// The `ProcessorFeature` property type specifies the processor features of a DB instance class status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorFeatureProperty := &processorFeatureProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnDBInstance_ProcessorFeatureProperty struct {
	// The name of the processor feature.
	//
	// Valid names are `coreCount` and `threadsPerCore` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of a processor feature name.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


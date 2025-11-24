package mixinsawsbatch


// Specifies the order of a service environment for a job queue.
//
// This determines the priority order when multiple service environments are associated with the same job queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceEnvironmentOrderProperty := &ServiceEnvironmentOrderProperty{
//   	Order: jsii.Number(123),
//   	ServiceEnvironment: jsii.String("serviceEnvironment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-serviceenvironmentorder.html
//
type CfnJobQueuePropsMixin_ServiceEnvironmentOrderProperty struct {
	// The order of the service environment.
	//
	// Job queues with a higher priority are evaluated first when associated with the same service environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-serviceenvironmentorder.html#cfn-batch-jobqueue-serviceenvironmentorder-order
	//
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// The name or ARN of the service environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-serviceenvironmentorder.html#cfn-batch-jobqueue-serviceenvironmentorder-serviceenvironment
	//
	ServiceEnvironment *string `field:"optional" json:"serviceEnvironment" yaml:"serviceEnvironment"`
}


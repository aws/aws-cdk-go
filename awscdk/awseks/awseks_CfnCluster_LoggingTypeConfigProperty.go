package awseks


// The enabled logging type.
//
// For a list of the valid logging types, see the [`types` property of `LogSetup`](https://docs.aws.amazon.com/eks/latest/APIReference/API_LogSetup.html#AmazonEKS-Type-LogSetup-types) in the *Amazon EKS API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingTypeConfigProperty := &loggingTypeConfigProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnCluster_LoggingTypeConfigProperty struct {
	// The name of the log type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterQueueProperty := &DeadLetterQueueProperty{
//   	TargetArn: jsii.String("targetArn"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deadletterqueue.html
//
type CfnFunction_DeadLetterQueueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deadletterqueue.html#cfn-serverless-function-deadletterqueue-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-deadletterqueue.html#cfn-serverless-function-deadletterqueue-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}


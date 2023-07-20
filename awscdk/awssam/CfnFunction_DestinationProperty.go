package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &DestinationProperty{
//   	Destination: jsii.String("destination"),
//
//   	// the properties below are optional
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-destination.html
//
type CfnFunction_DestinationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-destination.html#cfn-serverless-function-destination-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-destination.html#cfn-serverless-function-destination-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


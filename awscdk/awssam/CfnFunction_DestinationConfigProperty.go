package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigProperty := &DestinationConfigProperty{
//   	OnFailure: &DestinationProperty{
//   		Destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-destinationconfig.html
//
type CfnFunction_DestinationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-destinationconfig.html#cfn-serverless-function-destinationconfig-onfailure
	//
	OnFailure interface{} `field:"required" json:"onFailure" yaml:"onFailure"`
}


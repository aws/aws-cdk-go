package awsdatazone


// MLflow Properties Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mlflowPropertiesInputProperty := &MlflowPropertiesInputProperty{
//   	TrackingServerArn: jsii.String("trackingServerArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-mlflowpropertiesinput.html
//
type CfnConnection_MlflowPropertiesInputProperty struct {
	// The ARN of the MLflow tracking server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-mlflowpropertiesinput.html#cfn-datazone-connection-mlflowpropertiesinput-trackingserverarn
	//
	TrackingServerArn *string `field:"optional" json:"trackingServerArn" yaml:"trackingServerArn"`
}


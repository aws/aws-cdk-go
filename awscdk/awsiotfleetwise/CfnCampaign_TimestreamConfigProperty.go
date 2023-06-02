package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestreamConfigProperty := &TimestreamConfigProperty{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	TimestreamTableArn: jsii.String("timestreamTableArn"),
//   }
//
type CfnCampaign_TimestreamConfigProperty struct {
	// `CfnCampaign.TimestreamConfigProperty.ExecutionRoleArn`.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// `CfnCampaign.TimestreamConfigProperty.TimestreamTableArn`.
	TimestreamTableArn *string `field:"required" json:"timestreamTableArn" yaml:"timestreamTableArn"`
}


package awsiotsitewise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sessionConfigProperty := &SessionConfigProperty{
//   	SessionEndTime: jsii.String("sessionEndTime"),
//   	SessionStartTime: jsii.String("sessionStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-sessionconfig.html
//
type CfnDatasetPropsMixin_SessionConfigProperty struct {
	// The end time of the session as an ISO 8601 UTC instant, for example 2024-12-31T23:59:59Z.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-sessionconfig.html#cfn-iotsitewise-dataset-sessionconfig-sessionendtime
	//
	SessionEndTime *string `field:"optional" json:"sessionEndTime" yaml:"sessionEndTime"`
	// The start time of the session as an ISO 8601 UTC instant, for example 2024-01-01T00:00:00Z.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-sessionconfig.html#cfn-iotsitewise-dataset-sessionconfig-sessionstarttime
	//
	SessionStartTime *string `field:"optional" json:"sessionStartTime" yaml:"sessionStartTime"`
}


package awsiotsitewise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetConfigProperty := &DatasetConfigProperty{
//   	Session: &SessionConfigProperty{
//   		SessionEndTime: jsii.String("sessionEndTime"),
//   		SessionStartTime: jsii.String("sessionStartTime"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetconfig.html
//
type CfnDataset_DatasetConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetconfig.html#cfn-iotsitewise-dataset-datasetconfig-session
	//
	Session interface{} `field:"optional" json:"session" yaml:"session"`
}


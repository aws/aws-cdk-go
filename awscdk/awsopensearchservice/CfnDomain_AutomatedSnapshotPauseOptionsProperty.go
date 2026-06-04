package awsopensearchservice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automatedSnapshotPauseOptionsProperty := &AutomatedSnapshotPauseOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-automatedsnapshotpauseoptions.html
//
type CfnDomain_AutomatedSnapshotPauseOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-automatedsnapshotpauseoptions.html#cfn-opensearchservice-domain-automatedsnapshotpauseoptions-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-automatedsnapshotpauseoptions.html#cfn-opensearchservice-domain-automatedsnapshotpauseoptions-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-automatedsnapshotpauseoptions.html#cfn-opensearchservice-domain-automatedsnapshotpauseoptions-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}


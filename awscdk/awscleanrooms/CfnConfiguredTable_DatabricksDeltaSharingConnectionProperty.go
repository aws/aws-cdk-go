package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databricksDeltaSharingConnectionProperty := &DatabricksDeltaSharingConnectionProperty{
//   	DeltaShareEndpoint: jsii.String("deltaShareEndpoint"),
//   	ShareName: jsii.String("shareName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksdeltasharingconnection.html
//
type CfnConfiguredTable_DatabricksDeltaSharingConnectionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksdeltasharingconnection.html#cfn-cleanrooms-configuredtable-databricksdeltasharingconnection-deltashareendpoint
	//
	DeltaShareEndpoint *string `field:"required" json:"deltaShareEndpoint" yaml:"deltaShareEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksdeltasharingconnection.html#cfn-cleanrooms-configuredtable-databricksdeltasharingconnection-sharename
	//
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
}


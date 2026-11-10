package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   databricksDeltaSharingConnectionProperty := &DatabricksDeltaSharingConnectionProperty{
//   	DeltaShareEndpoint: jsii.String("deltaShareEndpoint"),
//   	ShareName: jsii.String("shareName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksdeltasharingconnection.html
//
type CfnConfiguredTablePropsMixin_DatabricksDeltaSharingConnectionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksdeltasharingconnection.html#cfn-cleanrooms-configuredtable-databricksdeltasharingconnection-deltashareendpoint
	//
	DeltaShareEndpoint *string `field:"optional" json:"deltaShareEndpoint" yaml:"deltaShareEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksdeltasharingconnection.html#cfn-cleanrooms-configuredtable-databricksdeltasharingconnection-sharename
	//
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
}


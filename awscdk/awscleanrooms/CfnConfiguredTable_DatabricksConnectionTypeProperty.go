package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databricksConnectionTypeProperty := &DatabricksConnectionTypeProperty{
//   	DeltaSharing: &DatabricksDeltaSharingConnectionProperty{
//   		DeltaShareEndpoint: jsii.String("deltaShareEndpoint"),
//   		ShareName: jsii.String("shareName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksconnectiontype.html
//
type CfnConfiguredTable_DatabricksConnectionTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksconnectiontype.html#cfn-cleanrooms-configuredtable-databricksconnectiontype-deltasharing
	//
	DeltaSharing interface{} `field:"required" json:"deltaSharing" yaml:"deltaSharing"`
}


package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnConfiguredTablePropsMixin_DatabricksConnectionTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-databricksconnectiontype.html#cfn-cleanrooms-configuredtable-databricksconnectiontype-deltasharing
	//
	DeltaSharing interface{} `field:"optional" json:"deltaSharing" yaml:"deltaSharing"`
}


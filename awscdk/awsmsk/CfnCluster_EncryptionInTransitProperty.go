package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionInTransitProperty := &EncryptionInTransitProperty{
//   	ClientBroker: jsii.String("clientBroker"),
//   	InCluster: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionintransit.html
//
type CfnCluster_EncryptionInTransitProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionintransit.html#cfn-msk-cluster-encryptionintransit-clientbroker
	//
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionintransit.html#cfn-msk-cluster-encryptionintransit-incluster
	//
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}


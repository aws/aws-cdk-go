package mixinsawsmsk


// The settings for encrypting data in transit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionInTransitProperty := &EncryptionInTransitProperty{
//   	ClientBroker: jsii.String("clientBroker"),
//   	InCluster: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionintransit.html
//
type CfnClusterPropsMixin_EncryptionInTransitProperty struct {
	// Indicates the encryption setting for data in transit between clients and brokers.
	//
	// You must set it to one of the following values.
	//
	// - `TLS` : Indicates that client-broker communication is enabled with TLS only.
	// - `TLS_PLAINTEXT` : Indicates that client-broker communication is enabled for both TLS-encrypted, as well as plaintext data.
	// - `PLAINTEXT` : Indicates that client-broker communication is enabled in plaintext only.
	//
	// The default value is `TLS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionintransit.html#cfn-msk-cluster-encryptionintransit-clientbroker
	//
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// When set to true, it indicates that data communication among the broker nodes of the cluster is encrypted.
	//
	// When set to false, the communication happens in plaintext.
	//
	// The default value is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionintransit.html#cfn-msk-cluster-encryptionintransit-incluster
	//
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}


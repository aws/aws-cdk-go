package previewawsmskmixins


// Configuration for specifying replicated topic names will be the same as their corresponding upstream topics or prefixed with source cluster alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicationTopicNameConfigurationProperty := &ReplicationTopicNameConfigurationProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationtopicnameconfiguration.html
//
type CfnReplicatorPropsMixin_ReplicationTopicNameConfigurationProperty struct {
	// The type of replication topic name configuration, identical to upstream topic name or prefixed with source cluster alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationtopicnameconfiguration.html#cfn-msk-replicator-replicationtopicnameconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


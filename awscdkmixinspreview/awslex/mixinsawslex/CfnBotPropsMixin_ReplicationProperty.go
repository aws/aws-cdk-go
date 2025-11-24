package mixinsawslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicationProperty := &ReplicationProperty{
//   	ReplicaRegions: []*string{
//   		jsii.String("replicaRegions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-replication.html
//
type CfnBotPropsMixin_ReplicationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-replication.html#cfn-lex-bot-replication-replicaregions
	//
	ReplicaRegions *[]*string `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
}


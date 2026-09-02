package awsglue


// Specifies the connections used by the session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectionsListProperty := &ConnectionsListProperty{
//   	Connections: []*string{
//   		jsii.String("connections"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-session-connectionslist.html
//
type CfnSessionPropsMixin_ConnectionsListProperty struct {
	// A list of connection names used by the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-session-connectionslist.html#cfn-glue-session-connectionslist-connections
	//
	Connections *[]*string `field:"optional" json:"connections" yaml:"connections"`
}


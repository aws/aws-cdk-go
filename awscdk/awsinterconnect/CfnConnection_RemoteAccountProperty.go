package awsinterconnect


// The remote account identifier for the connection.
//
// Required when creating a connection through AWS. Replaces RemoteOwnerAccount.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remoteAccountProperty := &RemoteAccountProperty{
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-remoteaccount.html
//
type CfnConnection_RemoteAccountProperty struct {
	// The identifier of the remote account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-interconnect-connection-remoteaccount.html#cfn-interconnect-connection-remoteaccount-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}


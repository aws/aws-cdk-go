package awsglue


// Specifies the connections used by a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionsListProperty := &connectionsListProperty{
//   	connections: []*string{
//   		jsii.String("connections"),
//   	},
//   }
//
type CfnJob_ConnectionsListProperty struct {
	// A list of connections used by the job.
	Connections *[]*string `field:"optional" json:"connections" yaml:"connections"`
}


package awsglue


// Specifies the connections used by a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionsListProperty := &ConnectionsListProperty{
//   	Connections: []*string{
//   		jsii.String("connections"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-connectionslist.html
//
type CfnJob_ConnectionsListProperty struct {
	// A list of connections used by the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-connectionslist.html#cfn-glue-job-connectionslist-connections
	//
	Connections *[]*string `field:"optional" json:"connections" yaml:"connections"`
}


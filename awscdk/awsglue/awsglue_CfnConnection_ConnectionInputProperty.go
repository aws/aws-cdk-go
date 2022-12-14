package awsglue


// A structure that is used to specify a connection to create or update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var connectionProperties interface{}
//
//   connectionInputProperty := &connectionInputProperty{
//   	connectionType: jsii.String("connectionType"),
//
//   	// the properties below are optional
//   	connectionProperties: connectionProperties,
//   	description: jsii.String("description"),
//   	matchCriteria: []*string{
//   		jsii.String("matchCriteria"),
//   	},
//   	name: jsii.String("name"),
//   	physicalConnectionRequirements: &physicalConnectionRequirementsProperty{
//   		availabilityZone: jsii.String("availabilityZone"),
//   		securityGroupIdList: []*string{
//   			jsii.String("securityGroupIdList"),
//   		},
//   		subnetId: jsii.String("subnetId"),
//   	},
//   }
//
type CfnConnection_ConnectionInputProperty struct {
	// The type of the connection. Currently, these types are supported:.
	//
	// - `JDBC` - Designates a connection to a database through Java Database Connectivity (JDBC).
	// - `KAFKA` - Designates a connection to an Apache Kafka streaming platform.
	// - `MONGODB` - Designates a connection to a MongoDB document database.
	// - `NETWORK` - Designates a network connection to a data source within an Amazon Virtual Private Cloud environment (Amazon VPC).
	//
	// SFTP is not supported.
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// These key-value pairs define parameters for the connection.
	ConnectionProperties interface{} `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// The description of the connection.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	MatchCriteria *[]*string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// The name of the connection.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map of physical connection requirements, such as virtual private cloud (VPC) and `SecurityGroup` , that are needed to successfully make this connection.
	PhysicalConnectionRequirements interface{} `field:"optional" json:"physicalConnectionRequirements" yaml:"physicalConnectionRequirements"`
}


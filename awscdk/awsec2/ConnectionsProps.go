package awsec2


// Properties to intialize a new Connections object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var peer iPeer
//   var port port
//   var securityGroup securityGroup
//
//   connectionsProps := &ConnectionsProps{
//   	DefaultPort: port,
//   	Peer: peer,
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type ConnectionsProps struct {
	// Default port range for initiating connections to and from this object.
	// Default: - No default port.
	//
	DefaultPort Port `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// Class that represents the rule by which others can connect to this connectable.
	//
	// This object is required, but will be derived from securityGroup if that is passed.
	// Default: Derived from securityGroup if set.
	//
	Peer IPeer `field:"optional" json:"peer" yaml:"peer"`
	// What securityGroup(s) this object is managing connections for.
	// Default: No security groups.
	//
	SecurityGroups *[]ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}


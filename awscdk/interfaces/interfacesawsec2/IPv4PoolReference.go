package interfacesawsec2


// A reference to a IPv4Pool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPv4PoolReference := map[string]*string{
//   	"iPv4PoolArn": jsii.String("iPv4PoolArn"),
//   }
//
type IPv4PoolReference struct {
	// The Arn of the IPv4Pool resource.
	IPv4PoolArn *string `field:"required" json:"iPv4PoolArn" yaml:"iPv4PoolArn"`
}


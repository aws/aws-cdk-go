package interfacesawswafregional


// A reference to a IPSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetReference := map[string]*string{
//   	"ipSetId": jsii.String("ipSetId"),
//   }
//
type IPSetReference struct {
	// The Id of the IPSet resource.
	IpSetId *string `field:"required" json:"ipSetId" yaml:"ipSetId"`
}


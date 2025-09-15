package awswafv2


// A reference to a IPSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetReference := map[string]*string{
//   	"ipSetArn": jsii.String("ipSetArn"),
//   	"ipSetId": jsii.String("ipSetId"),
//   	"ipSetName": jsii.String("ipSetName"),
//   	"scope": jsii.String("scope"),
//   }
//
type IPSetReference struct {
	// The ARN of the IPSet resource.
	IpSetArn *string `field:"required" json:"ipSetArn" yaml:"ipSetArn"`
	// The Id of the IPSet resource.
	IpSetId *string `field:"required" json:"ipSetId" yaml:"ipSetId"`
	// The Name of the IPSet resource.
	IpSetName *string `field:"required" json:"ipSetName" yaml:"ipSetName"`
	// The Scope of the IPSet resource.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}


package interfacesawsdirectoryservice


// A reference to a MicrosoftAD resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microsoftADReference := &MicrosoftADReference{
//   	MicrosoftAdId: jsii.String("microsoftAdId"),
//   }
//
type MicrosoftADReference struct {
	// The Id of the MicrosoftAD resource.
	MicrosoftAdId *string `field:"required" json:"microsoftAdId" yaml:"microsoftAdId"`
}


package awssmsvoice


// A reference to a SenderId resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   senderIdReference := &SenderIdReference{
//   	IsoCountryCode: jsii.String("isoCountryCode"),
//   	SenderId: jsii.String("senderId"),
//   	SenderIdArn: jsii.String("senderIdArn"),
//   }
//
type SenderIdReference struct {
	// The IsoCountryCode of the SenderId resource.
	IsoCountryCode *string `field:"required" json:"isoCountryCode" yaml:"isoCountryCode"`
	// The SenderId of the SenderId resource.
	SenderId *string `field:"required" json:"senderId" yaml:"senderId"`
	// The ARN of the SenderId resource.
	SenderIdArn *string `field:"required" json:"senderIdArn" yaml:"senderIdArn"`
}


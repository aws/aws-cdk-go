package awskinesisanalytics


// A reference to a ApplicationOutput resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationOutputV2Reference := &ApplicationOutputV2Reference{
//   	ApplicationOutputId: jsii.String("applicationOutputId"),
//   }
//
type ApplicationOutputV2Reference struct {
	// The Id of the ApplicationOutput resource.
	ApplicationOutputId *string `field:"required" json:"applicationOutputId" yaml:"applicationOutputId"`
}


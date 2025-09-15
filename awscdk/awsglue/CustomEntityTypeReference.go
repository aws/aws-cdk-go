package awsglue


// A reference to a CustomEntityType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customEntityTypeReference := &CustomEntityTypeReference{
//   	CustomEntityTypeId: jsii.String("customEntityTypeId"),
//   }
//
type CustomEntityTypeReference struct {
	// The Id of the CustomEntityType resource.
	CustomEntityTypeId *string `field:"required" json:"customEntityTypeId" yaml:"customEntityTypeId"`
}


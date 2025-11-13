package interfacesawsbedrock


// A reference to a Blueprint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blueprintReference := &BlueprintReference{
//   	BlueprintArn: jsii.String("blueprintArn"),
//   }
//
type BlueprintReference struct {
	// The BlueprintArn of the Blueprint resource.
	BlueprintArn *string `field:"required" json:"blueprintArn" yaml:"blueprintArn"`
}


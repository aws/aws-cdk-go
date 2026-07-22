package interfacesawscontrolcatalog


// A reference to a CommonControl resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commonControlReference := &CommonControlReference{
//   	CommonControlArn: jsii.String("commonControlArn"),
//   	CommonControlId: jsii.String("commonControlId"),
//   }
//
type CommonControlReference struct {
	// The ARN of the CommonControl resource.
	CommonControlArn *string `field:"required" json:"commonControlArn" yaml:"commonControlArn"`
	// The CommonControlId of the CommonControl resource.
	CommonControlId *string `field:"required" json:"commonControlId" yaml:"commonControlId"`
}


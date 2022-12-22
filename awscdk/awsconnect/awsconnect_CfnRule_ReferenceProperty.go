package awsconnect


// Information about the reference when the `referenceType` is `URL` .
//
// Otherwise, null. (Supports variable injection in the `Value` field.)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceProperty := &referenceProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnRule_ReferenceProperty struct {
	// The type of the reference. `DATE` must be of type Epoch timestamp.
	//
	// *Allowed values* : `URL` | `ATTACHMENT` | `NUMBER` | `STRING` | `DATE` | `EMAIL`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A valid value for the reference.
	//
	// For example, for a URL reference, a formatted URL that is displayed to an agent in the Contact Control Panel (CCP).
	Value *string `field:"required" json:"value" yaml:"value"`
}


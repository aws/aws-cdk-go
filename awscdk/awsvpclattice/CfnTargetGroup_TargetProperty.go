package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetProperty := &TargetProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Port: jsii.Number(123),
//   }
//
type CfnTargetGroup_TargetProperty struct {
	// `CfnTargetGroup.TargetProperty.Id`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// `CfnTargetGroup.TargetProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}


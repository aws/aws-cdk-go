package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hooksProperty := &hooksProperty{
//   	postTraffic: jsii.String("postTraffic"),
//   	preTraffic: jsii.String("preTraffic"),
//   }
//
type CfnFunction_HooksProperty struct {
	// `CfnFunction.HooksProperty.PostTraffic`.
	PostTraffic *string `field:"optional" json:"postTraffic" yaml:"postTraffic"`
	// `CfnFunction.HooksProperty.PreTraffic`.
	PreTraffic *string `field:"optional" json:"preTraffic" yaml:"preTraffic"`
}


package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInputProperty := &modelInputProperty{
//   	dataInputConfig: jsii.String("dataInputConfig"),
//   }
//
type CfnModelPackage_ModelInputProperty struct {
	// `CfnModelPackage.ModelInputProperty.DataInputConfig`.
	DataInputConfig *string `field:"required" json:"dataInputConfig" yaml:"dataInputConfig"`
}


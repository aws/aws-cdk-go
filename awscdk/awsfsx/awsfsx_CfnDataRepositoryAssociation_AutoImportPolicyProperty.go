package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoImportPolicyProperty := &autoImportPolicyProperty{
//   	events: []*string{
//   		jsii.String("events"),
//   	},
//   }
//
type CfnDataRepositoryAssociation_AutoImportPolicyProperty struct {
	// `CfnDataRepositoryAssociation.AutoImportPolicyProperty.Events`.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
}


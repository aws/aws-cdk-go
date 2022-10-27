package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &s3Property{
//   	autoExportPolicy: &autoExportPolicyProperty{
//   		events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   	autoImportPolicy: &autoImportPolicyProperty{
//   		events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   }
//
type CfnDataRepositoryAssociation_S3Property struct {
	// `CfnDataRepositoryAssociation.S3Property.AutoExportPolicy`.
	AutoExportPolicy interface{} `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// `CfnDataRepositoryAssociation.S3Property.AutoImportPolicy`.
	AutoImportPolicy interface{} `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}


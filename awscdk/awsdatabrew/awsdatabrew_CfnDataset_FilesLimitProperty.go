package awsdatabrew


// Represents a limit imposed on number of Amazon S3 files that should be selected for a dataset from a connected Amazon S3 path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filesLimitProperty := &filesLimitProperty{
//   	maxFiles: jsii.Number(123),
//
//   	// the properties below are optional
//   	order: jsii.String("order"),
//   	orderedBy: jsii.String("orderedBy"),
//   }
//
type CfnDataset_FilesLimitProperty struct {
	// The number of Amazon S3 files to select.
	MaxFiles *float64 `field:"required" json:"maxFiles" yaml:"maxFiles"`
	// A criteria to use for Amazon S3 files sorting before their selection.
	//
	// By default uses DESCENDING order, i.e. most recent files are selected first. Anotherpossible value is ASCENDING.
	Order *string `field:"optional" json:"order" yaml:"order"`
	// A criteria to use for Amazon S3 files sorting before their selection.
	//
	// By default uses LAST_MODIFIED_DATE as a sorting criteria. Currently it's the only allowed value.
	OrderedBy *string `field:"optional" json:"orderedBy" yaml:"orderedBy"`
}


package awsquicksight


// A transform operation that projects columns.
//
// Operations that come after a projection can only refer to projected columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectOperationProperty := &projectOperationProperty{
//   	projectedColumns: []*string{
//   		jsii.String("projectedColumns"),
//   	},
//   }
//
type CfnDataSet_ProjectOperationProperty struct {
	// Projected columns.
	ProjectedColumns *[]*string `field:"required" json:"projectedColumns" yaml:"projectedColumns"`
}


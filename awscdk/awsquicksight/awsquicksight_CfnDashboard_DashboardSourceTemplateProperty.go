package awsquicksight


// Dashboard source template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardSourceTemplateProperty := &dashboardSourceTemplateProperty{
//   	arn: jsii.String("arn"),
//   	dataSetReferences: []interface{}{
//   		&dataSetReferenceProperty{
//   			dataSetArn: jsii.String("dataSetArn"),
//   			dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   		},
//   	},
//   }
//
type CfnDashboard_DashboardSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Dataset references.
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}


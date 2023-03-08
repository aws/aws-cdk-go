package awssagemaker


// The configuration for a baseline model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelBiasBaselineConfigProperty := &modelBiasBaselineConfigProperty{
//   	baseliningJobName: jsii.String("baseliningJobName"),
//   	constraintsResource: &constraintsResourceProperty{
//   		s3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasBaselineConfigProperty struct {
	// The name of the baseline model bias job.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The constraints resource for a monitoring job.
	ConstraintsResource interface{} `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
}


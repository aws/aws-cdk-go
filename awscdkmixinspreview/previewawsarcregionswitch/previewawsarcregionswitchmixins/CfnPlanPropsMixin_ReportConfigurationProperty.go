package previewawsarcregionswitchmixins


// Configuration for automatic report generation for plan executions.
//
// When configured, Region switch automatically generates a report after each plan execution that includes execution events, plan configuration, and CloudWatch alarm states.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   reportConfigurationProperty := &ReportConfigurationProperty{
//   	ReportOutput: []interface{}{
//   		&ReportOutputConfigurationProperty{
//   			S3Configuration: &S3ReportOutputConfigurationProperty{
//   				BucketOwner: jsii.String("bucketOwner"),
//   				BucketPath: jsii.String("bucketPath"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-reportconfiguration.html
//
type CfnPlanPropsMixin_ReportConfigurationProperty struct {
	// The output configuration for the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-reportconfiguration.html#cfn-arcregionswitch-plan-reportconfiguration-reportoutput
	//
	ReportOutput interface{} `field:"optional" json:"reportOutput" yaml:"reportOutput"`
}


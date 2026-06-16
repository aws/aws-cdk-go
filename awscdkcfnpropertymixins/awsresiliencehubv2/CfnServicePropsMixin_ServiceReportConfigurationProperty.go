package awsresiliencehubv2


// Configuration for automatic report generation on a Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   serviceReportConfigurationProperty := &ServiceReportConfigurationProperty{
//   	ReportOutput: []interface{}{
//   		&ReportOutputConfigurationProperty{
//   			S3: &S3ReportOutputConfigurationProperty{
//   				BucketOwner: jsii.String("bucketOwner"),
//   				BucketPath: jsii.String("bucketPath"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-servicereportconfiguration.html
//
type CfnServicePropsMixin_ServiceReportConfigurationProperty struct {
	// Output destinations for generated reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-servicereportconfiguration.html#cfn-resiliencehubv2-service-servicereportconfiguration-reportoutput
	//
	ReportOutput interface{} `field:"optional" json:"reportOutput" yaml:"reportOutput"`
}


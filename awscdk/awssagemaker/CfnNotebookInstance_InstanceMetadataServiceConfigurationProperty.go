package awssagemaker


// Information on the IMDS configuration of the notebook instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceMetadataServiceConfigurationProperty := &InstanceMetadataServiceConfigurationProperty{
//   	MinimumInstanceMetadataServiceVersion: jsii.String("minimumInstanceMetadataServiceVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-notebookinstance-instancemetadataserviceconfiguration.html
//
type CfnNotebookInstance_InstanceMetadataServiceConfigurationProperty struct {
	// Indicates the minimum IMDS version that the notebook instance supports.
	//
	// When passed as part of `CreateNotebookInstance` , if no value is selected, then it defaults to IMDSv1. This means that both IMDSv1 and IMDSv2 are supported. If passed as part of `UpdateNotebookInstance` , there is no default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-notebookinstance-instancemetadataserviceconfiguration.html#cfn-sagemaker-notebookinstance-instancemetadataserviceconfiguration-minimuminstancemetadataserviceversion
	//
	MinimumInstanceMetadataServiceVersion *string `field:"required" json:"minimumInstanceMetadataServiceVersion" yaml:"minimumInstanceMetadataServiceVersion"`
}


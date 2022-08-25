package awssagemaker


// The input for the data quality monitoring job.
//
// Currently endpoints are supported for input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityJobInputProperty := &dataQualityJobInputProperty{
//   	endpointInput: &endpointInputProperty{
//   		endpointName: jsii.String("endpointName"),
//   		localPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		s3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnDataQualityJobDefinition_DataQualityJobInputProperty struct {
	// `CfnDataQualityJobDefinition.DataQualityJobInputProperty.EndpointInput`.
	EndpointInput interface{} `field:"required" json:"endpointInput" yaml:"endpointInput"`
}


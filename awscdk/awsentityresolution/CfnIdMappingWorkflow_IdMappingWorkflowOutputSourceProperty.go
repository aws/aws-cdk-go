package awsentityresolution


// A list of `IdMappingWorkflowOutputSource` objects, each of which contains fields `outputS3Path` and `KMSArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingWorkflowOutputSourceProperty := &IdMappingWorkflowOutputSourceProperty{
//   	OutputS3Path: jsii.String("outputS3Path"),
//
//   	// the properties below are optional
//   	KmsArn: jsii.String("kmsArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowoutputsource.html
//
type CfnIdMappingWorkflow_IdMappingWorkflowOutputSourceProperty struct {
	// The S3 path to which AWS Entity Resolution will write the output table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowoutputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowoutputsource-outputs3path
	//
	OutputS3Path *string `field:"required" json:"outputS3Path" yaml:"outputS3Path"`
	// Customer AWS KMS ARN for encryption at rest.
	//
	// If not provided, system will use an AWS Entity Resolution managed KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowoutputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowoutputsource-kmsarn
	//
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}


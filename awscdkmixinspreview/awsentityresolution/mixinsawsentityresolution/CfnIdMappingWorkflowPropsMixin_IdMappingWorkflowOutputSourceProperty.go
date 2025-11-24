package mixinsawsentityresolution


// A list of `IdMappingWorkflowOutputSource` objects, each of which contains fields `outputS3Path` and `KMSArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idMappingWorkflowOutputSourceProperty := &IdMappingWorkflowOutputSourceProperty{
//   	KmsArn: jsii.String("kmsArn"),
//   	OutputS3Path: jsii.String("outputS3Path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowoutputsource.html
//
type CfnIdMappingWorkflowPropsMixin_IdMappingWorkflowOutputSourceProperty struct {
	// Customer AWS  ARN for encryption at rest.
	//
	// If not provided, system will use an AWS Entity Resolution managed KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowoutputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowoutputsource-kmsarn
	//
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// The S3 path to which AWS Entity Resolution will write the output table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-idmappingworkflowoutputsource.html#cfn-entityresolution-idmappingworkflow-idmappingworkflowoutputsource-outputs3path
	//
	OutputS3Path *string `field:"optional" json:"outputS3Path" yaml:"outputS3Path"`
}


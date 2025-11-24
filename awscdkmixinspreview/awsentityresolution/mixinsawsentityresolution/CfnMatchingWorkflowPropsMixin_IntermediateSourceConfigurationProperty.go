package mixinsawsentityresolution


// The Amazon S3 location that temporarily stores your data while it processes.
//
// Your information won't be saved permanently.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   intermediateSourceConfigurationProperty := &IntermediateSourceConfigurationProperty{
//   	IntermediateS3Path: jsii.String("intermediateS3Path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-intermediatesourceconfiguration.html
//
type CfnMatchingWorkflowPropsMixin_IntermediateSourceConfigurationProperty struct {
	// The Amazon S3 location (bucket and prefix).
	//
	// For example: `s3://provider_bucket/DOC-EXAMPLE-BUCKET`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-intermediatesourceconfiguration.html#cfn-entityresolution-matchingworkflow-intermediatesourceconfiguration-intermediates3path
	//
	IntermediateS3Path *string `field:"optional" json:"intermediateS3Path" yaml:"intermediateS3Path"`
}


package mixinsawsentityresolution


// An object containing the `providerServiceARN` , `intermediateSourceConfiguration` , and `providerConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   providerPropertiesProperty := &ProviderPropertiesProperty{
//   	IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   		IntermediateS3Path: jsii.String("intermediateS3Path"),
//   	},
//   	ProviderConfiguration: map[string]*string{
//   		"providerConfigurationKey": jsii.String("providerConfiguration"),
//   	},
//   	ProviderServiceArn: jsii.String("providerServiceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-providerproperties.html
//
type CfnIdMappingWorkflowPropsMixin_ProviderPropertiesProperty struct {
	// The Amazon S3 location that temporarily stores your data while it processes.
	//
	// Your information won't be saved permanently.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-providerproperties.html#cfn-entityresolution-idmappingworkflow-providerproperties-intermediatesourceconfiguration
	//
	IntermediateSourceConfiguration interface{} `field:"optional" json:"intermediateSourceConfiguration" yaml:"intermediateSourceConfiguration"`
	// The required configuration fields to use with the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-providerproperties.html#cfn-entityresolution-idmappingworkflow-providerproperties-providerconfiguration
	//
	ProviderConfiguration interface{} `field:"optional" json:"providerConfiguration" yaml:"providerConfiguration"`
	// The ARN of the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idmappingworkflow-providerproperties.html#cfn-entityresolution-idmappingworkflow-providerproperties-providerservicearn
	//
	ProviderServiceArn *string `field:"optional" json:"providerServiceArn" yaml:"providerServiceArn"`
}


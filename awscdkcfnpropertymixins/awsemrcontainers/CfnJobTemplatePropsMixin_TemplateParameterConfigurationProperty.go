package awsemrcontainers


// The configuration of a job template parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   templateParameterConfigurationProperty := &TemplateParameterConfigurationProperty{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-templateparameterconfiguration.html
//
type CfnJobTemplatePropsMixin_TemplateParameterConfigurationProperty struct {
	// The default value for the job template parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-templateparameterconfiguration.html#cfn-emrcontainers-jobtemplate-templateparameterconfiguration-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The type of the job template parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-templateparameterconfiguration.html#cfn-emrcontainers-jobtemplate-templateparameterconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


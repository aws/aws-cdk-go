package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &logConfigurationProperty{
//   	logDriver: jsii.String("logDriver"),
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	secretOptions: []interface{}{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
type CfnService_LogConfigurationProperty struct {
	// `CfnService.LogConfigurationProperty.LogDriver`.
	LogDriver *string `field:"optional" json:"logDriver" yaml:"logDriver"`
	// `CfnService.LogConfigurationProperty.Options`.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// `CfnService.LogConfigurationProperty.SecretOptions`.
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


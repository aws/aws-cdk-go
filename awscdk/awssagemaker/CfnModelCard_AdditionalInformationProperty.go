package awssagemaker


// Additional information about the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalInformationProperty := &AdditionalInformationProperty{
//   	CaveatsAndRecommendations: jsii.String("caveatsAndRecommendations"),
//   	CustomDetails: map[string]*string{
//   		"customDetailsKey": jsii.String("customDetails"),
//   	},
//   	EthicalConsiderations: jsii.String("ethicalConsiderations"),
//   }
//
type CfnModelCard_AdditionalInformationProperty struct {
	// Caveats and recommendations for those who might use this model in their applications.
	CaveatsAndRecommendations *string `field:"optional" json:"caveatsAndRecommendations" yaml:"caveatsAndRecommendations"`
	// Any additional information to document about the model.
	CustomDetails interface{} `field:"optional" json:"customDetails" yaml:"customDetails"`
	// Any ethical considerations documented by the model card author.
	EthicalConsiderations *string `field:"optional" json:"ethicalConsiderations" yaml:"ethicalConsiderations"`
}


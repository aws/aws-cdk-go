package awsapigateway


// The attributes of an imported Stage.
//
// Example:
//   var restApi iRestApi
//
//   importedStage := apigateway.Stage_FromStageAttributes(this, jsii.String("imported-stage"), &StageAttributes{
//   	StageName: jsii.String("myStageName"),
//   	RestApi: RestApi,
//   })
//
//   importedStage.AddApiKey(jsii.String("MyApiKey"))
//
type StageAttributes struct {
	// The RestApi that the stage belongs to.
	RestApi IRestApi `field:"required" json:"restApi" yaml:"restApi"`
	// The name of the stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}


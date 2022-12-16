package awscodepipeline


// Allows you to control where to place a new Stage when it's added to the Pipeline.
//
// Note that you can provide only one of the below properties -
// specifying more than one will result in a validation error.
//
// Example:
//   // Insert a new Stage at an arbitrary point
//   var pipeline pipeline
//   var anotherStage iStage
//   var yetAnotherStage iStage
//
//
//   someStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("SomeStage"),
//   	placement: &stagePlacement{
//   		// note: you can only specify one of the below properties
//   		rightBefore: anotherStage,
//   		justAfter: yetAnotherStage,
//   	},
//   })
//
// See: #justAfter.
//
type StagePlacement struct {
	// Inserts the new Stage as a child of the given Stage (changing its current child Stage, if it had one).
	JustAfter IStage `field:"optional" json:"justAfter" yaml:"justAfter"`
	// Inserts the new Stage as a parent of the given Stage (changing its current parent Stage, if it had one).
	RightBefore IStage `field:"optional" json:"rightBefore" yaml:"rightBefore"`
}


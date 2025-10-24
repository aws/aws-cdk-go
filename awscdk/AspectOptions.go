package awscdk


// Options when Applying an Aspect.
//
// Example:
//   type mutatingAspect struct {
//   }
//
//   func (this *mutatingAspect) visit(node IConstruct) {}
//
//   type validationAspect struct {
//   }
//
//   func (this *validationAspect) visit(node IConstruct) {}
//
//   stack := awscdk.Newstack()
//
//   awscdk.Aspects_Of(stack).Add(NewMutatingAspect(), &AspectOptions{
//   	Priority: awscdk.AspectPriority_MUTATING(),
//   }) // Run first (mutating aspects)
//   awscdk.Aspects_Of(stack).Add(NewValidationAspect(), &AspectOptions{
//   	Priority: awscdk.AspectPriority_READONLY(),
//   })
//
type AspectOptions struct {
	// The priority value to apply on an Aspect. Priority must be a non-negative integer.
	//
	// Aspects that have same priority value are not guaranteed to be
	// executed in a consistent order.
	// Default: AspectPriority.DEFAULT
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}


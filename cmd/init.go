package cmd

// Init initializes all commands using Cobra
// Individual command init() functions in each command file will register themselves
func Init() error {
	return Execute()
}

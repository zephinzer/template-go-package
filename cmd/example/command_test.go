package main

func (t *Tests) Test_GetCommand() {
	cmd := GetCommand()
	t.Equal("example", cmd.Use)
	t.NotNil(cmd.Run)
}

package e2e_test

type fakeCompose struct {
	cmd []string
}

func (f *fakeCompose) buildCmd(m *machineTestBuilder) []string {
	cmd := []string{"compose"}
	cmd = append(cmd, "env")
	f.cmd = cmd
	return cmd
}

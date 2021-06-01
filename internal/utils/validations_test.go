package utils

import (
	"asimov-tool-cli/internal/defines"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IsLower_True(t *testing.T) {
	out := IsLower("ab")
	require.True(t, out)
}
func Test_IsLower_False(t *testing.T) {
	out := IsLower("aB")
	require.False(t, out)
}
func Test_IsValidVersion_True(t *testing.T) {
	out := IsValidVersion("123.456.789")
	require.True(t, out)
}
func Test_IsValidVersion_False(t *testing.T) {
	out := IsValidVersion("a.b.c")
	require.False(t, out)
}
func Test_IsValidTarget_True(t *testing.T) {
	out1 := IsValidTarget(defines.TargetARM)
	out2 := IsValidTarget(defines.TargetLinux)
	out3 := IsValidTarget(defines.TargetMac)
	out4 := IsValidTarget(defines.TargetWin)
	out5 := IsValidTarget("")

	require.True(t, out1)
	require.True(t, out2)
	require.True(t, out3)
	require.True(t, out4)
	require.True(t, out5)
}
func Test_IsValidTarget_False(t *testing.T) {
	out := IsValidTarget("a")

	require.False(t, out)
}
func Test_IsValidVersionName_True(t *testing.T) {
	out := IsValidVersionName("project-branch-v1.2.3")

	require.True(t, out)
}
func Test_IsValidVersionName_False(t *testing.T) {
	out := IsValidVersionName("Project-branch-v1.2.3")

	require.False(t, out)
}

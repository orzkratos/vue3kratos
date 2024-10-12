package demos

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/orzkratos/vue3kratos/internal/utils"
	"github.com/orzkratos/vue3kratos/vue3kratos2gen"
	"github.com/orzkratos/vue3kratos/vue3npm"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/runpath"
)

func TestGenWithGo(t *testing.T) {
	caseGen(t, vue3kratos2gen.NewGen())
}

func TestGenUseNpm(t *testing.T) {
	caseGen(t, vue3kratos2gen.NewGen().SetWkLanguage(vue3kratos2gen.JS).SetExecInRoot(vue3npm.SourceRoot()))
}

func caseGen(t *testing.T, gen *vue3kratos2gen.GenType) {
	dataB := done.VAE(os.ReadFile(runpath.PARENT.Join("greeter.client.ts-B.txt"))).Nice()
	dataA := done.VAE(os.ReadFile(runpath.PARENT.Join("greeter.client.ts-A.txt"))).Nice()

	targetRoot, err := os.MkdirTemp("", utils.NewUUID())
	require.NoError(t, err)
	defer func() {
		done.Done(os.RemoveAll(targetRoot))
	}()

	targetPath := filepath.Join(targetRoot, utils.NewUUID()+".client.ts")
	done.Done(os.WriteFile(targetPath, dataB, 0644))

	t.Log("gen")
	gen.GenGrpcViaHttpInRoot(targetRoot)
	t.Log("---")

	newContent := done.VAE(os.ReadFile(targetPath)).Nice()
	require.Equal(t, string(dataA), string(newContent))
}

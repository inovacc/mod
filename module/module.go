package module

import "github.com/inovacc/mod/module/internal/modload"

type Module struct {
	Path    string
	Version string
	ModTime string
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Params() (string, string) {
	return m.Path, m.Version
}

func (m *Module) GetModule(name, version string) ([]Module, error) {
	modload.ForceUseModules = true
	modload.ExplicitWriteGoMod = true
	modload.AllowMissingModuleImports()
	modload.Init()
	return nil, nil
}

// automatically generated by stateify.

// +build 386 amd64 arm64

package arch

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (m *MmapLayout) StateTypeName() string {
	return "pkg/sentry/arch.MmapLayout"
}

func (m *MmapLayout) StateFields() []string {
	return []string{
		"MinAddr",
		"MaxAddr",
		"BottomUpBase",
		"TopDownBase",
		"DefaultDirection",
		"MaxStackRand",
	}
}

func (m *MmapLayout) beforeSave() {}

// +checklocksignore
func (m *MmapLayout) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.MinAddr)
	stateSinkObject.Save(1, &m.MaxAddr)
	stateSinkObject.Save(2, &m.BottomUpBase)
	stateSinkObject.Save(3, &m.TopDownBase)
	stateSinkObject.Save(4, &m.DefaultDirection)
	stateSinkObject.Save(5, &m.MaxStackRand)
}

func (m *MmapLayout) afterLoad() {}

// +checklocksignore
func (m *MmapLayout) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.MinAddr)
	stateSourceObject.Load(1, &m.MaxAddr)
	stateSourceObject.Load(2, &m.BottomUpBase)
	stateSourceObject.Load(3, &m.TopDownBase)
	stateSourceObject.Load(4, &m.DefaultDirection)
	stateSourceObject.Load(5, &m.MaxStackRand)
}

func (a *AuxEntry) StateTypeName() string {
	return "pkg/sentry/arch.AuxEntry"
}

func (a *AuxEntry) StateFields() []string {
	return []string{
		"Key",
		"Value",
	}
}

func (a *AuxEntry) beforeSave() {}

// +checklocksignore
func (a *AuxEntry) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.Key)
	stateSinkObject.Save(1, &a.Value)
}

func (a *AuxEntry) afterLoad() {}

// +checklocksignore
func (a *AuxEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.Key)
	stateSourceObject.Load(1, &a.Value)
}

func (s *SignalStack) StateTypeName() string {
	return "pkg/sentry/arch.SignalStack"
}

func (s *SignalStack) StateFields() []string {
	return []string{
		"Addr",
		"Flags",
		"Size",
	}
}

func (s *SignalStack) beforeSave() {}

// +checklocksignore
func (s *SignalStack) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Addr)
	stateSinkObject.Save(1, &s.Flags)
	stateSinkObject.Save(2, &s.Size)
}

func (s *SignalStack) afterLoad() {}

// +checklocksignore
func (s *SignalStack) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Addr)
	stateSourceObject.Load(1, &s.Flags)
	stateSourceObject.Load(2, &s.Size)
}

func (s *SignalInfo) StateTypeName() string {
	return "pkg/sentry/arch.SignalInfo"
}

func (s *SignalInfo) StateFields() []string {
	return []string{
		"Signo",
		"Errno",
		"Code",
		"Fields",
	}
}

func (s *SignalInfo) beforeSave() {}

// +checklocksignore
func (s *SignalInfo) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Signo)
	stateSinkObject.Save(1, &s.Errno)
	stateSinkObject.Save(2, &s.Code)
	stateSinkObject.Save(3, &s.Fields)
}

func (s *SignalInfo) afterLoad() {}

// +checklocksignore
func (s *SignalInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Signo)
	stateSourceObject.Load(1, &s.Errno)
	stateSourceObject.Load(2, &s.Code)
	stateSourceObject.Load(3, &s.Fields)
}

func init() {
	state.Register((*MmapLayout)(nil))
	state.Register((*AuxEntry)(nil))
	state.Register((*SignalStack)(nil))
	state.Register((*SignalInfo)(nil))
}

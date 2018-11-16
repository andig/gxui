// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

type SwitchMode int
type SwitchButtonMode int

const (
	Linear SwitchMode = iota
	Circular
	Smart SwitchButtonMode = iota
	Constant
)

type PanelHolder interface {
	Control
	AddPanel(panel Control, name string)
	AddPanelAt(panel Control, name string, index int)
	RemovePanel(panel Control)
	Select(int)
	SelectNext()
	SelectPrev()
	PanelCount() int
	PanelIndex(Control) int
	Panel(int) Control
	Tab(int) Control
	Begin() int
	End() int
	SetMaxLabelLength(int)
	SwitchMode() SwitchMode
	SetSwitchMode(SwitchMode)
	SwitchButtonMode() SwitchButtonMode
	SetSwitchButtonMode(SwitchButtonMode)
}

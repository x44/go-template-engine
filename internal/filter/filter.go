package filter

import (
	"fmt"
	"strings"
)

const (
	CONDITION_START_TRUE  = "//?"
	CONDITION_START_FALSE = "//!"
	CONDITION_END         = "//-"
)

type Filter struct {
	lines             []string
	vars              map[string]bool
	conditionStack    *ConditionStack
	conditionSum      bool
	isOnConditionLine bool
}

func New(lines []string) *Filter {
	return &Filter{
		lines:          prepareLines(lines),
		vars:           make(map[string]bool),
		conditionStack: NewConditionStack(),
		conditionSum:   false,
	}
}

func prepareLines(lines []string) []string {
	out := make([]string, len(lines))
	for i, line := range lines {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, CONDITION_START_TRUE) || strings.HasPrefix(t, CONDITION_START_FALSE) || strings.HasPrefix(t, CONDITION_END) {
			out[i] = t
		} else {
			out[i] = line
		}
	}
	return out
}

func (p *Filter) SetVar(name string, b bool) {
	p.vars[name] = b
}

func (p *Filter) getVar(name string) bool {
	return p.vars[name]
}

func (p *Filter) ClearFlags() {
	for k := range p.vars {
		delete(p.vars, k)
	}
}

func (p *Filter) Process() ([]string, error) {
	p.resetConditions()
	out := make([]string, 0)

	for _, line := range p.lines {
		err := p.processCondition(line)
		if err != nil {
			return nil, err
		}
		if p.checkConditions() {
			out = append(out, line)
		}
	}

	if p.conditionStack.Size() != 0 {
		return nil, fmt.Errorf("missing end-of-condition")
	}

	return out, nil
}

func (p *Filter) resetConditions() {
	p.conditionStack.Clear()
	p.conditionSum = true
}

func (p *Filter) processCondition(line string) error {
	wantsTrue := strings.HasPrefix(line, CONDITION_START_TRUE)
	wantsFalse := strings.HasPrefix(line, CONDITION_START_FALSE)

	if wantsTrue || wantsFalse {
		p.isOnConditionLine = true
		flagName := line[3:]
		flag := p.getVar(flagName)
		condition := (wantsTrue && flag) || (wantsFalse && !flag)
		p.conditionStack.Push(condition)
		p.updateConditionSum()
	} else if strings.HasPrefix(line, CONDITION_END) {
		p.isOnConditionLine = true
		_, err := p.conditionStack.Pop()
		if err != nil {
			return fmt.Errorf("end-of-condition without start-of-condition")
		}
		p.updateConditionSum()
	} else {
		p.isOnConditionLine = false
	}
	return nil
}

func (p *Filter) checkConditions() bool {
	if p.isOnConditionLine {
		return false
	}
	return p.conditionSum
}

func (p *Filter) updateConditionSum() {
	sum, err := p.conditionStack.Sum()
	if err != nil {
		sum = true
	}
	p.conditionSum = sum
}

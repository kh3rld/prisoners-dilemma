package outcome

import "github.com/kh3rld/prisoners-dilemma/pkg/common"

type Outcome struct {
	common.Outcome
}

func (o *Outcome) GetOutcome() common.Outcome {
	return o.Outcome
}

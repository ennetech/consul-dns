package common

type ChainText struct {
	txt string
}

func NewChainText() *ChainText {
	return &ChainText{}
}

func (cl *ChainText) Append(txt string) {
	cl.txt += txt + "\n"
}

func (cl *ChainText) String() string {
	return cl.txt
}

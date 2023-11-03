package memento

type InputText struct {
	content string
}

// 追加数据
func (in *InputText) Append(content string) {
	in.content += content
}

// 获取数据
func (in *InputText) GetText() string {
	return in.content
}

// 创建快照
func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

// 从快照恢复
func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}

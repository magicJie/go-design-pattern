package state

import "fmt"

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approve(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

type IState interface {
	// 审批通过
	Approve(m *Machine)
	// 驳回
	Reject(m *Machine)
	// 获取当前状态名称
	GetName() string
}

// 直属领导审批
type leaderApproveState struct{}

func (leaderApproveState) Approve(m *Machine) {
	fmt.Println(("leader 审批成功"))
	m.SetState(GetFinanceApproveState())
}

func (leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

// 财务审批
type financeApproveState struct{}

func (f financeApproveState) Approve(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("触发打款操作")
}

// 拒绝
func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

// 获取名字
func (f financeApproveState) GetName() string {
	return "FinanceApproveState"
}

func GetFinanceApproveState() IState {
	return &financeApproveState{}
}

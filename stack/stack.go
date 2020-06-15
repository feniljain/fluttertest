package stack

//Error represents errors in stack
type stackError interface {
	error
	IsUnderFlowError() bool
}

//UnderFlowError represents the stack underflow error
type UnderFlowError struct{}

func (e *UnderFlowError) Error() string {
	return "stack: Stack Underflow Error"
}

//IsUnderFlowError represents whether it is underflow error or not
func (e *UnderFlowError) IsUnderFlowError() bool {
	return true
}

//Stack struct holds the information about stack
type Stack struct {
	stack []string
	Top   int
}

//Init function is used to intialize a stack
func Init() Stack {
	return Stack{
		stack: make([]string, 0),
		Top:   -1,
	}
}

//Push function pushes the given string in the stack
func (st *Stack) Push(s string) {
	st.stack = append(st.stack, s)
	st.Top++
}

//Pop function pops the top string from the stack
func (st *Stack) Pop() (string, error) {
	if st.Top == 0 {
		return "", &UnderFlowError{}
	}
	lastElement := st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	st.Top--
	return lastElement, nil
}

//TopElement returns the top element of the stack
func (st *Stack) TopElement() (string, error) {
	if st.Top == -1 {
		return "", &UnderFlowError{}
	}
	return st.stack[st.Top], nil
}

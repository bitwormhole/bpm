package bpm

// CommandHandler 表示一个命令处理器
type CommandHandler interface {
	Execute(cc *CommandContext) error
}

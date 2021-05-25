package decorator

import "fmt"

type INotification interface {
	SetDecorator(INotification) INotification
	Send() error
}

type CEMailNotification struct {
	notify INotification
}

// SetDecorator 设置装饰器,为其组装成链式结构
func (cemn *CEMailNotification) SetDecorator(n INotification) INotification {
	if cemn.notify != nil {
		cemn.notify.SetDecorator(n)
	} else {
		cemn.notify = n
	}
	return cemn
}

func (cemn *CEMailNotification) Send() error {
	fmt.Println("发送了邮件")
	if cemn.notify != nil {
		return cemn.notify.Send()
	}
	return nil
}

type CPushNotification struct {
	notify INotification
}

// SetDecorator 设置装饰器,为其组装成链式结构
func (cemn *CPushNotification) SetDecorator(n INotification) INotification {
	// 此处装饰设置,决定着调用链是正序还是反序.
	if cemn.notify != nil {
		cemn.notify.SetDecorator(n)
	} else {
		cemn.notify = n
	}
	return cemn
}

func (cemn *CPushNotification) Send() error {
	fmt.Println("发送了推送消息")
	if cemn.notify != nil {
		return cemn.notify.Send()
	}
	return nil
}

type CIMNotification struct {
	notify INotification
}

// SetDecorator 设置装饰器,为其组装成链式结构
func (cemn *CIMNotification) SetDecorator(n INotification) INotification {
	// 此处装饰设置,决定着调用链是正序还是反序.
	if cemn.notify != nil {
		cemn.notify.SetDecorator(n)
	} else {
		cemn.notify = n
	}
	return cemn
}

func (cemn *CIMNotification) Send() error {
	fmt.Println("发送了IM消息")
	if cemn.notify != nil {
		return cemn.notify.Send()
	}
	return nil
}
